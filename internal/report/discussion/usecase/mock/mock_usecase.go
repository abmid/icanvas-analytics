// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	entity "github.com/abmid/icanvas-analytics/internal/report/entity"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReportDiscussionUseCase is a mock of ReportDiscussionUseCase interface
type MockReportDiscussionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockReportDiscussionUseCaseMockRecorder
}

// MockReportDiscussionUseCaseMockRecorder is the mock recorder for MockReportDiscussionUseCase
type MockReportDiscussionUseCaseMockRecorder struct {
	mock *MockReportDiscussionUseCase
}

// NewMockReportDiscussionUseCase creates a new mock instance
func NewMockReportDiscussionUseCase(ctrl *gomock.Controller) *MockReportDiscussionUseCase {
	mock := &MockReportDiscussionUseCase{ctrl: ctrl}
	mock.recorder = &MockReportDiscussionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReportDiscussionUseCase) EXPECT() *MockReportDiscussionUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockReportDiscussionUseCase) Create(ctx context.Context, reportDiss *entity.ReportDiscussion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, reportDiss)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockReportDiscussionUseCaseMockRecorder) Create(ctx, reportDiss interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockReportDiscussionUseCase)(nil).Create), ctx, reportDiss)
}

// Read mocks base method
func (m *MockReportDiscussionUseCase) Read(ctx context.Context) ([]entity.ReportDiscussion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].([]entity.ReportDiscussion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReportDiscussionUseCaseMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReportDiscussionUseCase)(nil).Read), ctx)
}

// Update mocks base method
func (m *MockReportDiscussionUseCase) Update(ctx context.Context, reportDiss *entity.ReportDiscussion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, reportDiss)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockReportDiscussionUseCaseMockRecorder) Update(ctx, reportDiss interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockReportDiscussionUseCase)(nil).Update), ctx, reportDiss)
}

// Delete mocks base method
func (m *MockReportDiscussionUseCase) Delete(ctx context.Context, reportDissID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, reportDissID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockReportDiscussionUseCaseMockRecorder) Delete(ctx, reportDissID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReportDiscussionUseCase)(nil).Delete), ctx, reportDissID)
}

// CreateOrUpdateByFilter mocks base method
func (m *MockReportDiscussionUseCase) CreateOrUpdateByFilter(ctx context.Context, filter entity.ReportDiscussion, assigment *entity.ReportDiscussion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateByFilter", ctx, filter, assigment)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateByFilter indicates an expected call of CreateOrUpdateByFilter
func (mr *MockReportDiscussionUseCaseMockRecorder) CreateOrUpdateByFilter(ctx, filter, assigment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateByFilter", reflect.TypeOf((*MockReportDiscussionUseCase)(nil).CreateOrUpdateByFilter), ctx, filter, assigment)
}

// FindFilter mocks base method
func (m *MockReportDiscussionUseCase) FindFilter(ctx context.Context, filter entity.ReportDiscussion) ([]entity.ReportDiscussion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilter", ctx, filter)
	ret0, _ := ret[0].([]entity.ReportDiscussion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFilter indicates an expected call of FindFilter
func (mr *MockReportDiscussionUseCaseMockRecorder) FindFilter(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilter", reflect.TypeOf((*MockReportDiscussionUseCase)(nil).FindFilter), ctx, filter)
}