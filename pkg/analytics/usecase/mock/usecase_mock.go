// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	pagination "github.com/abmid/icanvas-analytics/internal/pagination"
	entity "github.com/abmid/icanvas-analytics/pkg/analytics/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAnalyticsUseCase is a mock of AnalyticsUseCase interface
type MockAnalyticsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticsUseCaseMockRecorder
}

// MockAnalyticsUseCaseMockRecorder is the mock recorder for MockAnalyticsUseCase
type MockAnalyticsUseCaseMockRecorder struct {
	mock *MockAnalyticsUseCase
}

// NewMockAnalyticsUseCase creates a new mock instance
func NewMockAnalyticsUseCase(ctrl *gomock.Controller) *MockAnalyticsUseCase {
	mock := &MockAnalyticsUseCase{ctrl: ctrl}
	mock.recorder = &MockAnalyticsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnalyticsUseCase) EXPECT() *MockAnalyticsUseCaseMockRecorder {
	return m.recorder
}

// FindBestCourseByFilter mocks base method
func (m *MockAnalyticsUseCase) FindBestCourseByFilter(ctx context.Context, filter entity.FilterAnalytics) ([]entity.AnalyticsCourse, pagination.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBestCourseByFilter", ctx, filter)
	ret0, _ := ret[0].([]entity.AnalyticsCourse)
	ret1, _ := ret[1].(pagination.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindBestCourseByFilter indicates an expected call of FindBestCourseByFilter
func (mr *MockAnalyticsUseCaseMockRecorder) FindBestCourseByFilter(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBestCourseByFilter", reflect.TypeOf((*MockAnalyticsUseCase)(nil).FindBestCourseByFilter), ctx, filter)
}
