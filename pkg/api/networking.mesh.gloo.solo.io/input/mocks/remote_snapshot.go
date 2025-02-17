// Code generated by MockGen. DO NOT EDIT.
// Source: ./remote_snapshot.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"
	v1sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1/sets"
	input "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	v1beta1sets0 "github.com/solo-io/gloo-mesh/pkg/api/xds.agent.enterprise.mesh.gloo.solo.io/v1beta1/sets"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockRemoteSnapshot is a mock of RemoteSnapshot interface.
type MockRemoteSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteSnapshotMockRecorder
}

// MockRemoteSnapshotMockRecorder is the mock recorder for MockRemoteSnapshot.
type MockRemoteSnapshotMockRecorder struct {
	mock *MockRemoteSnapshot
}

// NewMockRemoteSnapshot creates a new mock instance.
func NewMockRemoteSnapshot(ctrl *gomock.Controller) *MockRemoteSnapshot {
	mock := &MockRemoteSnapshot{ctrl: ctrl}
	mock.recorder = &MockRemoteSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteSnapshot) EXPECT() *MockRemoteSnapshotMockRecorder {
	return m.recorder
}

// AuthorizationPolicies mocks base method.
func (m *MockRemoteSnapshot) AuthorizationPolicies() v1beta1sets.AuthorizationPolicySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationPolicies")
	ret0, _ := ret[0].(v1beta1sets.AuthorizationPolicySet)
	return ret0
}

// AuthorizationPolicies indicates an expected call of AuthorizationPolicies.
func (mr *MockRemoteSnapshotMockRecorder) AuthorizationPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationPolicies", reflect.TypeOf((*MockRemoteSnapshot)(nil).AuthorizationPolicies))
}

// DestinationRules mocks base method.
func (m *MockRemoteSnapshot) DestinationRules() v1alpha3sets.DestinationRuleSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestinationRules")
	ret0, _ := ret[0].(v1alpha3sets.DestinationRuleSet)
	return ret0
}

// DestinationRules indicates an expected call of DestinationRules.
func (mr *MockRemoteSnapshotMockRecorder) DestinationRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestinationRules", reflect.TypeOf((*MockRemoteSnapshot)(nil).DestinationRules))
}

// EnvoyFilters mocks base method.
func (m *MockRemoteSnapshot) EnvoyFilters() v1alpha3sets.EnvoyFilterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvoyFilters")
	ret0, _ := ret[0].(v1alpha3sets.EnvoyFilterSet)
	return ret0
}

// EnvoyFilters indicates an expected call of EnvoyFilters.
func (mr *MockRemoteSnapshotMockRecorder) EnvoyFilters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvoyFilters", reflect.TypeOf((*MockRemoteSnapshot)(nil).EnvoyFilters))
}

// Gateways mocks base method.
func (m *MockRemoteSnapshot) Gateways() v1alpha3sets.GatewaySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gateways")
	ret0, _ := ret[0].(v1alpha3sets.GatewaySet)
	return ret0
}

// Gateways indicates an expected call of Gateways.
func (mr *MockRemoteSnapshotMockRecorder) Gateways() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gateways", reflect.TypeOf((*MockRemoteSnapshot)(nil).Gateways))
}

// IssuedCertificates mocks base method.
func (m *MockRemoteSnapshot) IssuedCertificates() v1sets.IssuedCertificateSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssuedCertificates")
	ret0, _ := ret[0].(v1sets.IssuedCertificateSet)
	return ret0
}

// IssuedCertificates indicates an expected call of IssuedCertificates.
func (mr *MockRemoteSnapshotMockRecorder) IssuedCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssuedCertificates", reflect.TypeOf((*MockRemoteSnapshot)(nil).IssuedCertificates))
}

// MarshalJSON mocks base method.
func (m *MockRemoteSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON.
func (mr *MockRemoteSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockRemoteSnapshot)(nil).MarshalJSON))
}

// PodBounceDirectives mocks base method.
func (m *MockRemoteSnapshot) PodBounceDirectives() v1sets.PodBounceDirectiveSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodBounceDirectives")
	ret0, _ := ret[0].(v1sets.PodBounceDirectiveSet)
	return ret0
}

// PodBounceDirectives indicates an expected call of PodBounceDirectives.
func (mr *MockRemoteSnapshotMockRecorder) PodBounceDirectives() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodBounceDirectives", reflect.TypeOf((*MockRemoteSnapshot)(nil).PodBounceDirectives))
}

