// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/rbac.enterprise.mesh.gloo.solo.io/v1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/rbac.enterprise.mesh.gloo.solo.io/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterRoleReconciler is a mock of MulticlusterRoleReconciler interface.
type MockMulticlusterRoleReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleReconcilerMockRecorder
}

// MockMulticlusterRoleReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleReconciler.
type MockMulticlusterRoleReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleReconciler
}

// NewMockMulticlusterRoleReconciler creates a new mock instance.
func NewMockMulticlusterRoleReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleReconciler {
	mock := &MockMulticlusterRoleReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRoleReconciler) EXPECT() *MockMulticlusterRoleReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRole mocks base method.
func (m *MockMulticlusterRoleReconciler) ReconcileRole(clusterName string, obj *v1.Role) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRole", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRole indicates an expected call of ReconcileRole.
func (mr *MockMulticlusterRoleReconcilerMockRecorder) ReconcileRole(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRole", reflect.TypeOf((*MockMulticlusterRoleReconciler)(nil).ReconcileRole), clusterName, obj)
}

// MockMulticlusterRoleDeletionReconciler is a mock of MulticlusterRoleDeletionReconciler interface.
type MockMulticlusterRoleDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleDeletionReconcilerMockRecorder
}

// MockMulticlusterRoleDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleDeletionReconciler.
type MockMulticlusterRoleDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleDeletionReconciler
}

// NewMockMulticlusterRoleDeletionReconciler creates a new mock instance.
func NewMockMulticlusterRoleDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleDeletionReconciler {
	mock := &MockMulticlusterRoleDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRoleDeletionReconciler) EXPECT() *MockMulticlusterRoleDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleDeletion mocks base method.
func (m *MockMulticlusterRoleDeletionReconciler) ReconcileRoleDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoleDeletion indicates an expected call of ReconcileRoleDeletion.
func (mr *MockMulticlusterRoleDeletionReconcilerMockRecorder) ReconcileRoleDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleDeletion", reflect.TypeOf((*MockMulticlusterRoleDeletionReconciler)(nil).ReconcileRoleDeletion), clusterName, req)
}

// MockMulticlusterRoleReconcileLoop is a mock of MulticlusterRoleReconcileLoop interface.
type MockMulticlusterRoleReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleReconcileLoopMockRecorder
}

// MockMulticlusterRoleReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRoleReconcileLoop.
type MockMulticlusterRoleReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRoleReconcileLoop
}

// NewMockMulticlusterRoleReconcileLoop creates a new mock instance.
func NewMockMulticlusterRoleReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRoleReconcileLoop {
	mock := &MockMulticlusterRoleReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRoleReconcileLoop) EXPECT() *MockMulticlusterRoleReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRoleReconciler mocks base method.
func (m *MockMulticlusterRoleReconcileLoop) AddMulticlusterRoleReconciler(ctx context.Context, rec controller.MulticlusterRoleReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRoleReconciler", varargs...)
}

// AddMulticlusterRoleReconciler indicates an expected call of AddMulticlusterRoleReconciler.
func (mr *MockMulticlusterRoleReconcileLoopMockRecorder) AddMulticlusterRoleReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRoleReconciler", reflect.TypeOf((*MockMulticlusterRoleReconcileLoop)(nil).AddMulticlusterRoleReconciler), varargs...)
}

// MockMulticlusterRoleBindingReconciler is a mock of MulticlusterRoleBindingReconciler interface.
type MockMulticlusterRoleBindingReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleBindingReconcilerMockRecorder
}

// MockMulticlusterRoleBindingReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleBindingReconciler.
type MockMulticlusterRoleBindingReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleBindingReconciler
}

// NewMockMulticlusterRoleBindingReconciler creates a new mock instance.
func NewMockMulticlusterRoleBindingReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleBindingReconciler {
	mock := &MockMulticlusterRoleBindingReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleBindingReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRoleBindingReconciler) EXPECT() *MockMulticlusterRoleBindingReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleBinding mocks base method.
func (m *MockMulticlusterRoleBindingReconciler) ReconcileRoleBinding(clusterName string, obj *v1.RoleBinding) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBinding", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileRoleBinding indicates an expected call of ReconcileRoleBinding.
func (mr *MockMulticlusterRoleBindingReconcilerMockRecorder) ReconcileRoleBinding(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBinding", reflect.TypeOf((*MockMulticlusterRoleBindingReconciler)(nil).ReconcileRoleBinding), clusterName, obj)
}

// MockMulticlusterRoleBindingDeletionReconciler is a mock of MulticlusterRoleBindingDeletionReconciler interface.
type MockMulticlusterRoleBindingDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleBindingDeletionReconcilerMockRecorder
}

// MockMulticlusterRoleBindingDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterRoleBindingDeletionReconciler.
type MockMulticlusterRoleBindingDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterRoleBindingDeletionReconciler
}

// NewMockMulticlusterRoleBindingDeletionReconciler creates a new mock instance.
func NewMockMulticlusterRoleBindingDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterRoleBindingDeletionReconciler {
	mock := &MockMulticlusterRoleBindingDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleBindingDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRoleBindingDeletionReconciler) EXPECT() *MockMulticlusterRoleBindingDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileRoleBindingDeletion mocks base method.
func (m *MockMulticlusterRoleBindingDeletionReconciler) ReconcileRoleBindingDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileRoleBindingDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileRoleBindingDeletion indicates an expected call of ReconcileRoleBindingDeletion.
func (mr *MockMulticlusterRoleBindingDeletionReconcilerMockRecorder) ReconcileRoleBindingDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileRoleBindingDeletion", reflect.TypeOf((*MockMulticlusterRoleBindingDeletionReconciler)(nil).ReconcileRoleBindingDeletion), clusterName, req)
}

// MockMulticlusterRoleBindingReconcileLoop is a mock of MulticlusterRoleBindingReconcileLoop interface.
type MockMulticlusterRoleBindingReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterRoleBindingReconcileLoopMockRecorder
}

// MockMulticlusterRoleBindingReconcileLoopMockRecorder is the mock recorder for MockMulticlusterRoleBindingReconcileLoop.
type MockMulticlusterRoleBindingReconcileLoopMockRecorder struct {
	mock *MockMulticlusterRoleBindingReconcileLoop
}

// NewMockMulticlusterRoleBindingReconcileLoop creates a new mock instance.
func NewMockMulticlusterRoleBindingReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterRoleBindingReconcileLoop {
	mock := &MockMulticlusterRoleBindingReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterRoleBindingReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterRoleBindingReconcileLoop) EXPECT() *MockMulticlusterRoleBindingReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterRoleBindingReconciler mocks base method.
func (m *MockMulticlusterRoleBindingReconcileLoop) AddMulticlusterRoleBindingReconciler(ctx context.Context, rec controller.MulticlusterRoleBindingReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterRoleBindingReconciler", varargs...)
}

// AddMulticlusterRoleBindingReconciler indicates an expected call of AddMulticlusterRoleBindingReconciler.
func (mr *MockMulticlusterRoleBindingReconcileLoopMockRecorder) AddMulticlusterRoleBindingReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterRoleBindingReconciler", reflect.TypeOf((*MockMulticlusterRoleBindingReconcileLoop)(nil).AddMulticlusterRoleBindingReconciler), varargs...)
}
