// Code generated by mockery v2.45.1. DO NOT EDIT.

package rest

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// GetApiV1UsersUserIdTodosResponseObject is an autogenerated mock type for the GetApiV1UsersUserIdTodosResponseObject type
type GetApiV1UsersUserIdTodosResponseObject struct {
	mock.Mock
}

type GetApiV1UsersUserIdTodosResponseObject_Expecter struct {
	mock *mock.Mock
}

func (_m *GetApiV1UsersUserIdTodosResponseObject) EXPECT() *GetApiV1UsersUserIdTodosResponseObject_Expecter {
	return &GetApiV1UsersUserIdTodosResponseObject_Expecter{mock: &_m.Mock}
}

// VisitGetApiV1UsersUserIdTodosResponse provides a mock function with given fields: w
func (_m *GetApiV1UsersUserIdTodosResponseObject) VisitGetApiV1UsersUserIdTodosResponse(w http.ResponseWriter) error {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for VisitGetApiV1UsersUserIdTodosResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisitGetApiV1UsersUserIdTodosResponse'
type GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call struct {
	*mock.Call
}

// VisitGetApiV1UsersUserIdTodosResponse is a helper method to define mock.On call
//   - w http.ResponseWriter
func (_e *GetApiV1UsersUserIdTodosResponseObject_Expecter) VisitGetApiV1UsersUserIdTodosResponse(w interface{}) *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call {
	return &GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call{Call: _e.mock.On("VisitGetApiV1UsersUserIdTodosResponse", w)}
}

func (_c *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call) Run(run func(w http.ResponseWriter)) *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter))
	})
	return _c
}

func (_c *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call) Return(_a0 error) *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call) RunAndReturn(run func(http.ResponseWriter) error) *GetApiV1UsersUserIdTodosResponseObject_VisitGetApiV1UsersUserIdTodosResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetApiV1UsersUserIdTodosResponseObject creates a new instance of GetApiV1UsersUserIdTodosResponseObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetApiV1UsersUserIdTodosResponseObject(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetApiV1UsersUserIdTodosResponseObject {
	mock := &GetApiV1UsersUserIdTodosResponseObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
