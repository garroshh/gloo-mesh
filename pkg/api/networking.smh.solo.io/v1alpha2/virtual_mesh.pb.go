// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/virtual_mesh.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//If ENABLED, by default disallow traffic to all Services in the VirtualMesh unless explicitly allowed through AccessControlPolicies.
//If DISABLED, by default allow traffic to all Services in the VirtualMesh.
//If MESH_DEFAULT, the default value depends on the type service mesh:
//Istio: false
//Appmesh: true
type VirtualMeshSpec_EnforcementPolicy int32

const (
	VirtualMeshSpec_MESH_DEFAULT VirtualMeshSpec_EnforcementPolicy = 0
	VirtualMeshSpec_ENABLED      VirtualMeshSpec_EnforcementPolicy = 1
	VirtualMeshSpec_DISABLED     VirtualMeshSpec_EnforcementPolicy = 2
)

var VirtualMeshSpec_EnforcementPolicy_name = map[int32]string{
	0: "MESH_DEFAULT",
	1: "ENABLED",
	2: "DISABLED",
}

var VirtualMeshSpec_EnforcementPolicy_value = map[string]int32{
	"MESH_DEFAULT": 0,
	"ENABLED":      1,
	"DISABLED":     2,
}

func (x VirtualMeshSpec_EnforcementPolicy) String() string {
	return proto.EnumName(VirtualMeshSpec_EnforcementPolicy_name, int32(x))
}

func (VirtualMeshSpec_EnforcementPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 0}
}