// ServiceEntries mocks base method.
func (m *MockRemoteSnapshot) ServiceEntries() v1alpha3sets.ServiceEntrySet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceEntries")
	ret0, _ := ret[0].(v1alpha3sets.ServiceEntrySet)
	return ret0
}

// ServiceEntries indicates an expected call of ServiceEntries.
func (mr *MockRemoteSnapshotMockRecorder) ServiceEntries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceEntries", reflect.TypeOf((*MockRemoteSnapshot)(nil).ServiceEntries))
}

// Sidecars mocks base method.
func (m *MockRemoteSnapshot) Sidecars() v1alpha3sets.SidecarSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sidecars")
	ret0, _ := ret[0].(v1alpha3sets.SidecarSet)
	return ret0
}

// Sidecars indicates an expected call of Sidecars.
func (mr *MockRemoteSnapshotMockRecorder) Sidecars() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sidecars", reflect.TypeOf((*MockRemoteSnapshot)(nil).Sidecars))
}

// SyncStatuses mocks base method.
func (m *MockRemoteSnapshot) SyncStatuses(ctx context.Context, c client.Client, opts input.RemoteSyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatuses", ctx, c, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatuses indicates an expected call of SyncStatuses.
func (mr *MockRemoteSnapshotMockRecorder) SyncStatuses(ctx, c, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatuses", reflect.TypeOf((*MockRemoteSnapshot)(nil).SyncStatuses), ctx, c, opts)
}

// SyncStatusesMultiCluster mocks base method.
func (m *MockRemoteSnapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts input.RemoteSyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusesMultiCluster", ctx, mcClient, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatusesMultiCluster indicates an expected call of SyncStatusesMultiCluster.
func (mr *MockRemoteSnapshotMockRecorder) SyncStatusesMultiCluster(ctx, mcClient, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusesMultiCluster", reflect.TypeOf((*MockRemoteSnapshot)(nil).SyncStatusesMultiCluster), ctx, mcClient, opts)
}

// VirtualServices mocks base method.
func (m *MockRemoteSnapshot) VirtualServices() v1alpha3sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualServices")
	ret0, _ := ret[0].(v1alpha3sets.VirtualServiceSet)
	return ret0
}

// VirtualServices indicates an expected call of VirtualServices.
func (mr *MockRemoteSnapshotMockRecorder) VirtualServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualServices", reflect.TypeOf((*MockRemoteSnapshot)(nil).VirtualServices))
}

// XdsConfigs mocks base method.
func (m *MockRemoteSnapshot) XdsConfigs() v1beta1sets0.XdsConfigSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XdsConfigs")
	ret0, _ := ret[0].(v1beta1sets0.XdsConfigSet)
	return ret0
}

// XdsConfigs indicates an expected call of XdsConfigs.
func (mr *MockRemoteSnapshotMockRecorder) XdsConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XdsConfigs", reflect.TypeOf((*MockRemoteSnapshot)(nil).XdsConfigs))
}

// MockRemoteBuilder is a mock of RemoteBuilder interface.
type MockRemoteBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteBuilderMockRecorder
}

// MockRemoteBuilderMockRecorder is the mock recorder for MockRemoteBuilder.
type MockRemoteBuilderMockRecorder struct {
	mock *MockRemoteBuilder
}

// NewMockRemoteBuilder creates a new mock instance.
func NewMockRemoteBuilder(ctrl *gomock.Controller) *MockRemoteBuilder {
	mock := &MockRemoteBuilder{ctrl: ctrl}
	mock.recorder = &MockRemoteBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteBuilder) EXPECT() *MockRemoteBuilderMockRecorder {
	return m.recorder
}

// BuildSnapshot mocks base method.
func (m *MockRemoteBuilder) BuildSnapshot(ctx context.Context, name string, opts input.RemoteBuildOptions) (input.RemoteSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSnapshot", ctx, name, opts)
	ret0, _ := ret[0].(input.RemoteSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSnapshot indicates an expected call of BuildSnapshot.
func (mr *MockRemoteBuilderMockRecorder) BuildSnapshot(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSnapshot", reflect.TypeOf((*MockRemoteBuilder)(nil).BuildSnapshot), ctx, name, opts)
}
