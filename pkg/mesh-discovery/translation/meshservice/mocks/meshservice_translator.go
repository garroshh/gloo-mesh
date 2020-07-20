// Code generated by MockGen. DO NOT EDIT.
// Source: ./meshservice_translator.go

// Package mock_meshservice is a generated GoMock package.
package mock_meshservice

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
)

// MockTranslator is a mock of Translator interface.
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator.
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance.
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// TranslateMeshServices mocks base method.
func (m *MockTranslator) TranslateMeshServices(services v1sets.ServiceSet, meshWorkloads v1alpha2sets.MeshWorkloadSet) v1alpha2sets.MeshServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateMeshServices", services, meshWorkloads)
	ret0, _ := ret[0].(v1alpha2sets.MeshServiceSet)
	return ret0
}

// TranslateMeshServices indicates an expected call of TranslateMeshServices.
func (mr *MockTranslatorMockRecorder) TranslateMeshServices(services, meshWorkloads interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateMeshServices", reflect.TypeOf((*MockTranslator)(nil).TranslateMeshServices), services, meshWorkloads)
}
