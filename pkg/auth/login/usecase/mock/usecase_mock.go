// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	entity "github.com/abmid/icanvas-analytics/pkg/user/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLoginUseCase is a mock of LoginUseCase interface
type MockLoginUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockLoginUseCaseMockRecorder
}

// MockLoginUseCaseMockRecorder is the mock recorder for MockLoginUseCase
type MockLoginUseCaseMockRecorder struct {
	mock *MockLoginUseCase
}

// NewMockLoginUseCase creates a new mock instance
func NewMockLoginUseCase(ctrl *gomock.Controller) *MockLoginUseCase {
	mock := &MockLoginUseCase{ctrl: ctrl}
	mock.recorder = &MockLoginUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoginUseCase) EXPECT() *MockLoginUseCaseMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockLoginUseCase) Login(email, password string) (*entity.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login
func (mr *MockLoginUseCaseMockRecorder) Login(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginUseCase)(nil).Login), email, password)
}