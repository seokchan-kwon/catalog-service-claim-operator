// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceClaim) DeepCopyInto(out *CatalogServiceClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceClaim.
func (in *CatalogServiceClaim) DeepCopy() *CatalogServiceClaim {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogServiceClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceClaimList) DeepCopyInto(out *CatalogServiceClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogServiceClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceClaimList.
func (in *CatalogServiceClaimList) DeepCopy() *CatalogServiceClaimList {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogServiceClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceClaimSpec) DeepCopyInto(out *CatalogServiceClaimSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceClaimSpec.
func (in *CatalogServiceClaimSpec) DeepCopy() *CatalogServiceClaimSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceClaimStatus) DeepCopyInto(out *CatalogServiceClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceClaimStatus.
func (in *CatalogServiceClaimStatus) DeepCopy() *CatalogServiceClaimStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceClaimStatus)
	in.DeepCopyInto(out)
	return out
}