// A VirtualMesh represents a logical grouping of meshes for
// shared configuration and cross-mesh interoperability.
//
// VirtualMeshes are used to configure things like shared trust roots (for mTLS)
// and federation of services (for cross-cluster networking).
//
// Currently, VirtualMeshes can only be constructed from Istio
// meshes.
type VirtualMeshSpec struct {
	// User-provided display name for the virtual mesh.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The meshes contained in this virtual mesh.
	Meshes []*v1.ObjectRef `protobuf:"bytes,2,rep,name=meshes,proto3" json:"meshes,omitempty"`
	// Sets a shared Certificate Authority across the defined meshes.
	CertificateAuthority *VirtualMeshSpec_CertificateAuthority `protobuf:"bytes,3,opt,name=certificate_authority,json=certificateAuthority,proto3" json:"certificate_authority,omitempty"`
	// Expose services to traffic from any workload in
	// the virtual mesh using Service Federation.
	Federation *VirtualMeshSpec_Federation `protobuf:"bytes,4,opt,name=federation,proto3" json:"federation,omitempty"`
	// Types that are valid to be assigned to TrustModel:
	//	*VirtualMeshSpec_Shared
	//	*VirtualMeshSpec_Limited
	TrustModel           isVirtualMeshSpec_TrustModel      `protobuf_oneof:"trust_model"`
	EnforceAccessControl VirtualMeshSpec_EnforcementPolicy `protobuf:"varint,7,opt,name=enforce_access_control,json=enforceAccessControl,proto3,enum=networking.smh.solo.io.VirtualMeshSpec_EnforcementPolicy" json:"enforce_access_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *VirtualMeshSpec) Reset()         { *m = VirtualMeshSpec{} }
func (m *VirtualMeshSpec) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec) ProtoMessage()    {}
func (*VirtualMeshSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0}
}
func (m *VirtualMeshSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec.Unmarshal(m, b)
}
func (m *VirtualMeshSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec.Merge(m, src)
}
func (m *VirtualMeshSpec) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec.Size(m)
}
func (m *VirtualMeshSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec proto.InternalMessageInfo

type isVirtualMeshSpec_TrustModel interface {
	isVirtualMeshSpec_TrustModel()
	Equal(interface{}) bool
}

type VirtualMeshSpec_Shared struct {
	Shared *VirtualMeshSpec_SharedTrust `protobuf:"bytes,5,opt,name=shared,proto3,oneof" json:"shared,omitempty"`
}
type VirtualMeshSpec_Limited struct {
	Limited *VirtualMeshSpec_LimitedTrust `protobuf:"bytes,6,opt,name=limited,proto3,oneof" json:"limited,omitempty"`
}

func (*VirtualMeshSpec_Shared) isVirtualMeshSpec_TrustModel()  {}
func (*VirtualMeshSpec_Limited) isVirtualMeshSpec_TrustModel() {}

func (m *VirtualMeshSpec) GetTrustModel() isVirtualMeshSpec_TrustModel {
	if m != nil {
		return m.TrustModel
	}
	return nil
}

func (m *VirtualMeshSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualMeshSpec) GetMeshes() []*v1.ObjectRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *VirtualMeshSpec) GetCertificateAuthority() *VirtualMeshSpec_CertificateAuthority {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

func (m *VirtualMeshSpec) GetFederation() *VirtualMeshSpec_Federation {
	if m != nil {
		return m.Federation
	}
	return nil
}

func (m *VirtualMeshSpec) GetShared() *VirtualMeshSpec_SharedTrust {
	if x, ok := m.GetTrustModel().(*VirtualMeshSpec_Shared); ok {
		return x.Shared
	}
	return nil
}

func (m *VirtualMeshSpec) GetLimited() *VirtualMeshSpec_LimitedTrust {
	if x, ok := m.GetTrustModel().(*VirtualMeshSpec_Limited); ok {
		return x.Limited
	}
	return nil
}

func (m *VirtualMeshSpec) GetEnforceAccessControl() VirtualMeshSpec_EnforcementPolicy {
	if m != nil {
		return m.EnforceAccessControl
	}
	return VirtualMeshSpec_MESH_DEFAULT
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_Shared)(nil),
		(*VirtualMeshSpec_Limited)(nil),
	}
}

type VirtualMeshSpec_CertificateAuthority struct {
	// If omitted, defaults to builtin.
	//
	// Types that are valid to be assigned to Type:
	//	*VirtualMeshSpec_CertificateAuthority_Builtin_
	//	*VirtualMeshSpec_CertificateAuthority_Provided_
	Type                 isVirtualMeshSpec_CertificateAuthority_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority) Reset()         { *m = VirtualMeshSpec_CertificateAuthority{} }
func (m *VirtualMeshSpec_CertificateAuthority) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_CertificateAuthority) ProtoMessage()    {}
func (*VirtualMeshSpec_CertificateAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 0}
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority proto.InternalMessageInfo

type isVirtualMeshSpec_CertificateAuthority_Type interface {
	isVirtualMeshSpec_CertificateAuthority_Type()
	Equal(interface{}) bool
}

type VirtualMeshSpec_CertificateAuthority_Builtin_ struct {
	Builtin *VirtualMeshSpec_CertificateAuthority_Builtin `protobuf:"bytes,1,opt,name=builtin,proto3,oneof" json:"builtin,omitempty"`
}
type VirtualMeshSpec_CertificateAuthority_Provided_ struct {
	Provided *VirtualMeshSpec_CertificateAuthority_Provided `protobuf:"bytes,2,opt,name=provided,proto3,oneof" json:"provided,omitempty"`
}

func (*VirtualMeshSpec_CertificateAuthority_Builtin_) isVirtualMeshSpec_CertificateAuthority_Type() {}
func (*VirtualMeshSpec_CertificateAuthority_Provided_) isVirtualMeshSpec_CertificateAuthority_Type() {
}

func (m *VirtualMeshSpec_CertificateAuthority) GetType() isVirtualMeshSpec_CertificateAuthority_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *VirtualMeshSpec_CertificateAuthority) GetBuiltin() *VirtualMeshSpec_CertificateAuthority_Builtin {
	if x, ok := m.GetType().(*VirtualMeshSpec_CertificateAuthority_Builtin_); ok {
		return x.Builtin
	}
	return nil
}

func (m *VirtualMeshSpec_CertificateAuthority) GetProvided() *VirtualMeshSpec_CertificateAuthority_Provided {
	if x, ok := m.GetType().(*VirtualMeshSpec_CertificateAuthority_Provided_); ok {
		return x.Provided
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec_CertificateAuthority) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_CertificateAuthority_Builtin_)(nil),
		(*VirtualMeshSpec_CertificateAuthority_Provided_)(nil),
	}
}

//
//Configuration for auto-generated root certificate unique to the VirtualMesh
//Uses the X.509 format, RFC5280
type VirtualMeshSpec_CertificateAuthority_Builtin struct {
	// Number of days before root cert expires. Defaults to 365.
	TtlDays uint32 `protobuf:"varint,1,opt,name=ttl_days,json=ttlDays,proto3" json:"ttl_days,omitempty"`
	// Size in bytes of the root cert's private key. Defaults to 4096
	RsaKeySizeBytes uint32 `protobuf:"varint,2,opt,name=rsa_key_size_bytes,json=rsaKeySizeBytes,proto3" json:"rsa_key_size_bytes,omitempty"`
	// Root cert organization name. Defaults to "service-mesh-hub"
	OrgName              string   `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) Reset() {
	*m = VirtualMeshSpec_CertificateAuthority_Builtin{}
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshSpec_CertificateAuthority_Builtin) ProtoMessage() {}
func (*VirtualMeshSpec_CertificateAuthority_Builtin) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 0, 0}
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority_Builtin) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Builtin proto.InternalMessageInfo

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetTtlDays() uint32 {
	if m != nil {
		return m.TtlDays
	}
	return 0
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetRsaKeySizeBytes() uint32 {
	if m != nil {
		return m.RsaKeySizeBytes
	}
	return 0
}

func (m *VirtualMeshSpec_CertificateAuthority_Builtin) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

// Configuration for user-provided root certificate.
type VirtualMeshSpec_CertificateAuthority_Provided struct {
	// Reference to a Secret object containing the root certificate.
	Certificate          *v1.ClusterObjectRef `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VirtualMeshSpec_CertificateAuthority_Provided) Reset() {
	*m = VirtualMeshSpec_CertificateAuthority_Provided{}
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshSpec_CertificateAuthority_Provided) ProtoMessage() {}
func (*VirtualMeshSpec_CertificateAuthority_Provided) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 0, 1}
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Merge(m, src)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.Size(m)
}
func (m *VirtualMeshSpec_CertificateAuthority_Provided) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_CertificateAuthority_Provided proto.InternalMessageInfo

func (m *VirtualMeshSpec_CertificateAuthority_Provided) GetCertificate() *v1.ClusterObjectRef {
	if m != nil {
		return m.Certificate
	}
	return nil
}

// In Service Mesh Hub, Federation refers to the ability
// to expose services on with a global DNS name
// for traffic originating from any service within the
// virtual mesh.
type VirtualMeshSpec_Federation struct {
	// The "mode" in which to run federate services within this virtual mesh.
	//
	// Types that are valid to be assigned to Mode:
	//	*VirtualMeshSpec_Federation_Permissive
	Mode                 isVirtualMeshSpec_Federation_Mode `protobuf_oneof:"mode"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *VirtualMeshSpec_Federation) Reset()         { *m = VirtualMeshSpec_Federation{} }
func (m *VirtualMeshSpec_Federation) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_Federation) ProtoMessage()    {}
func (*VirtualMeshSpec_Federation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 1}
}
func (m *VirtualMeshSpec_Federation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_Federation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_Federation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_Federation.Merge(m, src)
}
func (m *VirtualMeshSpec_Federation) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_Federation.Size(m)
}
func (m *VirtualMeshSpec_Federation) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_Federation.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_Federation proto.InternalMessageInfo

