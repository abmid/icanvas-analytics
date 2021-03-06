// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	entity "github.com/abmid/icanvas-analytics/pkg/user/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserUseCase is a mock of UserUseCase interface
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserUseCase) Create(user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserUseCaseMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserUseCase)(nil).Create), user)
}

// Find mocks base method
func (m *MockUserUseCase) Find(email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserUseCaseMockRecorder) Find(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserUseCase)(nil).Find), email)
}

// All mocks base method
func (m *MockUserUseCase) All() ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockUserUseCaseMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockUserUseCase)(nil).All))
}
