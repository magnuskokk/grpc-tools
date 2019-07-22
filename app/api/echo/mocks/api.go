// Code generated by MockGen. DO NOT EDIT.
// Source: app/generated/idl/echo (interfaces: EchoAPIServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	echo "app/generated/idl/echo"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEchoAPIServer is a mock of EchoAPIServer interface
type MockEchoAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockEchoAPIServerMockRecorder
}

// MockEchoAPIServerMockRecorder is the mock recorder for MockEchoAPIServer
type MockEchoAPIServerMockRecorder struct {
	mock *MockEchoAPIServer
}

// NewMockEchoAPIServer creates a new mock instance
func NewMockEchoAPIServer(ctrl *gomock.Controller) *MockEchoAPIServer {
	mock := &MockEchoAPIServer{ctrl: ctrl}
	mock.recorder = &MockEchoAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEchoAPIServer) EXPECT() *MockEchoAPIServerMockRecorder {
	return m.recorder
}

// Echo mocks base method
func (m *MockEchoAPIServer) Echo(arg0 context.Context, arg1 *echo.EchoRequest) (*echo.EchoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo", arg0, arg1)
	ret0, _ := ret[0].(*echo.EchoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Echo indicates an expected call of Echo
func (mr *MockEchoAPIServerMockRecorder) Echo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockEchoAPIServer)(nil).Echo), arg0, arg1)
}
