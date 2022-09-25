// Code generated by MockGen. DO NOT EDIT.
// Source: ./downstream.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/kodefluence/altair/plugin/oauth/entity"
	db "github.com/kodefluence/monorepo/db"
	exception "github.com/kodefluence/monorepo/exception"
	kontext "github.com/kodefluence/monorepo/kontext"
)

// MockOauthApplicationRepository is a mock of OauthApplicationRepository interface.
type MockOauthApplicationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationRepositoryMockRecorder
}

// MockOauthApplicationRepositoryMockRecorder is the mock recorder for MockOauthApplicationRepository.
type MockOauthApplicationRepositoryMockRecorder struct {
	mock *MockOauthApplicationRepository
}

// NewMockOauthApplicationRepository creates a new mock instance.
func NewMockOauthApplicationRepository(ctrl *gomock.Controller) *MockOauthApplicationRepository {
	mock := &MockOauthApplicationRepository{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthApplicationRepository) EXPECT() *MockOauthApplicationRepositoryMockRecorder {
	return m.recorder
}

// OneByUIDandSecret mocks base method.
func (m *MockOauthApplicationRepository) OneByUIDandSecret(ktx kontext.Context, clientUID, clientSecret string, tx db.TX) (entity.OauthApplication, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByUIDandSecret", ktx, clientUID, clientSecret, tx)
	ret0, _ := ret[0].(entity.OauthApplication)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// OneByUIDandSecret indicates an expected call of OneByUIDandSecret.
func (mr *MockOauthApplicationRepositoryMockRecorder) OneByUIDandSecret(ktx, clientUID, clientSecret, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByUIDandSecret", reflect.TypeOf((*MockOauthApplicationRepository)(nil).OneByUIDandSecret), ktx, clientUID, clientSecret, tx)
}

// MockRouterPath is a mock of RouterPath interface.
type MockRouterPath struct {
	ctrl     *gomock.Controller
	recorder *MockRouterPathMockRecorder
}

// MockRouterPathMockRecorder is the mock recorder for MockRouterPath.
type MockRouterPathMockRecorder struct {
	mock *MockRouterPath
}

// NewMockRouterPath creates a new mock instance.
func NewMockRouterPath(ctrl *gomock.Controller) *MockRouterPath {
	mock := &MockRouterPath{ctrl: ctrl}
	mock.recorder = &MockRouterPathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterPath) EXPECT() *MockRouterPathMockRecorder {
	return m.recorder
}

// GetAuth mocks base method.
func (m *MockRouterPath) GetAuth() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuth")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAuth indicates an expected call of GetAuth.
func (mr *MockRouterPathMockRecorder) GetAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*MockRouterPath)(nil).GetAuth))
}

// GetSCope mocks base method.
func (m *MockRouterPath) GetSCope() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSCope")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSCope indicates an expected call of GetSCope.
func (mr *MockRouterPathMockRecorder) GetSCope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSCope", reflect.TypeOf((*MockRouterPath)(nil).GetSCope))
}
