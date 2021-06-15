// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/enterprise/observability/v1/access_logging.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	_ "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	v11 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

//
//Describes a record of access logs sourced from a set of workloads and optionally
//filtered based on request criteria.
type AccessLogRecordSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select the workloads to be configured to emit access logs.
	// Leave empty to apply to all workloads managed by Gloo Mesh.
	WorkloadSelectors []*v1.WorkloadSelector `protobuf:"bytes,1,rep,name=workload_selectors,json=workloadSelectors,proto3" json:"workload_selectors,omitempty"`
	// Configure criteria for determining which access logs will be recorded.
	// The list is disjunctive, a request will be recorded if it matches any filter.
	// Leave empty to emit all access logs.
	Filters []*AccessLogRecordSpec_Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	// Specify request headers to include in access logs.
	IncludedRequestHeaders []string `protobuf:"bytes,3,rep,name=included_request_headers,json=includedRequestHeaders,proto3" json:"included_request_headers,omitempty"`
	// Specify response headers to include in access logs.
	IncludedResponseHeaders []string `protobuf:"bytes,4,rep,name=included_response_headers,json=includedResponseHeaders,proto3" json:"included_response_headers,omitempty"`
	// Specify response trailers to include in access logs.
	IncludedResponseTrailers []string `protobuf:"bytes,5,rep,name=included_response_trailers,json=includedResponseTrailers,proto3" json:"included_response_trailers,omitempty"`
	// Specify filter state objects to include in access logs
	IncludedFilterStateObjects []string `protobuf:"bytes,6,rep,name=included_filter_state_objects,json=includedFilterStateObjects,proto3" json:"included_filter_state_objects,omitempty"`
}

func (x *AccessLogRecordSpec) Reset() {
	*x = AccessLogRecordSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogRecordSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogRecordSpec) ProtoMessage() {}

func (x *AccessLogRecordSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogRecordSpec.ProtoReflect.Descriptor instead.
func (*AccessLogRecordSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescGZIP(), []int{0}
}

func (x *AccessLogRecordSpec) GetWorkloadSelectors() []*v1.WorkloadSelector {
	if x != nil {
		return x.WorkloadSelectors
	}
	return nil
}

func (x *AccessLogRecordSpec) GetFilters() []*AccessLogRecordSpec_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *AccessLogRecordSpec) GetIncludedRequestHeaders() []string {
	if x != nil {
		return x.IncludedRequestHeaders
	}
	return nil
}

func (x *AccessLogRecordSpec) GetIncludedResponseHeaders() []string {
	if x != nil {
		return x.IncludedResponseHeaders
	}
	return nil
}

func (x *AccessLogRecordSpec) GetIncludedResponseTrailers() []string {
	if x != nil {
		return x.IncludedResponseTrailers
	}
	return nil
}

func (x *AccessLogRecordSpec) GetIncludedFilterStateObjects() []string {
	if x != nil {
		return x.IncludedFilterStateObjects
	}
	return nil
}

type AccessLogRecordStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the AccessLogRecord metadata.
	// If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource, will only show accepted if it has been successfully
	// applied to all target workloads.
	State v1.ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=common.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// Any errors encountered during processing. Also reported to any Workloads that this object applies to.
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	// references to workloads that this AccessLogRecord applies to
	Workloads []*v11.ObjectRef `protobuf:"bytes,4,rep,name=workloads,proto3" json:"workloads,omitempty"`
}

func (x *AccessLogRecordStatus) Reset() {
	*x = AccessLogRecordStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogRecordStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogRecordStatus) ProtoMessage() {}

func (x *AccessLogRecordStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogRecordStatus.ProtoReflect.Descriptor instead.
func (*AccessLogRecordStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescGZIP(), []int{1}
}

func (x *AccessLogRecordStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *AccessLogRecordStatus) GetState() v1.ApprovalState {
	if x != nil {
		return x.State
	}
	return v1.ApprovalState_PENDING
}

func (x *AccessLogRecordStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *AccessLogRecordStatus) GetWorkloads() []*v11.ObjectRef {
	if x != nil {
		return x.Workloads
	}
	return nil
}

// Specify criteria for recording access logs. A request must match all specified criteria to be recorded.
type AccessLogRecordSpec_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Criteria type.
	//
	// Types that are assignable to Type:
	//	*AccessLogRecordSpec_Filter_StatusCodeMatcher
	//	*AccessLogRecordSpec_Filter_HeaderMatcher
	Type isAccessLogRecordSpec_Filter_Type `protobuf_oneof:"type"`
}

