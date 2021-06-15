// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/networking/v1/base.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Describes how to match a given string in HTTP headers. Match is case-sensitive.
type StringMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string match type.
	//
	// Types that are assignable to MatchType:
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Regex
	//	*StringMatch_Suffix
	MatchType isStringMatch_MatchType `protobuf_oneof:"match_type"`
	//If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no effect for the regex match.
	IgnoreCase bool `protobuf:"varint,5,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
}

func (x *StringMatch) Reset() {
	*x = StringMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatch) ProtoMessage() {}

func (x *StringMatch) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatch.ProtoReflect.Descriptor instead.
func (*StringMatch) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescGZIP(), []int{0}
}

func (m *StringMatch) GetMatchType() isStringMatch_MatchType {
	if m != nil {
		return m.MatchType
	}
	return nil
}

func (x *StringMatch) GetExact() string {
	if x, ok := x.GetMatchType().(*StringMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *StringMatch) GetPrefix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *StringMatch) GetRegex() string {
	if x, ok := x.GetMatchType().(*StringMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

func (x *StringMatch) GetSuffix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (x *StringMatch) GetIgnoreCase() bool {
	if x != nil {
		return x.IgnoreCase
	}
	return false
}

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
}

type StringMatch_Exact struct {
	// Exact string match.
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatch_Prefix struct {
	// Prefix-based match.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatch_Regex struct {
	// ECMAscript style regex-based match.
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

type StringMatch_Suffix struct {
	// Suffix-based match.
	Suffix string `protobuf:"bytes,4,opt,name=suffix,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType() {}

func (*StringMatch_Prefix) isStringMatch_MatchType() {}

func (*StringMatch_Regex) isStringMatch_MatchType() {}

func (*StringMatch_Suffix) isStringMatch_MatchType() {}

// ObjectMeta is a simplified clone of the Kubernetes ObjectMeta used to represent object metadata
// for Kubernetes objects passed as messages in the NetworkingExtensions API.
type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the Kubernetes name of the object
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the Kubernetes namespace of the object
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// the Kubernetes clusterName of the object (used internally by Gloo Mesh)
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// the Kubernetes resourceVersion of the object
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// the Kubernetes labels on the object
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the Kubernetes annotations on the object
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectMeta) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectMeta) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ObjectMeta) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *ObjectMeta) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ObjectMeta) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// Copy of Kubernetes ObjectMeta from https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/v1/generated.proto
// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
type TypeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	// In CamelCase.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion string `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
}

func (x *TypeMeta) Reset() {
	*x = TypeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeMeta) ProtoMessage() {}

func (x *TypeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeMeta.ProtoReflect.Descriptor instead.
func (*TypeMeta) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescGZIP(), []int{2}
}

func (x *TypeMeta) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *TypeMeta) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01,
	0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a,
	0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x16, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61,
	0x73, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xb2, 0x03, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x4c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5b, 0x0a,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x4a, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_goTypes = []interface{}{
	(*StringMatch)(nil), // 0: networking.mesh.gloo.solo.io.StringMatch
	(*ObjectMeta)(nil),  // 1: networking.mesh.gloo.solo.io.ObjectMeta
	(*TypeMeta)(nil),    // 2: networking.mesh.gloo.solo.io.TypeMeta
	nil,                 // 3: networking.mesh.gloo.solo.io.ObjectMeta.LabelsEntry
	nil,                 // 4: networking.mesh.gloo.solo.io.ObjectMeta.AnnotationsEntry
}
var file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_depIdxs = []int32{
	3, // 0: networking.mesh.gloo.solo.io.ObjectMeta.labels:type_name -> networking.mesh.gloo.solo.io.ObjectMeta.LabelsEntry
	4, // 1: networking.mesh.gloo.solo.io.ObjectMeta.annotations:type_name -> networking.mesh.gloo.solo.io.ObjectMeta.AnnotationsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_networking_v1_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Regex)(nil),
		(*StringMatch_Suffix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_base_proto_depIdxs = nil
}
