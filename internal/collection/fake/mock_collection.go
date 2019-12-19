// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-labs/enduro/internal/collection (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	collection "github.com/artefactual-labs/enduro/internal/api/gen/collection"
	collection0 "github.com/artefactual-labs/enduro/internal/collection"
	gomock "github.com/golang/mock/gomock"
	http "goa.design/goa/v3/http"
	http0 "net/http"
	reflect "reflect"
	time "time"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockService) Create(arg0 context.Context, arg1 *collection0.Collection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0, arg1)
}

// Goa mocks base method
func (m *MockService) Goa() collection.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Goa")
	ret0, _ := ret[0].(collection.Service)
	return ret0
}

// Goa indicates an expected call of Goa
func (mr *MockServiceMockRecorder) Goa() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Goa", reflect.TypeOf((*MockService)(nil).Goa))
}

// HTTPDownload mocks base method
func (m *MockService) HTTPDownload(arg0 http.Muxer, arg1 func(*http0.Request) http.Decoder) http0.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPDownload", arg0, arg1)
	ret0, _ := ret[0].(http0.HandlerFunc)
	return ret0
}

// HTTPDownload indicates an expected call of HTTPDownload
func (mr *MockServiceMockRecorder) HTTPDownload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPDownload", reflect.TypeOf((*MockService)(nil).HTTPDownload), arg0, arg1)
}

// UpdateWorkflowStatus mocks base method
func (m *MockService) UpdateWorkflowStatus(arg0 context.Context, arg1 uint, arg2, arg3, arg4, arg5, arg6, arg7 string, arg8 collection0.Status, arg9 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkflowStatus", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWorkflowStatus indicates an expected call of UpdateWorkflowStatus
func (mr *MockServiceMockRecorder) UpdateWorkflowStatus(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowStatus", reflect.TypeOf((*MockService)(nil).UpdateWorkflowStatus), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}