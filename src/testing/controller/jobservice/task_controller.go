// Code generated by mockery v2.1.0. DO NOT EDIT.

package jobservice

import (
	context "context"

	jobservice "github.com/goharbor/harbor/src/controller/jobservice"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// TaskController is an autogenerated mock type for the TaskController type
type TaskController struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, vendorType, id
func (_m *TaskController) Get(ctx context.Context, vendorType string, id int64) (*jobservice.Task, error) {
	ret := _m.Called(ctx, vendorType, id)

	var r0 *jobservice.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *jobservice.Task); ok {
		r0 = rf(ctx, vendorType, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jobservice.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, vendorType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLog provides a mock function with given fields: ctx, vendorType, id
func (_m *TaskController) GetLog(ctx context.Context, vendorType string, id int64) ([]byte, error) {
	ret := _m.Called(ctx, vendorType, id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []byte); ok {
		r0 = rf(ctx, vendorType, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, vendorType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, vendorType, query
func (_m *TaskController) List(ctx context.Context, vendorType string, query *q.Query) ([]*jobservice.Task, error) {
	ret := _m.Called(ctx, vendorType, query)

	var r0 []*jobservice.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, *q.Query) []*jobservice.Task); ok {
		r0 = rf(ctx, vendorType, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*jobservice.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *q.Query) error); ok {
		r1 = rf(ctx, vendorType, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}