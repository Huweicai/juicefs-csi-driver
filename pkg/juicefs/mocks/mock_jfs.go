// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juicedata/juicefs-csi-driver/pkg/juicefs (interfaces: Jfs)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJfs is a mock of Jfs interface.
type MockJfs struct {
	ctrl     *gomock.Controller
	recorder *MockJfsMockRecorder
}

// MockJfsMockRecorder is the mock recorder for MockJfs.
type MockJfsMockRecorder struct {
	mock *MockJfs
}

// NewMockJfs creates a new mock instance.
func NewMockJfs(ctrl *gomock.Controller) *MockJfs {
	mock := &MockJfs{ctrl: ctrl}
	mock.recorder = &MockJfsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJfs) EXPECT() *MockJfsMockRecorder {
	return m.recorder
}

// CreateVol mocks base method.
func (m *MockJfs) CreateVol(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVol", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVol indicates an expected call of CreateVol.
func (mr *MockJfsMockRecorder) CreateVol(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVol", reflect.TypeOf((*MockJfs)(nil).CreateVol), arg0, arg1)
}

// DeleteVol mocks base method.
func (m *MockJfs) DeleteVol(arg0 string, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVol", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVol indicates an expected call of DeleteVol.
func (mr *MockJfsMockRecorder) DeleteVol(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVol", reflect.TypeOf((*MockJfs)(nil).DeleteVol), arg0, arg1)
}

// GetBasePath mocks base method.
func (m *MockJfs) GetBasePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBasePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBasePath indicates an expected call of GetBasePath.
func (mr *MockJfsMockRecorder) GetBasePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBasePath", reflect.TypeOf((*MockJfs)(nil).GetBasePath))
}