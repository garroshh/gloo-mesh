// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for networking.smh.solo.io/v1alpha2 resources

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for TrafficPolicy

func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *TrafficPolicy) DeepCopy() *TrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

func (in *TrafficPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *TrafficPolicyList) DeepCopyInto(out *TrafficPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrafficPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *TrafficPolicyList) DeepCopy() *TrafficPolicyList {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicyList)
	in.DeepCopyInto(out)
	return out
}

func (in *TrafficPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for AccessPolicy

func (in *AccessPolicy) DeepCopyInto(out *AccessPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *AccessPolicy) DeepCopy() *AccessPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessPolicy)
	in.DeepCopyInto(out)
	return out
}

func (in *AccessPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AccessPolicyList) DeepCopyInto(out *AccessPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *AccessPolicyList) DeepCopy() *AccessPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyList)
	in.DeepCopyInto(out)
	return out
}

func (in *AccessPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for VirtualMesh

func (in *VirtualMesh) DeepCopyInto(out *VirtualMesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *VirtualMesh) DeepCopy() *VirtualMesh {
	if in == nil {
		return nil
	}
	out := new(VirtualMesh)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMeshList) DeepCopyInto(out *VirtualMeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMeshList) DeepCopy() *VirtualMeshList {
	if in == nil {
		return nil
	}
	out := new(VirtualMeshList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
