// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/route_table.proto

package v1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
//
//RouteTable is a resource which can be referenced either from the top level VirtualHost resource, or from
//other RouteTables. It's primary use is to organizationally and logically separate the configuration of routes,
//so that the responsbilities of route configuration and maintenance can be divided between teams where appropriate.
type RouteTableSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of HTTP routes define routing actions to be taken for incoming HTTP requests whose host header matches
	// this virtual host. If the request matches more than one route in the list, the first route matched will be selected.
	// If the list of routes is empty, the virtual host will be ignored by Gloo.
	Routes []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteTableSpec) Reset() {
	*x = RouteTableSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableSpec) ProtoMessage() {}

func (x *RouteTableSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableSpec.ProtoReflect.Descriptor instead.
func (*RouteTableSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescGZIP(), []int{0}
}

func (x *RouteTableSpec) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type RouteTableStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the RouteTable metadata.
	// If the `observedGeneration` does not match `metadata.generation`,
	// Gloo Mesh has not processed the most recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// Any warnings found while processing this generation of the resource.
	Warnings []string `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// List of resources which have selected this RouteTable. Can be VirtualHosts or other RouteTables
	SelectedBy []*SelectedBy `protobuf:"bytes,4,rep,name=selected_by,json=selectedBy,proto3" json:"selected_by,omitempty"`
	// List of child RouteTables that this RouteTable delegates to
	SelectedRouteTables []*v1.ObjectRef `protobuf:"bytes,5,rep,name=selected_route_tables,json=selectedRouteTables,proto3" json:"selected_route_tables,omitempty"`
}

func (x *RouteTableStatus) Reset() {
	*x = RouteTableStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTableStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTableStatus) ProtoMessage() {}

func (x *RouteTableStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTableStatus.ProtoReflect.Descriptor instead.
func (*RouteTableStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescGZIP(), []int{1}
}

func (x *RouteTableStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *RouteTableStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RouteTableStatus) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *RouteTableStatus) GetSelectedBy() []*SelectedBy {
	if x != nil {
		return x.SelectedBy
	}
	return nil
}

func (x *RouteTableStatus) GetSelectedRouteTables() []*v1.ObjectRef {
	if x != nil {
		return x.SelectedRouteTables
	}
	return nil
}

type SelectedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of resource selecting this RouteTable
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Namespace of resource selecting this RouteTable
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Type of resource selecting this RouteTable. Can be
	// VirtualGateway, VirtualHost, or RouteTable.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SelectedBy) Reset() {
	*x = SelectedBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectedBy) ProtoMessage() {}

func (x *SelectedBy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectedBy.ProtoReflect.Descriptor instead.
func (*SelectedBy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescGZIP(), []int{2}
}

func (x *SelectedBy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SelectedBy) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SelectedBy) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDesc = []byte{
	0x0a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x27, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x46, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x54, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x50, 0x0a, 0x15, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x52, 0x13, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x5a, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_goTypes = []interface{}{
	(*RouteTableSpec)(nil),   // 0: networking.enterprise.mesh.gloo.solo.io.RouteTableSpec
	(*RouteTableStatus)(nil), // 1: networking.enterprise.mesh.gloo.solo.io.RouteTableStatus
	(*SelectedBy)(nil),       // 2: networking.enterprise.mesh.gloo.solo.io.SelectedBy
	(*Route)(nil),            // 3: networking.enterprise.mesh.gloo.solo.io.Route
	(*v1.ObjectRef)(nil),     // 4: core.skv2.solo.io.ObjectRef
}
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_depIdxs = []int32{
	3, // 0: networking.enterprise.mesh.gloo.solo.io.RouteTableSpec.routes:type_name -> networking.enterprise.mesh.gloo.solo.io.Route
	2, // 1: networking.enterprise.mesh.gloo.solo.io.RouteTableStatus.selected_by:type_name -> networking.enterprise.mesh.gloo.solo.io.SelectedBy
	4, // 2: networking.enterprise.mesh.gloo.solo.io.RouteTableStatus.selected_route_tables:type_name -> core.skv2.solo.io.ObjectRef
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTableStatus); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectedBy); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_route_table_proto_depIdxs = nil
}
