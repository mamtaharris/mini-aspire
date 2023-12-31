// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	requests "github.com/mamtaharris/mini-aspire/internal/models/requests"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ValidateUserAndGenerateToken mocks base method.
func (m *MockUserService) ValidateUserAndGenerateToken(ctx context.Context, loginReq requests.UserLoginReq) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUserAndGenerateToken", ctx, loginReq)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUserAndGenerateToken indicates an expected call of ValidateUserAndGenerateToken.
func (mr *MockUserServiceMockRecorder) ValidateUserAndGenerateToken(ctx, loginReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUserAndGenerateToken", reflect.TypeOf((*MockUserService)(nil).ValidateUserAndGenerateToken), ctx, loginReq)
}