func (x *AccessLogRecordSpec_Filter) Reset() {
	*x = AccessLogRecordSpec_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLogRecordSpec_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLogRecordSpec_Filter) ProtoMessage() {}

func (x *AccessLogRecordSpec_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLogRecordSpec_Filter.ProtoReflect.Descriptor instead.
func (*AccessLogRecordSpec_Filter) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AccessLogRecordSpec_Filter) GetType() isAccessLogRecordSpec_Filter_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *AccessLogRecordSpec_Filter) GetStatusCodeMatcher() *v1.StatusCodeMatcher {
	if x, ok := x.GetType().(*AccessLogRecordSpec_Filter_StatusCodeMatcher); ok {
		return x.StatusCodeMatcher
	}
	return nil
}

func (x *AccessLogRecordSpec_Filter) GetHeaderMatcher() *v1.HeaderMatcher {
	if x, ok := x.GetType().(*AccessLogRecordSpec_Filter_HeaderMatcher); ok {
		return x.HeaderMatcher
	}
	return nil
}

type isAccessLogRecordSpec_Filter_Type interface {
	isAccessLogRecordSpec_Filter_Type()
}

type AccessLogRecordSpec_Filter_StatusCodeMatcher struct {
	// Matches against a response status code. Omit to match any status code.
	StatusCodeMatcher *v1.StatusCodeMatcher `protobuf:"bytes,1,opt,name=status_code_matcher,json=statusCodeMatcher,proto3,oneof"`
}

type AccessLogRecordSpec_Filter_HeaderMatcher struct {
	// Matches against a request or response header. Omit to match any headers.
	HeaderMatcher *v1.HeaderMatcher `protobuf:"bytes,2,opt,name=header_matcher,json=headerMatcher,proto3,oneof"`
}

func (*AccessLogRecordSpec_Filter_StatusCodeMatcher) isAccessLogRecordSpec_Filter_Type() {}

func (*AccessLogRecordSpec_Filter_HeaderMatcher) isAccessLogRecordSpec_Filter_Type() {}

var File_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDesc = []byte{
	0x0a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x05, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x59, 0x0a, 0x12, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x60, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3c,
	0x0a, 0x1a, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x18, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x1d,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x1a, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a,
	0xc1, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x13, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x0e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a,
	0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_goTypes = []interface{}{
	(*AccessLogRecordSpec)(nil),        // 0: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec
	(*AccessLogRecordStatus)(nil),      // 1: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordStatus
	(*AccessLogRecordSpec_Filter)(nil), // 2: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec.Filter
	(*v1.WorkloadSelector)(nil),        // 3: common.mesh.gloo.solo.io.WorkloadSelector
	(v1.ApprovalState)(0),              // 4: common.mesh.gloo.solo.io.ApprovalState
	(*v11.ObjectRef)(nil),              // 5: core.skv2.solo.io.ObjectRef
	(*v1.StatusCodeMatcher)(nil),       // 6: common.mesh.gloo.solo.io.StatusCodeMatcher
	(*v1.HeaderMatcher)(nil),           // 7: common.mesh.gloo.solo.io.HeaderMatcher
}
var file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_depIdxs = []int32{
	3, // 0: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec.workload_selectors:type_name -> common.mesh.gloo.solo.io.WorkloadSelector
	2, // 1: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec.filters:type_name -> observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec.Filter
	4, // 2: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordStatus.state:type_name -> common.mesh.gloo.solo.io.ApprovalState
	5, // 3: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordStatus.workloads:type_name -> core.skv2.solo.io.ObjectRef
	6, // 4: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec.Filter.status_code_matcher:type_name -> common.mesh.gloo.solo.io.StatusCodeMatcher
	7, // 5: observability.enterprise.mesh.gloo.solo.io.AccessLogRecordSpec.Filter.header_matcher:type_name -> common.mesh.gloo.solo.io.HeaderMatcher
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogRecordSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogRecordStatus); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLogRecordSpec_Filter); i {
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
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AccessLogRecordSpec_Filter_StatusCodeMatcher)(nil),
		(*AccessLogRecordSpec_Filter_HeaderMatcher)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_observability_v1_access_logging_proto_depIdxs = nil
}
