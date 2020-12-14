// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/johnnywidth/9ty/api (interfaces: PortDomainClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/johnnywidth/9ty/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockPortDomainClient is a mock of PortDomainClient interface
type MockPortDomainClient struct {
	ctrl     *gomock.Controller
	recorder *MockPortDomainClientMockRecorder
}

// MockPortDomainClientMockRecorder is the mock recorder for MockPortDomainClient
type MockPortDomainClientMockRecorder struct {
	mock *MockPortDomainClient
}

// NewMockPortDomainClient creates a new mock instance
func NewMockPortDomainClient(ctrl *gomock.Controller) *MockPortDomainClient {
	mock := &MockPortDomainClient{ctrl: ctrl}
	mock.recorder = &MockPortDomainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortDomainClient) EXPECT() *MockPortDomainClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPortDomainClient) Create(arg0 context.Context, arg1 *api.PortRequest, arg2 ...grpc.CallOption) (*api.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*api.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPortDomainClientMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPortDomainClient)(nil).Create), varargs...)
}

// Get mocks base method
func (m *MockPortDomainClient) Get(arg0 context.Context, arg1 *api.GetRequest, arg2 ...grpc.CallOption) (*api.PortResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*api.PortResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPortDomainClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPortDomainClient)(nil).Get), varargs...)
}
