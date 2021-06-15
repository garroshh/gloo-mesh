// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for discovery.mesh.gloo.solo.io/v1 resources

package v1

import (
    runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for Destination

func (in *Destination) DeepCopyInto(out *Destination) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *Destination) DeepCopy() *Destination {
    if in == nil {
        return nil
    }
    out := new(Destination)
    in.DeepCopyInto(out)
    return out
}

func (in *Destination) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *DestinationList) DeepCopyInto(out *DestinationList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]Destination, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *DestinationList) DeepCopy() *DestinationList {
    if in == nil {
        return nil
    }
    out := new(DestinationList)
    in.DeepCopyInto(out)
    return out
}

func (in *DestinationList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for Workload

func (in *Workload) DeepCopyInto(out *Workload) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *Workload) DeepCopy() *Workload {
    if in == nil {
        return nil
    }
    out := new(Workload)
    in.DeepCopyInto(out)
    return out
}

func (in *Workload) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]Workload, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *WorkloadList) DeepCopy() *WorkloadList {
    if in == nil {
        return nil
    }
    out := new(WorkloadList)
    in.DeepCopyInto(out)
    return out
}

func (in *WorkloadList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for Mesh

func (in *Mesh) DeepCopyInto(out *Mesh) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *Mesh) DeepCopy() *Mesh {
    if in == nil {
        return nil
    }
    out := new(Mesh)
    in.DeepCopyInto(out)
    return out
}

func (in *Mesh) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *MeshList) DeepCopyInto(out *MeshList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]Mesh, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *MeshList) DeepCopy() *MeshList {
    if in == nil {
        return nil
    }
    out := new(MeshList)
    in.DeepCopyInto(out)
    return out
}

func (in *MeshList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

