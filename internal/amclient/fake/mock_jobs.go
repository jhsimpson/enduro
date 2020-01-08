// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-labs/enduro/internal/amclient (interfaces: JobsService)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	amclient "github.com/artefactual-labs/enduro/internal/amclient"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJobsService is a mock of JobsService interface
type MockJobsService struct {
	ctrl     *gomock.Controller
	recorder *MockJobsServiceMockRecorder
}

// MockJobsServiceMockRecorder is the mock recorder for MockJobsService
type MockJobsServiceMockRecorder struct {
	mock *MockJobsService
}

// NewMockJobsService creates a new mock instance
func NewMockJobsService(ctrl *gomock.Controller) *MockJobsService {
	mock := &MockJobsService{ctrl: ctrl}
	mock.recorder = &MockJobsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobsService) EXPECT() *MockJobsServiceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockJobsService) List(arg0 context.Context, arg1 string, arg2 *amclient.JobsListRequest) ([]amclient.Job, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]amclient.Job)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockJobsServiceMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockJobsService)(nil).List), arg0, arg1, arg2)
}