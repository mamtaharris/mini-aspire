// Code generated by MockGen. DO NOT EDIT.
// Source: loan.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	requests "github.com/mamtaharris/mini-aspire/internal/models/requests"
)

// MockLoanReqValidatorInterface is a mock of LoanReqValidatorInterface interface.
type MockLoanReqValidatorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLoanReqValidatorInterfaceMockRecorder
}

// MockLoanReqValidatorInterfaceMockRecorder is the mock recorder for MockLoanReqValidatorInterface.
type MockLoanReqValidatorInterfaceMockRecorder struct {
	mock *MockLoanReqValidatorInterface
}

// NewMockLoanReqValidatorInterface creates a new mock instance.
func NewMockLoanReqValidatorInterface(ctrl *gomock.Controller) *MockLoanReqValidatorInterface {
	mock := &MockLoanReqValidatorInterface{ctrl: ctrl}
	mock.recorder = &MockLoanReqValidatorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoanReqValidatorInterface) EXPECT() *MockLoanReqValidatorInterfaceMockRecorder {
	return m.recorder
}

// ValidateCreateLoanReq mocks base method.
func (m *MockLoanReqValidatorInterface) ValidateCreateLoanReq(ctx *gin.Context) (requests.CreateLoanReq, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCreateLoanReq", ctx)
	ret0, _ := ret[0].(requests.CreateLoanReq)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateCreateLoanReq indicates an expected call of ValidateCreateLoanReq.
func (mr *MockLoanReqValidatorInterfaceMockRecorder) ValidateCreateLoanReq(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCreateLoanReq", reflect.TypeOf((*MockLoanReqValidatorInterface)(nil).ValidateCreateLoanReq), ctx)
}

// ValidateGetLoanReq mocks base method.
func (m *MockLoanReqValidatorInterface) ValidateGetLoanReq(ctx *gin.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGetLoanReq", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateGetLoanReq indicates an expected call of ValidateGetLoanReq.
func (mr *MockLoanReqValidatorInterfaceMockRecorder) ValidateGetLoanReq(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGetLoanReq", reflect.TypeOf((*MockLoanReqValidatorInterface)(nil).ValidateGetLoanReq), ctx)
}

// ValidateRepayLoanReq mocks base method.
func (m *MockLoanReqValidatorInterface) ValidateRepayLoanReq(ctx *gin.Context) (requests.RepayLoanReq, int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRepayLoanReq", ctx)
	ret0, _ := ret[0].(requests.RepayLoanReq)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ValidateRepayLoanReq indicates an expected call of ValidateRepayLoanReq.
func (mr *MockLoanReqValidatorInterfaceMockRecorder) ValidateRepayLoanReq(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRepayLoanReq", reflect.TypeOf((*MockLoanReqValidatorInterface)(nil).ValidateRepayLoanReq), ctx)
}

// ValidateUpdateLoanReq mocks base method.
func (m *MockLoanReqValidatorInterface) ValidateUpdateLoanReq(ctx *gin.Context) (requests.UpdateLoanReq, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUpdateLoanReq", ctx)
	ret0, _ := ret[0].(requests.UpdateLoanReq)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateUpdateLoanReq indicates an expected call of ValidateUpdateLoanReq.
func (mr *MockLoanReqValidatorInterfaceMockRecorder) ValidateUpdateLoanReq(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUpdateLoanReq", reflect.TypeOf((*MockLoanReqValidatorInterface)(nil).ValidateUpdateLoanReq), ctx)
}