type isVirtualMeshSpec_Federation_Mode interface {
	isVirtualMeshSpec_Federation_Mode()
	Equal(interface{}) bool
}

type VirtualMeshSpec_Federation_Permissive struct {
	Permissive *types.Empty `protobuf:"bytes,1,opt,name=permissive,proto3,oneof" json:"permissive,omitempty"`
}

func (*VirtualMeshSpec_Federation_Permissive) isVirtualMeshSpec_Federation_Mode() {}

func (m *VirtualMeshSpec_Federation) GetMode() isVirtualMeshSpec_Federation_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *VirtualMeshSpec_Federation) GetPermissive() *types.Empty {
	if x, ok := m.GetMode().(*VirtualMeshSpec_Federation_Permissive); ok {
		return x.Permissive
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*VirtualMeshSpec_Federation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*VirtualMeshSpec_Federation_Permissive)(nil),
	}
}

//
//Shared trust is a virtual mesh trust model requiring a shared root certificate, as well as shared identity
//between all entities which wish to communicate within the virtual mesh.
//
//The best current example of this would be the replicated control planes example from Istio:
//https://preliminary.istio.io/docs/setup/install/multicluster/gateways/
type VirtualMeshSpec_SharedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_SharedTrust) Reset()         { *m = VirtualMeshSpec_SharedTrust{} }
func (m *VirtualMeshSpec_SharedTrust) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_SharedTrust) ProtoMessage()    {}
func (*VirtualMeshSpec_SharedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 2}
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_SharedTrust.Merge(m, src)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_SharedTrust.Size(m)
}
func (m *VirtualMeshSpec_SharedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_SharedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_SharedTrust proto.InternalMessageInfo

//
//Limited trust is a virtual mesh trust model which does not require all meshes sharing the same root certificate
//or identity model. But rather, the limited trust creates trust between meshes running on different clusters
//by connecting their ingress/egress gateways with a common cert/identity. In this model all requests
//between different have the following request path when communicating between clusters
//
//cluster 1 MTLS               shared MTLS                  cluster 2 MTLS
//client/workload <-----------> egress gateway <----------> ingress gateway <--------------> server
//
//This approach has the downside of not maintaining identity from client to server, but allows for ad-hoc
//addition of additional clusters into a virtual mesh.
type VirtualMeshSpec_LimitedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshSpec_LimitedTrust) Reset()         { *m = VirtualMeshSpec_LimitedTrust{} }
func (m *VirtualMeshSpec_LimitedTrust) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshSpec_LimitedTrust) ProtoMessage()    {}
func (*VirtualMeshSpec_LimitedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{0, 3}
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Unmarshal(m, b)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Marshal(b, m, deterministic)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Merge(m, src)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshSpec_LimitedTrust.Size(m)
}
func (m *VirtualMeshSpec_LimitedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshSpec_LimitedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshSpec_LimitedTrust proto.InternalMessageInfo

type VirtualMeshStatus struct {
	// The most recent generation observed in the the TrafficPolicy metadata.
	// if the observedGeneration does not match generation, the controller has not received the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// the state of the overall resource.
	// will only show accepted if it has been successfully
	// applied to all target meshes.
	State ValidationState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.smh.solo.io.ValidationState" json:"state,omitempty"`
	// The status of the VirtualMesh for each Mesh to which it has been applied.
	// A TrafficPolicy may be Accepted for some Meshes and rejected for others.
	Meshes               map[string]*ValidationStatus `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *VirtualMeshStatus) Reset()         { *m = VirtualMeshStatus{} }
func (m *VirtualMeshStatus) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshStatus) ProtoMessage()    {}
func (*VirtualMeshStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8ec20aa69a8d951, []int{1}
}
func (m *VirtualMeshStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshStatus.Unmarshal(m, b)
}
func (m *VirtualMeshStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshStatus.Marshal(b, m, deterministic)
}
func (m *VirtualMeshStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshStatus.Merge(m, src)
}
func (m *VirtualMeshStatus) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshStatus.Size(m)
}
func (m *VirtualMeshStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshStatus proto.InternalMessageInfo

func (m *VirtualMeshStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *VirtualMeshStatus) GetState() ValidationState {
	if m != nil {
		return m.State
	}
	return ValidationState_PENDING
}

func (m *VirtualMeshStatus) GetMeshes() map[string]*ValidationStatus {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func init() {
	proto.RegisterEnum("networking.smh.solo.io.VirtualMeshSpec_EnforcementPolicy", VirtualMeshSpec_EnforcementPolicy_name, VirtualMeshSpec_EnforcementPolicy_value)
	proto.RegisterType((*VirtualMeshSpec)(nil), "networking.smh.solo.io.VirtualMeshSpec")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority)(nil), "networking.smh.solo.io.VirtualMeshSpec.CertificateAuthority")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority_Builtin)(nil), "networking.smh.solo.io.VirtualMeshSpec.CertificateAuthority.Builtin")
	proto.RegisterType((*VirtualMeshSpec_CertificateAuthority_Provided)(nil), "networking.smh.solo.io.VirtualMeshSpec.CertificateAuthority.Provided")
	proto.RegisterType((*VirtualMeshSpec_Federation)(nil), "networking.smh.solo.io.VirtualMeshSpec.Federation")
	proto.RegisterType((*VirtualMeshSpec_SharedTrust)(nil), "networking.smh.solo.io.VirtualMeshSpec.SharedTrust")
	proto.RegisterType((*VirtualMeshSpec_LimitedTrust)(nil), "networking.smh.solo.io.VirtualMeshSpec.LimitedTrust")
	proto.RegisterType((*VirtualMeshStatus)(nil), "networking.smh.solo.io.VirtualMeshStatus")
	proto.RegisterMapType((map[string]*ValidationStatus)(nil), "networking.smh.solo.io.VirtualMeshStatus.MeshesEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/virtual_mesh.proto", fileDescriptor_d8ec20aa69a8d951)
}

var fileDescriptor_d8ec20aa69a8d951 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6f, 0xdb, 0x36,
	0x10, 0xc7, 0x63, 0xbb, 0xb1, 0xd3, 0x73, 0x92, 0xa6, 0x5c, 0x16, 0x78, 0xde, 0x50, 0x64, 0xd9,
	0xc3, 0x0c, 0x0c, 0x91, 0x50, 0xb7, 0x03, 0xba, 0x61, 0xbf, 0xe2, 0x58, 0x5d, 0x86, 0x26, 0x59,
	0x2a, 0x77, 0x7b, 0xd8, 0x8b, 0x46, 0x4b, 0x67, 0x99, 0xb3, 0x24, 0x6a, 0x24, 0xe5, 0x42, 0xfd,
	0x83, 0x86, 0x3d, 0xec, 0x7d, 0xff, 0xcf, 0xfe, 0x92, 0x81, 0x94, 0xe4, 0x08, 0x99, 0x07, 0x18,
	0xc8, 0x93, 0x28, 0x92, 0xf7, 0x39, 0x1e, 0xef, 0xee, 0x4b, 0xb8, 0x0e, 0x99, 0x9a, 0x67, 0x53,
	0xcb, 0xe7, 0xb1, 0x2d, 0x79, 0xc4, 0x4f, 0x19, 0xb7, 0x25, 0x8a, 0x25, 0xf3, 0xf1, 0x34, 0x46,
	0x39, 0x3f, 0x9d, 0x67, 0x53, 0x9b, 0xa6, 0xcc, 0x4e, 0x50, 0xbd, 0xe5, 0x62, 0xc1, 0x92, 0xd0,
	0x5e, 0x3e, 0xa5, 0x51, 0x3a, 0xa7, 0x43, 0x7b, 0xc9, 0x84, 0xca, 0x68, 0xe4, 0xe9, 0x8d, 0x56,
	0x2a, 0xb8, 0xe2, 0xe4, 0xe8, 0x76, 0x9f, 0x25, 0xe3, 0xb9, 0xa5, 0x99, 0x16, 0xe3, 0x7d, 0x6b,
	0x9d, 0x9f, 0xc5, 0x72, 0x68, 0xd8, 0x3e, 0x17, 0x68, 0x2f, 0x9f, 0x9a, 0x6f, 0xc1, 0xe9, 0x7f,
	0xbb, 0xf9, 0x21, 0x68, 0xc4, 0x02, 0xaa, 0x18, 0x4f, 0x3c, 0xa9, 0xa8, 0xaa, 0x00, 0x1f, 0x86,
	0x9c, 0x87, 0x11, 0xda, 0xe6, 0x6f, 0x9a, 0xcd, 0x6c, 0x8c, 0x53, 0x95, 0x97, 0x8b, 0x4f, 0xee,
	0x2e, 0xbe, 0x15, 0x34, 0x4d, 0x51, 0xc8, 0x72, 0xfd, 0x30, 0xe4, 0x21, 0x37, 0x43, 0x5b, 0x8f,
	0x8a, 0xd9, 0x93, 0xbf, 0x1e, 0xc2, 0xa3, 0x9f, 0x8b, 0x90, 0xaf, 0x50, 0xce, 0x27, 0x29, 0xfa,
	0xe4, 0x63, 0xd8, 0x0d, 0x98, 0x4c, 0x23, 0x9a, 0x7b, 0x09, 0x8d, 0xb1, 0xd7, 0x38, 0x6e, 0x0c,
	0x1e, 0xba, 0xdd, 0x72, 0xee, 0x9a, 0xc6, 0x48, 0x9e, 0x43, 0x5b, 0x07, 0x81, 0xb2, 0xd7, 0x3c,
	0x6e, 0x0d, 0xba, 0xc3, 0x8f, 0x2c, 0x13, 0xa7, 0x8e, 0xbe, 0xba, 0x1e, 0xeb, 0xc7, 0xe9, 0x6f,
	0xe8, 0x2b, 0x17, 0x67, 0x6e, 0xb9, 0x97, 0xfc, 0x0e, 0xef, 0xfb, 0x28, 0x14, 0x9b, 0x31, 0x9f,
	0x2a, 0xf4, 0x68, 0xa6, 0xe6, 0x5c, 0x30, 0x95, 0xf7, 0x5a, 0xc7, 0x8d, 0x41, 0x77, 0xf8, 0x95,
	0xb5, 0xfe, 0xa2, 0xad, 0x3b, 0x07, 0xb4, 0xce, 0x6f, 0x21, 0x67, 0x15, 0xc3, 0x3d, 0xf4, 0xd7,
	0xcc, 0x12, 0x17, 0x60, 0x86, 0x01, 0x0a, 0x73, 0x99, 0xbd, 0x07, 0xc6, 0xcf, 0x70, 0x53, 0x3f,
	0x2f, 0x57, 0x96, 0x6e, 0x8d, 0x42, 0xae, 0xa0, 0x2d, 0xe7, 0x54, 0x60, 0xd0, 0xdb, 0x36, 0xbc,
	0x67, 0x9b, 0xf2, 0x26, 0xc6, 0xea, 0x8d, 0xc8, 0xa4, 0xba, 0xd8, 0x72, 0x4b, 0x08, 0xb9, 0x81,
	0x4e, 0xc4, 0x62, 0xa6, 0x30, 0xe8, 0xb5, 0x0d, 0xef, 0xf9, 0xa6, 0xbc, 0xcb, 0xc2, 0xac, 0x02,
	0x56, 0x18, 0xc2, 0xe1, 0x08, 0x93, 0x19, 0x17, 0x3e, 0x7a, 0xd4, 0xf7, 0x51, 0x4a, 0xcf, 0xe7,
	0x89, 0x12, 0x3c, 0xea, 0x75, 0x8e, 0x1b, 0x83, 0xfd, 0xe1, 0x17, 0x9b, 0x3a, 0x70, 0x0a, 0x4a,
	0x8c, 0x89, 0xba, 0xe1, 0x11, 0xf3, 0x73, 0xf7, 0xb0, 0x04, 0x9f, 0x19, 0xee, 0x79, 0x81, 0xed,
	0xff, 0xd1, 0x82, 0xc3, 0x75, 0x49, 0x21, 0xbf, 0x42, 0x67, 0x9a, 0xb1, 0x48, 0xb1, 0xc4, 0x54,
	0x51, 0x77, 0x38, 0xbe, 0x4f, 0x8e, 0xad, 0x51, 0xc1, 0xd2, 0xb1, 0x96, 0x58, 0xe2, 0xc3, 0x4e,
	0x2a, 0xf8, 0x92, 0x05, 0x18, 0xf4, 0x9a, 0xc6, 0x85, 0x73, 0x2f, 0x17, 0x37, 0x25, 0xec, 0x62,
	0xcb, 0x5d, 0x81, 0xfb, 0x11, 0x74, 0x4a, 0xd7, 0xe4, 0x03, 0xd8, 0x51, 0x2a, 0xf2, 0x02, 0x9a,
	0x4b, 0x13, 0xd2, 0x9e, 0xdb, 0x51, 0x2a, 0x1a, 0xd3, 0x5c, 0x92, 0xcf, 0x80, 0x08, 0x49, 0xbd,
	0x05, 0xe6, 0x9e, 0x64, 0xef, 0xd0, 0x9b, 0xe6, 0xca, 0x34, 0x88, 0xde, 0xf4, 0x48, 0x48, 0xfa,
	0x0a, 0xf3, 0x09, 0x7b, 0x87, 0x23, 0x3d, 0xad, 0x39, 0x5c, 0x84, 0x45, 0x83, 0xb5, 0x4c, 0x83,
	0x75, 0xb8, 0x08, 0x75, 0x73, 0xf5, 0x5f, 0xc3, 0x4e, 0x75, 0x0a, 0xe2, 0x40, 0xb7, 0x56, 0xd7,
	0x65, 0xa3, 0x7c, 0xb2, 0xa6, 0xdb, 0xce, 0xa3, 0x4c, 0x2a, 0x14, 0xb7, 0x4d, 0x57, 0xb7, 0x1b,
	0xb5, 0xe1, 0x81, 0xca, 0x53, 0xec, 0x5f, 0x03, 0xdc, 0x16, 0x35, 0x79, 0x01, 0x90, 0xa2, 0x88,
	0x99, 0x94, 0x6c, 0x89, 0x65, 0x82, 0x8e, 0xac, 0x42, 0x47, 0xac, 0x4a, 0x47, 0x2c, 0x47, 0x8b,
	0xcc, 0xc5, 0x96, 0x5b, 0xdb, 0xab, 0x79, 0x31, 0x0f, 0xb0, 0xbf, 0x07, 0xdd, 0x5a, 0x51, 0xf7,
	0xf7, 0x61, 0xb7, 0x5e, 0x93, 0x27, 0xdf, 0xc1, 0xe3, 0xff, 0x94, 0x10, 0x39, 0x80, 0xdd, 0x2b,
	0x67, 0x72, 0xe1, 0x8d, 0x9d, 0x97, 0x67, 0x3f, 0x5d, 0xbe, 0x39, 0xd8, 0x22, 0x5d, 0xe8, 0x38,
	0xd7, 0x67, 0xa3, 0x4b, 0x67, 0x7c, 0xd0, 0x20, 0xbb, 0xb0, 0x33, 0xfe, 0x61, 0x52, 0xfc, 0x35,
	0x47, 0x7b, 0xd0, 0x55, 0x1a, 0xe5, 0x69, 0x77, 0xd1, 0xc9, 0xdf, 0x4d, 0x78, 0x5c, 0x4f, 0xa3,
	0xa2, 0x2a, 0x93, 0xc4, 0x86, 0xf7, 0xf8, 0x54, 0x8b, 0x2b, 0x06, 0x5e, 0x88, 0x49, 0xd5, 0xed,
	0x3a, 0xa0, 0x96, 0x4b, 0xaa, 0xa5, 0xef, 0x57, 0x2b, 0xe4, 0x6b, 0xd8, 0x36, 0xba, 0x6a, 0x92,
	0xb3, 0x3f, 0xfc, 0xf4, 0x7f, 0x2b, 0x66, 0xa5, 0xc3, 0xda, 0x13, 0xba, 0x85, 0x95, 0x16, 0x80,
	0x52, 0xfd, 0x5a, 0x46, 0xfd, 0x3e, 0xdf, 0xa4, 0xe2, 0xcc, 0x51, 0xad, 0x2b, 0x63, 0xe7, 0x24,
	0x4a, 0xe4, 0x95, 0x2c, 0xf6, 0x7d, 0xe8, 0xd6, 0xa6, 0xc9, 0x01, 0xb4, 0x16, 0x98, 0x97, 0xaa,
	0xab, 0x87, 0xe4, 0x1b, 0xd8, 0x5e, 0xd2, 0x28, 0xc3, 0xb2, 0xc0, 0x07, 0x9b, 0x1d, 0x37, 0x93,
	0x6e, 0x61, 0xf6, 0x65, 0xf3, 0x45, 0x63, 0xf4, 0xfa, 0xcf, 0x7f, 0x9e, 0x34, 0x7e, 0x79, 0xb5,
	0xc9, 0xd3, 0x98, 0x2e, 0xc2, 0x3b, 0x2f, 0x53, 0xdd, 0xcb, 0xea, 0x95, 0x9a, 0xb6, 0x4d, 0x89,
	0x3c, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x7e, 0x74, 0x0f, 0x70, 0x07, 0x00, 0x00,
}

func (this *VirtualMeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec)
	if !ok {
		that2, ok := that.(VirtualMeshSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !this.CertificateAuthority.Equal(that1.CertificateAuthority) {
		return false
	}
	if !this.Federation.Equal(that1.Federation) {
		return false
	}
	if that1.TrustModel == nil {
		if this.TrustModel != nil {
			return false
		}
	} else if this.TrustModel == nil {
		return false
	} else if !this.TrustModel.Equal(that1.TrustModel) {
		return false
	}
	if this.EnforceAccessControl != that1.EnforceAccessControl {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Shared) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Shared)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Shared)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Shared.Equal(that1.Shared) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Limited) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Limited)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Limited)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Limited.Equal(that1.Limited) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Builtin_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Builtin_)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Builtin_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Builtin.Equal(that1.Builtin) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Provided_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Provided_)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Provided_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Provided.Equal(that1.Provided) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Builtin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Builtin)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Builtin)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TtlDays != that1.TtlDays {
		return false
	}
	if this.RsaKeySizeBytes != that1.RsaKeySizeBytes {
		return false
	}
	if this.OrgName != that1.OrgName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_CertificateAuthority_Provided) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_CertificateAuthority_Provided)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_CertificateAuthority_Provided)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Certificate.Equal(that1.Certificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Federation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Federation)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Federation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Mode == nil {
		if this.Mode != nil {
			return false
		}
	} else if this.Mode == nil {
		return false
	} else if !this.Mode.Equal(that1.Mode) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_Federation_Permissive) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_Federation_Permissive)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_Federation_Permissive)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Permissive.Equal(that1.Permissive) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_SharedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_SharedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_SharedTrust)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshSpec_LimitedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshSpec_LimitedTrust)
	if !ok {
		that2, ok := that.(VirtualMeshSpec_LimitedTrust)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshStatus)
	if !ok {
		that2, ok := that.(VirtualMeshStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
