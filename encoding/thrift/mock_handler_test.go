// Automatically generated by MockGen. DO NOT EDIT!
// Source: go.uber.org/yarpc/encoding/thrift (interfaces: Handler)

package thrift

import (
	gomock "github.com/golang/mock/gomock"
	wire "go.uber.org/thriftrw/wire"
	yarpc_go "go.uber.org/yarpc"
	context "golang.org/x/net/context"
)

// Mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockHandlerRecorder
}

// Recorder for MockHandler (not exported)
type _MockHandlerRecorder struct {
	mock *MockHandler
}

func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &_MockHandlerRecorder{mock}
	return mock
}

func (_m *MockHandler) EXPECT() *_MockHandlerRecorder {
	return _m.recorder
}

func (_m *MockHandler) Handle(_param0 context.Context, _param1 yarpc_go.ReqMeta, _param2 wire.Value) (Response, error) {
	ret := _m.ctrl.Call(_m, "Handle", _param0, _param1, _param2)
	ret0, _ := ret[0].(Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHandlerRecorder) Handle(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Handle", arg0, arg1, arg2)
}
