// Code generated by mockery v1.0.0. DO NOT EDIT.

package v1

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockSomewhatServer is an autogenerated mock type for the SomewhatServer type
type MockSomewhatServer struct {
	mock.Mock
}

// CreateSomething provides a mock function with given fields: _a0, _a1
func (_m *MockSomewhatServer) CreateSomething(_a0 context.Context, _a1 *CreateSomethingRequest) (*CreateSomethingResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *CreateSomethingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *CreateSomethingRequest) *CreateSomethingResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateSomethingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CreateSomethingRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSomething provides a mock function with given fields: _a0, _a1
func (_m *MockSomewhatServer) DeleteSomething(_a0 context.Context, _a1 *DeleteSomethingRequest) (*DeleteSomethingResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *DeleteSomethingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteSomethingRequest) *DeleteSomethingResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteSomethingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteSomethingRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSomething provides a mock function with given fields: _a0, _a1
func (_m *MockSomewhatServer) GetSomething(_a0 context.Context, _a1 *GetSomethingRequest) (*GetSomethingResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetSomethingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetSomethingRequest) *GetSomethingResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetSomethingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetSomethingRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSomething provides a mock function with given fields: _a0, _a1
func (_m *MockSomewhatServer) ListSomething(_a0 context.Context, _a1 *ListSomethingRequest) (*ListSomethingResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ListSomethingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ListSomethingRequest) *ListSomethingResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListSomethingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListSomethingRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSomething provides a mock function with given fields: _a0, _a1
func (_m *MockSomewhatServer) UpdateSomething(_a0 context.Context, _a1 *UpdateSomethingRequest) (*UpdateSomethingResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UpdateSomethingResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateSomethingRequest) *UpdateSomethingResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateSomethingResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateSomethingRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
