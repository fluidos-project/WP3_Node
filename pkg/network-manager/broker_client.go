// Copyright 2022-2024 FLUIDOS Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkmanager

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkv1alpha1 "github.com/fluidos-project/node/apis/network/v1alpha1"
	nodecorev1alpha1 "github.com/fluidos-project/node/apis/nodecore/v1alpha1"
	"github.com/fluidos-project/node/pkg/utils/flags"
	"github.com/fluidos-project/node/pkg/utils/getters"
	"github.com/fluidos-project/node/pkg/utils/namings"
	"github.com/fluidos-project/node/pkg/utils/resourceforge"
)

// clusterRole
// +kubebuilder:rbac:groups=network.fluidos.eu,resources=brokers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.fluidos.eu,resources=brokers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch

// NetworkManager keeps all the necessary class data.
type BrokerClient struct {
	ID *nodecorev1alpha1.NodeIdentity //this node

	subFlag bool
	pubFlag bool

	serverName string
	serverAddr string
	clCert     *corev1.Secret
	caCert     *corev1.Secret

	conn         *amqp.Connection
	ch           *amqp.Channel
	exchangeName string
	routingKey   string

	queueName   string
	msgs        <-chan amqp.Delivery
	outboundMsg []byte
	confirms    chan amqp.Confirmation
}

// Setup the Broker Client from NM reconcile
func (bc *BrokerClient) SetupBrokerClient(ctx context.Context, cl client.Client, broker *networkv1alpha1.Broker) error {
	klog.Info("Setting up Broker Client routines")

	var err error

	nodeIdentity := getters.GetNodeIdentity(ctx, cl)
	if nodeIdentity == nil {
		return fmt.Errorf("failed to get Node Identity")
	}

	bc.ID = nodeIdentity
	bc.serverName = broker.Spec.Name
	bc.serverAddr = broker.Spec.Address
	bc.clCert = broker.Spec.ClCert
	bc.caCert = broker.Spec.CaCert
	bc.exchangeName = "DefaultPeerRequest"

	//setting pub/sub
	if strings.EqualFold(broker.Spec.Role, "publisher") {
		bc.pubFlag = true
		bc.subFlag = false
	} else if strings.EqualFold(broker.Spec.Role, "subscriber") {
		bc.pubFlag = false
		bc.subFlag = true
	} else {
		bc.pubFlag = true
		bc.subFlag = true
	}

	bc.outboundMsg, err = json.Marshal(bc.ID)
	if err != nil {
		return err
	}

	// Extract certs and key
	clientCert, ok := bc.clCert.Data["tls.crt"]
	if !ok {
		klog.Fatalf("cert error: %v", ok)
		return fmt.Errorf("missing certificate: 'tls.crt' not found in clCert Data")
	}

	clientKey, ok := bc.clCert.Data["tls.key"]
	if !ok {
		klog.Fatalf("key error: %v", ok)
		return fmt.Errorf("missing key: 'tls.key' not found in clCert Data")
	}

	caCertData, ok := bc.caCert.Data["tls.crt"]
	if !ok {
		klog.Fatalf("CA cert error: %v", ok)
		return fmt.Errorf("missing certificate: 'tls.crt' not found in CACert Data")
	}

	// load client cert and privKey
	cert, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		klog.Fatalf("error X509KeyPair: %v", err)
		return err
	}

	// load CAcert
	caCertPool := x509.NewCertPool()
	ok = caCertPool.AppendCertsFromPEM(caCertData)
	if !ok {
		klog.Fatalf("AppendCertsFromPEM error: %v", ok)
	}

	// TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ServerName:   bc.serverName,
		//InsecureSkipVerify: false, //skippa/consenti verifica SSL in produzione
	}

	bc.routingKey, err = extractCNfromCert(&clientCert) // Routing key specifica per il topic
	if err != nil {
		klog.Fatalf("Common Name extraction error: %v", err)
	}
	bc.queueName = bc.routingKey

	bc.rabbitConfig(tlsConfig)

	return err
}

// Execute the Network Manager Broker routines.
func (bc *BrokerClient) ExecuteBrokerClient(ctx context.Context, cl client.Client) error {
	// Start sending messages

	klog.Info("EXECUTING Broker Client routines")
	var err error
	if bc.pubFlag {
		go func() {
			if err = bc.publishOnBroker(ctx); err != nil {
				klog.ErrorS(err, "Error sending advertisement")
			}
		}()
	}

	// Start receiving messages
	if bc.subFlag {
		go func() {
			if err = bc.readMsgOnBroker(ctx, cl); err != nil {
				klog.ErrorS(err, "Error receiving advertisement")
			}
		}()
	}
	return err
}

