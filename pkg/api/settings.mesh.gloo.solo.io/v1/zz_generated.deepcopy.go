// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for settings.mesh.gloo.solo.io/v1 resources

package v1

import (
    runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for Settings

func (in *Settings) DeepCopyInto(out *Settings) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *Settings) DeepCopy() *Settings {
    if in == nil {
        return nil
    }
    out := new(Settings)
    in.DeepCopyInto(out)
    return out
}

func (in *Settings) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *SettingsList) DeepCopyInto(out *SettingsList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]Settings, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *SettingsList) DeepCopy() *SettingsList {
    if in == nil {
        return nil
    }
    out := new(SettingsList)
    in.DeepCopyInto(out)
    return out
}

func (in *SettingsList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for Dashboard

func (in *Dashboard) DeepCopyInto(out *Dashboard) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *Dashboard) DeepCopy() *Dashboard {
    if in == nil {
        return nil
    }
    out := new(Dashboard)
    in.DeepCopyInto(out)
    return out
}

func (in *Dashboard) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *DashboardList) DeepCopyInto(out *DashboardList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]Dashboard, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *DashboardList) DeepCopy() *DashboardList {
    if in == nil {
        return nil
    }
    out := new(DashboardList)
    in.DeepCopyInto(out)
    return out
}

func (in *DashboardList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

