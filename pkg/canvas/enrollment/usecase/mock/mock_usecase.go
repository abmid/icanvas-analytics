// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	entity "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEnrollmentUseCase is a mock of EnrollmentUseCase interface
type MockEnrollmentUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockEnrollmentUseCaseMockRecorder
}

// MockEnrollmentUseCaseMockRecorder is the mock recorder for MockEnrollmentUseCase
type MockEnrollmentUseCaseMockRecorder struct {
	mock *MockEnrollmentUseCase
}

// NewMockEnrollmentUseCase creates a new mock instance
func NewMockEnrollmentUseCase(ctrl *gomock.Controller) *MockEnrollmentUseCase {
	mock := &MockEnrollmentUseCase{ctrl: ctrl}
	mock.recorder = &MockEnrollmentUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnrollmentUseCase) EXPECT() *MockEnrollmentUseCaseMockRecorder {
	return m.recorder
}

// ListEnrollmentByCourseID mocks base method
func (m *MockEnrollmentUseCase) ListEnrollmentByCourseID(courseID uint32) ([]entity.Enrollment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnrollmentByCourseID", courseID)
	ret0, _ := ret[0].([]entity.Enrollment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnrollmentByCourseID indicates an expected call of ListEnrollmentByCourseID
func (mr *MockEnrollmentUseCaseMockRecorder) ListEnrollmentByCourseID(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnrollmentByCourseID", reflect.TypeOf((*MockEnrollmentUseCase)(nil).ListEnrollmentByCourseID), courseID)
}