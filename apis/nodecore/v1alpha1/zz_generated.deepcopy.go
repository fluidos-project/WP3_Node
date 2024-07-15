//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allocation) DeepCopyInto(out *Allocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allocation.
func (in *Allocation) DeepCopy() *Allocation {
	if in == nil {
		return nil
	}
	out := new(Allocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Allocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationList) DeepCopyInto(out *AllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Allocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationList.
func (in *AllocationList) DeepCopy() *AllocationList {
	if in == nil {
		return nil
	}
	out := new(AllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationSpec) DeepCopyInto(out *AllocationSpec) {
	*out = *in
	out.Contract = in.Contract
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationSpec.
func (in *AllocationSpec) DeepCopy() *AllocationSpec {
	if in == nil {
		return nil
	}
	out := new(AllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationStatus) DeepCopyInto(out *AllocationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationStatus.
func (in *AllocationStatus) DeepCopy() *AllocationStatus {
	if in == nil {
		return nil
	}
	out := new(AllocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CarbonFootprint) DeepCopyInto(out *CarbonFootprint) {
	*out = *in
	if in.Operational != nil {
		in, out := &in.Operational, &out.Operational
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CarbonFootprint.
func (in *CarbonFootprint) DeepCopy() *CarbonFootprint {
	if in == nil {
		return nil
	}
	out := new(CarbonFootprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	in.ConfigurationData.DeepCopyInto(&out.ConfigurationData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flavor) DeepCopyInto(out *Flavor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flavor.
func (in *Flavor) DeepCopy() *Flavor {
	if in == nil {
		return nil
	}
	out := new(Flavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Flavor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorList) DeepCopyInto(out *FlavorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorList.
func (in *FlavorList) DeepCopy() *FlavorList {
	if in == nil {
		return nil
	}
	out := new(FlavorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlavorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorSpec) DeepCopyInto(out *FlavorSpec) {
	*out = *in
	in.FlavorType.DeepCopyInto(&out.FlavorType)
	out.Owner = in.Owner
	out.Price = in.Price
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(Location)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorSpec.
func (in *FlavorSpec) DeepCopy() *FlavorSpec {
	if in == nil {
		return nil
	}
	out := new(FlavorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorStatus) DeepCopyInto(out *FlavorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorStatus.
func (in *FlavorStatus) DeepCopy() *FlavorStatus {
	if in == nil {
		return nil
	}
	out := new(FlavorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorType) DeepCopyInto(out *FlavorType) {
	*out = *in
	in.TypeData.DeepCopyInto(&out.TypeData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorType.
func (in *FlavorType) DeepCopy() *FlavorType {
	if in == nil {
		return nil
	}
	out := new(FlavorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPU) DeepCopyInto(out *GPU) {
	*out = *in
	out.Cores = in.Cores.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPU.
func (in *GPU) DeepCopy() *GPU {
	if in == nil {
		return nil
	}
	out := new(GPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericRef) DeepCopyInto(out *GenericRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericRef.
func (in *GenericRef) DeepCopy() *GenericRef {
	if in == nil {
		return nil
	}
	out := new(GenericRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8Slice) DeepCopyInto(out *K8Slice) {
	*out = *in
	in.Characteristics.DeepCopyInto(&out.Characteristics)
	in.Properties.DeepCopyInto(&out.Properties)
	in.Policies.DeepCopyInto(&out.Policies)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8Slice.
func (in *K8Slice) DeepCopy() *K8Slice {
	if in == nil {
		return nil
	}
	out := new(K8Slice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SliceCharacteristics) DeepCopyInto(out *K8SliceCharacteristics) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	out.Pods = in.Pods.DeepCopy()
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(GPU)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SliceCharacteristics.
func (in *K8SliceCharacteristics) DeepCopy() *K8SliceCharacteristics {
	if in == nil {
		return nil
	}
	out := new(K8SliceCharacteristics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SliceConfiguration) DeepCopyInto(out *K8SliceConfiguration) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	out.Pods = in.Pods.DeepCopy()
	if in.Gpu != nil {
		in, out := &in.Gpu, &out.Gpu
		*out = new(GPU)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SliceConfiguration.
func (in *K8SliceConfiguration) DeepCopy() *K8SliceConfiguration {
	if in == nil {
		return nil
	}
	out := new(K8SliceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SliceSelector) DeepCopyInto(out *K8SliceSelector) {
	*out = *in
	if in.CPUFilter != nil {
		in, out := &in.CPUFilter, &out.CPUFilter
		*out = new(ResourceQuantityFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.MemoryFilter != nil {
		in, out := &in.MemoryFilter, &out.MemoryFilter
		*out = new(ResourceQuantityFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.PodsFilter != nil {
		in, out := &in.PodsFilter, &out.PodsFilter
		*out = new(ResourceQuantityFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageFilter != nil {
		in, out := &in.StorageFilter, &out.StorageFilter
		*out = new(ResourceQuantityFilter)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SliceSelector.
func (in *K8SliceSelector) DeepCopy() *K8SliceSelector {
	if in == nil {
		return nil
	}
	out := new(K8SliceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LiqoCredentials) DeepCopyInto(out *LiqoCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LiqoCredentials.
func (in *LiqoCredentials) DeepCopy() *LiqoCredentials {
	if in == nil {
		return nil
	}
	out := new(LiqoCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Location) DeepCopyInto(out *Location) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Location.
func (in *Location) DeepCopy() *Location {
	if in == nil {
		return nil
	}
	out := new(Location)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeIdentity) DeepCopyInto(out *NodeIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeIdentity.
func (in *NodeIdentity) DeepCopy() *NodeIdentity {
	if in == nil {
		return nil
	}
	out := new(NodeIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Partitionability) DeepCopyInto(out *Partitionability) {
	*out = *in
	out.CPUMin = in.CPUMin.DeepCopy()
	out.MemoryMin = in.MemoryMin.DeepCopy()
	out.PodsMin = in.PodsMin.DeepCopy()
	out.GpuMin = in.GpuMin.DeepCopy()
	out.CPUStep = in.CPUStep.DeepCopy()
	out.MemoryStep = in.MemoryStep.DeepCopy()
	out.PodsStep = in.PodsStep.DeepCopy()
	out.GpuStep = in.GpuStep.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Partitionability.
func (in *Partitionability) DeepCopy() *Partitionability {
	if in == nil {
		return nil
	}
	out := new(Partitionability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseStatus) DeepCopyInto(out *PhaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseStatus.
func (in *PhaseStatus) DeepCopy() *PhaseStatus {
	if in == nil {
		return nil
	}
	out := new(PhaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policies) DeepCopyInto(out *Policies) {
	*out = *in
	in.Partitionability.DeepCopyInto(&out.Partitionability)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policies.
func (in *Policies) DeepCopy() *Policies {
	if in == nil {
		return nil
	}
	out := new(Policies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Price) DeepCopyInto(out *Price) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Price.
func (in *Price) DeepCopy() *Price {
	if in == nil {
		return nil
	}
	out := new(Price)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Properties) DeepCopyInto(out *Properties) {
	*out = *in
	if in.SecurityStandards != nil {
		in, out := &in.SecurityStandards, &out.SecurityStandards
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CarbonFootprint != nil {
		in, out := &in.CarbonFootprint, &out.CarbonFootprint
		*out = new(CarbonFootprint)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Properties.
func (in *Properties) DeepCopy() *Properties {
	if in == nil {
		return nil
	}
	out := new(Properties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMatchSelector) DeepCopyInto(out *ResourceMatchSelector) {
	*out = *in
	out.Value = in.Value.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMatchSelector.
func (in *ResourceMatchSelector) DeepCopy() *ResourceMatchSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceMatchSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuantityFilter) DeepCopyInto(out *ResourceQuantityFilter) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuantityFilter.
func (in *ResourceQuantityFilter) DeepCopy() *ResourceQuantityFilter {
	if in == nil {
		return nil
	}
	out := new(ResourceQuantityFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRangeSelector) DeepCopyInto(out *ResourceRangeSelector) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRangeSelector.
func (in *ResourceRangeSelector) DeepCopy() *ResourceRangeSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceRangeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Solver) DeepCopyInto(out *Solver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Solver.
func (in *Solver) DeepCopy() *Solver {
	if in == nil {
		return nil
	}
	out := new(Solver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Solver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolverList) DeepCopyInto(out *SolverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Solver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolverList.
func (in *SolverList) DeepCopy() *SolverList {
	if in == nil {
		return nil
	}
	out := new(SolverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolverSpec) DeepCopyInto(out *SolverSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolverSpec.
func (in *SolverSpec) DeepCopy() *SolverSpec {
	if in == nil {
		return nil
	}
	out := new(SolverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolverStatus) DeepCopyInto(out *SolverStatus) {
	*out = *in
	out.SolverPhase = in.SolverPhase
	out.Allocation = in.Allocation
	out.Contract = in.Contract
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolverStatus.
func (in *SolverStatus) DeepCopy() *SolverStatus {
	if in == nil {
		return nil
	}
	out := new(SolverStatus)
	in.DeepCopyInto(out)
	return out
}
