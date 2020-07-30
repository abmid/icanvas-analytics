// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "github.com/abmid/icanvas-analytics/pkg/setting/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSettingRepository is a mock of SettingRepository interface
type MockSettingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSettingRepositoryMockRecorder
}

// MockSettingRepositoryMockRecorder is the mock recorder for MockSettingRepository
type MockSettingRepositoryMockRecorder struct {
	mock *MockSettingRepository
}

// NewMockSettingRepository creates a new mock instance
func NewMockSettingRepository(ctrl *gomock.Controller) *MockSettingRepository {
	mock := &MockSettingRepository{ctrl: ctrl}
	mock.recorder = &MockSettingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingRepository) EXPECT() *MockSettingRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockSettingRepository) Create(setting *entity.Setting) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", setting)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSettingRepositoryMockRecorder) Create(setting interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSettingRepository)(nil).Create), setting)
}

// Update mocks base method
func (m *MockSettingRepository) Update(id uint32, setting entity.Setting) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, setting)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSettingRepositoryMockRecorder) Update(id, setting interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSettingRepository)(nil).Update), id, setting)
}

// FindByFilter mocks base method
func (m *MockSettingRepository) FindByFilter(filter entity.Setting) ([]entity.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByFilter", filter)
	ret0, _ := ret[0].([]entity.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByFilter indicates an expected call of FindByFilter
func (mr *MockSettingRepositoryMockRecorder) FindByFilter(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByFilter", reflect.TypeOf((*MockSettingRepository)(nil).FindByFilter), filter)
}
