//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cron) DeepCopyInto(out *Cron) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cron.
func (in *Cron) DeepCopy() *Cron {
	if in == nil {
		return nil
	}
	out := new(Cron)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cron) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronInitParameters) DeepCopyInto(out *CronInitParameters) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronInitParameters.
func (in *CronInitParameters) DeepCopy() *CronInitParameters {
	if in == nil {
		return nil
	}
	out := new(CronInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronList) DeepCopyInto(out *CronList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cron, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronList.
func (in *CronList) DeepCopy() *CronList {
	if in == nil {
		return nil
	}
	out := new(CronList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronObservation) DeepCopyInto(out *CronObservation) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronObservation.
func (in *CronObservation) DeepCopy() *CronObservation {
	if in == nil {
		return nil
	}
	out := new(CronObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronParameters) DeepCopyInto(out *CronParameters) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronParameters.
func (in *CronParameters) DeepCopy() *CronParameters {
	if in == nil {
		return nil
	}
	out := new(CronParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronSpec) DeepCopyInto(out *CronSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronSpec.
func (in *CronSpec) DeepCopy() *CronSpec {
	if in == nil {
		return nil
	}
	out := new(CronSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronStatus) DeepCopyInto(out *CronStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronStatus.
func (in *CronStatus) DeepCopy() *CronStatus {
	if in == nil {
		return nil
	}
	out := new(CronStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainInitParameters) DeepCopyInto(out *DomainInitParameters) {
	*out = *in
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainInitParameters.
func (in *DomainInitParameters) DeepCopy() *DomainInitParameters {
	if in == nil {
		return nil
	}
	out := new(DomainInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainObservation) DeepCopyInto(out *DomainObservation) {
	*out = *in
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainObservation.
func (in *DomainObservation) DeepCopy() *DomainObservation {
	if in == nil {
		return nil
	}
	out := new(DomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainParameters) DeepCopyInto(out *DomainParameters) {
	*out = *in
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainParameters.
func (in *DomainParameters) DeepCopy() *DomainParameters {
	if in == nil {
		return nil
	}
	out := new(DomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionInitParameters) DeepCopyInto(out *FunctionInitParameters) {
	*out = *in
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.HTTPOption != nil {
		in, out := &in.HTTPOption, &out.HTTPOption
		*out = new(string)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.MaxScale != nil {
		in, out := &in.MaxScale, &out.MaxScale
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceIDRef != nil {
		in, out := &in.NamespaceIDRef, &out.NamespaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceIDSelector != nil {
		in, out := &in.NamespaceIDSelector, &out.NamespaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Privacy != nil {
		in, out := &in.Privacy, &out.Privacy
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
	if in.ZipHash != nil {
		in, out := &in.ZipHash, &out.ZipHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionInitParameters.
func (in *FunctionInitParameters) DeepCopy() *FunctionInitParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespace) DeepCopyInto(out *FunctionNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespace.
func (in *FunctionNamespace) DeepCopy() *FunctionNamespace {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespaceInitParameters) DeepCopyInto(out *FunctionNamespaceInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespaceInitParameters.
func (in *FunctionNamespaceInitParameters) DeepCopy() *FunctionNamespaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespaceList) DeepCopyInto(out *FunctionNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FunctionNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespaceList.
func (in *FunctionNamespaceList) DeepCopy() *FunctionNamespaceList {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespaceObservation) DeepCopyInto(out *FunctionNamespaceObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RegistryEndpoint != nil {
		in, out := &in.RegistryEndpoint, &out.RegistryEndpoint
		*out = new(string)
		**out = **in
	}
	if in.RegistryNamespaceID != nil {
		in, out := &in.RegistryNamespaceID, &out.RegistryNamespaceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespaceObservation.
func (in *FunctionNamespaceObservation) DeepCopy() *FunctionNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespaceParameters) DeepCopyInto(out *FunctionNamespaceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SecretEnvironmentVariablesSecretRef != nil {
		in, out := &in.SecretEnvironmentVariablesSecretRef, &out.SecretEnvironmentVariablesSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespaceParameters.
func (in *FunctionNamespaceParameters) DeepCopy() *FunctionNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespaceSpec) DeepCopyInto(out *FunctionNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespaceSpec.
func (in *FunctionNamespaceSpec) DeepCopy() *FunctionNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionNamespaceStatus) DeepCopyInto(out *FunctionNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionNamespaceStatus.
func (in *FunctionNamespaceStatus) DeepCopy() *FunctionNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.CPULimit != nil {
		in, out := &in.CPULimit, &out.CPULimit
		*out = new(float64)
		**out = **in
	}
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.HTTPOption != nil {
		in, out := &in.HTTPOption, &out.HTTPOption
		*out = new(string)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxScale != nil {
		in, out := &in.MaxScale, &out.MaxScale
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationID != nil {
		in, out := &in.OrganizationID, &out.OrganizationID
		*out = new(string)
		**out = **in
	}
	if in.Privacy != nil {
		in, out := &in.Privacy, &out.Privacy
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
	if in.ZipHash != nil {
		in, out := &in.ZipHash, &out.ZipHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.HTTPOption != nil {
		in, out := &in.HTTPOption, &out.HTTPOption
		*out = new(string)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.MaxScale != nil {
		in, out := &in.MaxScale, &out.MaxScale
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceIDRef != nil {
		in, out := &in.NamespaceIDRef, &out.NamespaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceIDSelector != nil {
		in, out := &in.NamespaceIDSelector, &out.NamespaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Privacy != nil {
		in, out := &in.Privacy, &out.Privacy
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.SecretEnvironmentVariablesSecretRef != nil {
		in, out := &in.SecretEnvironmentVariablesSecretRef, &out.SecretEnvironmentVariablesSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(float64)
		**out = **in
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
	if in.ZipHash != nil {
		in, out := &in.ZipHash, &out.ZipHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Token) DeepCopyInto(out *Token) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Token.
func (in *Token) DeepCopy() *Token {
	if in == nil {
		return nil
	}
	out := new(Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Token) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenInitParameters) DeepCopyInto(out *TokenInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceIDRef != nil {
		in, out := &in.NamespaceIDRef, &out.NamespaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceIDSelector != nil {
		in, out := &in.NamespaceIDSelector, &out.NamespaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenInitParameters.
func (in *TokenInitParameters) DeepCopy() *TokenInitParameters {
	if in == nil {
		return nil
	}
	out := new(TokenInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenList) DeepCopyInto(out *TokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Token, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenList.
func (in *TokenList) DeepCopy() *TokenList {
	if in == nil {
		return nil
	}
	out := new(TokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenObservation) DeepCopyInto(out *TokenObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenObservation.
func (in *TokenObservation) DeepCopy() *TokenObservation {
	if in == nil {
		return nil
	}
	out := new(TokenObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenParameters) DeepCopyInto(out *TokenParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceIDRef != nil {
		in, out := &in.NamespaceIDRef, &out.NamespaceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceIDSelector != nil {
		in, out := &in.NamespaceIDSelector, &out.NamespaceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenParameters.
func (in *TokenParameters) DeepCopy() *TokenParameters {
	if in == nil {
		return nil
	}
	out := new(TokenParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSpec) DeepCopyInto(out *TokenSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSpec.
func (in *TokenSpec) DeepCopy() *TokenSpec {
	if in == nil {
		return nil
	}
	out := new(TokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenStatus) DeepCopyInto(out *TokenStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenStatus.
func (in *TokenStatus) DeepCopy() *TokenStatus {
	if in == nil {
		return nil
	}
	out := new(TokenStatus)
	in.DeepCopyInto(out)
	return out
}
