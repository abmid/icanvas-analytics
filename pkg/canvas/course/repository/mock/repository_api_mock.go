// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "github.com/abmid/icanvas-analytics/pkg/canvas/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCourseRepository is a mock of CourseRepository interface
type MockCourseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCourseRepositoryMockRecorder
}

// MockCourseRepositoryMockRecorder is the mock recorder for MockCourseRepository
type MockCourseRepositoryMockRecorder struct {
	mock *MockCourseRepository
}

// NewMockCourseRepository creates a new mock instance
func NewMockCourseRepository(ctrl *gomock.Controller) *MockCourseRepository {
	mock := &MockCourseRepository{ctrl: ctrl}
	mock.recorder = &MockCourseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCourseRepository) EXPECT() *MockCourseRepositoryMockRecorder {
	return m.recorder
}

// Courses mocks base method
func (m *MockCourseRepository) Courses(accountId, page uint32) ([]entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Courses", accountId, page)
	ret0, _ := ret[0].([]entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Courses indicates an expected call of Courses
func (mr *MockCourseRepositoryMockRecorder) Courses(accountId, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Courses", reflect.TypeOf((*MockCourseRepository)(nil).Courses), accountId, page)
}

// ListUserInCourse mocks base method
func (m *MockCourseRepository) ListUserInCourse(courseID uint32, enrollmentRole string) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserInCourse", courseID, enrollmentRole)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserInCourse indicates an expected call of ListUserInCourse
func (mr *MockCourseRepositoryMockRecorder) ListUserInCourse(courseID, enrollmentRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserInCourse", reflect.TypeOf((*MockCourseRepository)(nil).ListUserInCourse), courseID, enrollmentRole)
}
