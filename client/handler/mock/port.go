// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/johnnywidth/9ty/client/handler (interfaces: PortUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/johnnywidth/9ty/client/entity"
	reflect "reflect"
)

// MockPortUsecase is a mock of PortUsecase interface
type MockPortUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPortUsecaseMockRecorder
}

// MockPortUsecaseMockRecorder is the mock recorder for MockPortUsecase
type MockPortUsecaseMockRecorder struct {
	mock *MockPortUsecase
}

// NewMockPortUsecase creates a new mock instance
func NewMockPortUsecase(ctrl *gomock.Controller) *MockPortUsecase {
	mock := &MockPortUsecase{ctrl: ctrl}
	mock.recorder = &MockPortUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPortUsecase) EXPECT() *MockPortUsecaseMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPortUsecase) Get(arg0 context.Context, arg1 string) (*entity.PortData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*entity.PortData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPortUsecaseMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPortUsecase)(nil).Get), arg0, arg1)
}