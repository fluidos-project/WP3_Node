/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	nodecorev1alpha1 "fluidos.eu/node/api/nodecore/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReservationSpec defines the desired state of Reservation
type ReservationSpec struct {

	// TransactionID is the ID of the transaction that this reservation is part of
	TransactionID string `json:"transactionID,omitempty"`

	// This is the Node identity of the buyer FLUIDOS Node.
	Buyer nodecorev1alpha1.NodeIdentity `json:"buyer"`

	// This is the Node identity of the seller FLUIDOS Node.
	Seller nodecorev1alpha1.NodeIdentity `json:"seller"`

	// Parition is the partition of the flavour that is being reserved
	Partition nodecorev1alpha1.FlavourSelector `json:"partition,omitempty"`

	// Reserve indicates if the reservation is a reserve or not
	Reserve bool `json:"reserve,omitempty"`

	// Purchase indicates if the reservation is an purchase or not
	Purchase bool `json:"purchase,omitempty"`

	// PeeringCandidate is the reference to the PeeringCandidate of the Reservation
	PeeringCandidate nodecorev1alpha1.GenericRef `json:"peeringCandidate,omitempty"`

	// Contract is the reference to the Contract of the Reservation
	Contract nodecorev1alpha1.GenericRef `json:"contract,omitempty"`
}

// ReservationStatus defines the observed state of Reservation
type ReservationStatus struct {
	// This is the current phase of the reservation
	Phase nodecorev1alpha1.PhaseStatus `json:"phase"`

	// ReservePhase is the current phase of the reservation
	ReservePhase nodecorev1alpha1.Phase `json:"reservePhase,omitempty"`

	// PurchasePhase is the current phase of the reservation
	PurchasePhase nodecorev1alpha1.Phase `json:"purchasePhase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Reservation is the Schema for the reservations API
type Reservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReservationSpec   `json:"spec,omitempty"`
	Status ReservationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReservationList contains a list of Reservation
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Reservation{}, &ReservationList{})
}
