// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/celestiaorg/celestia-node/nodebuilder/header (interfaces: Module)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	header "github.com/celestiaorg/celestia-node/header"
	header0 "github.com/celestiaorg/celestia-node/libs/header"
	gomock "github.com/golang/mock/gomock"
)

// MockModule is a mock of Module interface.
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule.
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance.
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockModule) Get(arg0 context.Context, arg1 header0.Hash) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockModuleMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockModule)(nil).Get), arg0, arg1)
}

// GetByHeight mocks base method.
func (m *MockModule) GetByHeight(arg0 context.Context, arg1 uint64) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHeight", arg0, arg1)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHeight indicates an expected call of GetByHeight.
func (mr *MockModuleMockRecorder) GetByHeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHeight", reflect.TypeOf((*MockModule)(nil).GetByHeight), arg0, arg1)
}

// GetVerifiedRangeByHeight mocks base method.
func (m *MockModule) GetVerifiedRangeByHeight(arg0 context.Context, arg1 *header.ExtendedHeader, arg2 uint64) ([]*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVerifiedRangeByHeight", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVerifiedRangeByHeight indicates an expected call of GetVerifiedRangeByHeight.
func (mr *MockModuleMockRecorder) GetVerifiedRangeByHeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVerifiedRangeByHeight", reflect.TypeOf((*MockModule)(nil).GetVerifiedRangeByHeight), arg0, arg1, arg2)
}

// Head mocks base method.
func (m *MockModule) Head(arg0 context.Context) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head", arg0)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockModuleMockRecorder) Head(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockModule)(nil).Head), arg0)
}

// IsSyncing mocks base method.
func (m *MockModule) IsSyncing(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSyncing", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSyncing indicates an expected call of IsSyncing.
func (mr *MockModuleMockRecorder) IsSyncing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSyncing", reflect.TypeOf((*MockModule)(nil).IsSyncing), arg0)
}

// SyncHead mocks base method.
func (m *MockModule) SyncHead(arg0 context.Context) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncHead", arg0)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncHead indicates an expected call of SyncHead.
func (mr *MockModuleMockRecorder) SyncHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncHead", reflect.TypeOf((*MockModule)(nil).SyncHead), arg0)
}

// WaitSync mocks base method.
func (m *MockModule) WaitSync(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitSync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitSync indicates an expected call of WaitSync.
func (mr *MockModuleMockRecorder) WaitSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitSync", reflect.TypeOf((*MockModule)(nil).WaitSync), arg0)
}