func (bc *BrokerClient) publishOnBroker(ctx context.Context) error {

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:

			// Pubblicazione del messaggio sull'exchange con la routing key
			err := bc.ch.Publish(
				bc.exchangeName, // Nome dell'exchange
				bc.routingKey,   // Routing key per instradare il messaggio
				true,            // Mandatory, se non routable errore
				false,           // Immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        bc.outboundMsg,
				})
			if err != nil {
				klog.Fatalf("Error pub message: %v", err)
			} else {
				//fmt.Printf("Message on exchange '%s' with routing key '%s': \n", bc.exchangeName, bc.routingKey)
			}

			select {
			case confirm := <-bc.confirms:
				if confirm.Ack {
					klog.Info("Message successfully published!")
				} else {
					klog.Info("Message failed to publish!")
				}
			case <-time.After(5 * time.Second): // Timeout
				klog.Info("No confirmation received, message status unknown.")
			}

		case <-ctx.Done():
			ticker.Stop()
			klog.Info("Ticker stopped\n")
			return nil
		}
	}
}

func (bc *BrokerClient) readMsgOnBroker(ctx context.Context, cl client.Client) error {

	klog.Info("Listening from Broker")
	for d := range bc.msgs {

		klog.Info("Received remote advertisement from BROKER\n")
		var remote NetworkManager
		err := json.Unmarshal(d.Body, &remote.ID)

		if err != nil {
			klog.Error("Error unmarshalling message: ", err)
			continue
		}
		// Check if received advertisement is remote
		if bc.ID.IP != remote.ID.IP {

			//create knownCluster CR
			kc := &networkv1alpha1.KnownCluster{}

			if err := cl.Get(ctx, client.ObjectKey{Name: namings.ForgeKnownClusterName(remote.ID.NodeID), Namespace: flags.FluidosNamespace}, kc); err != nil {
				if client.IgnoreNotFound(err) == nil {
					klog.Info("KnownCluster not found: creating")

					// Create new KnownCluster CR
					if err := cl.Create(ctx, resourceforge.ForgeKnownCluster(remote.ID.NodeID, remote.ID.IP)); err != nil {
						return err
					}
					klog.InfoS("KnownCluster created", "ID", remote.ID.NodeID)
				}
			} else {
				klog.Info("KnownCluster already present: updating")
				kc.UpdateStatus()

				// Update fetched KnownCluster CR
				err := cl.Status().Update(ctx, kc)
				if err != nil {
					return err
				}
				klog.InfoS("KnownCluster updated", "ID", kc.ObjectMeta.Name)
			}
		}
	}
	return nil
}

func extractCNfromCert(certPEM *[]byte) (string, error) { //certPath string

	// Decode PEM cert
	block, _ := pem.Decode(*certPEM)
	if block == nil {
		klog.Fatalf("Error decoding certificate PEM in CN extraction")
	}

	// Parsing X.509
	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		klog.Fatalf("Error parsing certificate X.509 in CN extraction: %v", err)
	}

	CN := cert.Subject.CommonName

	return strings.TrimSpace(CN), err
}

func (bc *BrokerClient) rabbitConfig(tlsConfig *tls.Config) error {

	var err error
	//____________________________RABBIT
	// EXTERNAL AUTH RabbitMQ
	config := amqp.Config{
		SASL:            []amqp.Authentication{&amqp.ExternalAuth{}}, // Autenticazione EXTERNAL
		TLSClientConfig: tlsConfig,                                   // Configurazione TLS
		Vhost:           "/",                                         // Vhost
		Heartbeat:       10 * time.Second,                            // Intervallo heartbeat
	}

	// Config connection

	rabbitMQURL := "amqps://fluidos.top-ix.org:5671/"
	// RABBITMQ conn
	bc.conn, err = amqp.DialConfig(rabbitMQURL, config) // conn, err := amqp.DialTLS(rabbitMQURL, tlsConfig)//conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		klog.Fatalf("RabbitMQ connection error: %v", err)
	}
	//defer bc.conn.Close()                                 GRACEFUL

	// channel creation
	bc.ch, err = bc.conn.Channel()
	if err != nil {
		klog.Fatalf("channel creation error: %v", err)
	}
	//defer bc.ch.Close()                                    GRACEFUL

	// Queue subscrition
	bc.msgs, err = bc.ch.Consume(
		bc.queueName, // Nome della coda
		"",           // Nome del consumatore (vuoto per uno generato automaticamente)
		true,         // AutoAck: conferma automatica della ricezione del messaggio
		false,        // Exclusive: solo questo consumatore può accedere alla coda
		true,         //false,        // NoLocal: non riceve messaggi pubblicati dalla stessa connessione       --TO BE SET--
		false,        // NoWait: non aspettare la conferma del server
		nil,          // Argomenti aggiuntivi
	)
	if err != nil {
		klog.Fatalf("Error subscribing queue: %s", err)
	}

	// Write confirm broker
	if err := bc.ch.Confirm(false); err != nil {
		log.Fatalf("Failed to enable publisher confirms: %v", err)
	}

	// Channels for write confirm
	bc.confirms = bc.ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	klog.InfoS("Node", "ID", bc.ID.NodeID, "Client Address", bc.ID.IP, "Server Address", bc.serverAddr, "RoutingKey", bc.routingKey)

	return err
}