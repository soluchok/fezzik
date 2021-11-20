// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	basic "github.com/inigolabs/fezzik/examples/basic/gen/basic"

	fezzik "github.com/inigolabs/fezzik"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// OneAdd provides a mock function with given fields: ctx, input
func (_m *Client) OneAdd(ctx context.Context, input *basic.OneAddInputArgs) (*basic.OneAddResponse, *fezzik.GQLErrors, error) {
	ret := _m.Called(ctx, input)

	var r0 *basic.OneAddResponse
	if rf, ok := ret.Get(0).(func(context.Context, *basic.OneAddInputArgs) *basic.OneAddResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneAddResponse)
		}
	}

	var r1 *fezzik.GQLErrors
	if rf, ok := ret.Get(1).(func(context.Context, *basic.OneAddInputArgs) *fezzik.GQLErrors); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*fezzik.GQLErrors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *basic.OneAddInputArgs) error); ok {
		r2 = rf(ctx, input)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OneAllTypes provides a mock function with given fields: ctx
func (_m *Client) OneAllTypes(ctx context.Context) (*basic.OneAllTypesResponse, *fezzik.GQLErrors, error) {
	ret := _m.Called(ctx)

	var r0 *basic.OneAllTypesResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.OneAllTypesResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneAllTypesResponse)
		}
	}

	var r1 *fezzik.GQLErrors
	if rf, ok := ret.Get(1).(func(context.Context) *fezzik.GQLErrors); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*fezzik.GQLErrors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OneWithSubSelections provides a mock function with given fields: ctx
func (_m *Client) OneWithSubSelections(ctx context.Context) (*basic.OneWithSubSelectionsResponse, *fezzik.GQLErrors, error) {
	ret := _m.Called(ctx)

	var r0 *basic.OneWithSubSelectionsResponse
	if rf, ok := ret.Get(0).(func(context.Context) *basic.OneWithSubSelectionsResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.OneWithSubSelectionsResponse)
		}
	}

	var r1 *fezzik.GQLErrors
	if rf, ok := ret.Get(1).(func(context.Context) *fezzik.GQLErrors); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*fezzik.GQLErrors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// QueryWithInputs provides a mock function with given fields: ctx, input
func (_m *Client) QueryWithInputs(ctx context.Context, input *basic.QueryWithInputsInputArgs) (*basic.QueryWithInputsResponse, *fezzik.GQLErrors, error) {
	ret := _m.Called(ctx, input)

	var r0 *basic.QueryWithInputsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *basic.QueryWithInputsInputArgs) *basic.QueryWithInputsResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basic.QueryWithInputsResponse)
		}
	}

	var r1 *fezzik.GQLErrors
	if rf, ok := ret.Get(1).(func(context.Context, *basic.QueryWithInputsInputArgs) *fezzik.GQLErrors); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*fezzik.GQLErrors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *basic.QueryWithInputsInputArgs) error); ok {
		r2 = rf(ctx, input)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
