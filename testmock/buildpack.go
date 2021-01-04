// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: Buildpack)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	lifecycle "github.com/buildpacks/lifecycle"
)

// MockBuildpack is a mock of Buildpack interface
type MockBuildpack struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackMockRecorder
}

// MockBuildpackMockRecorder is the mock recorder for MockBuildpack
type MockBuildpackMockRecorder struct {
	mock *MockBuildpack
}

// NewMockBuildpack creates a new mock instance
func NewMockBuildpack(ctrl *gomock.Controller) *MockBuildpack {
	mock := &MockBuildpack{ctrl: ctrl}
	mock.recorder = &MockBuildpackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuildpack) EXPECT() *MockBuildpackMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockBuildpack) Build(arg0 lifecycle.BuildpackPlan, arg1 lifecycle.BuildConfig) (lifecycle.BuildResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1)
	ret0, _ := ret[0].(lifecycle.BuildResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockBuildpackMockRecorder) Build(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuildpack)(nil).Build), arg0, arg1)
}