// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	users "investaBackend/business/users"

	mock "github.com/stretchr/testify/mock"
)

// UserRepoInterface is an autogenerated mock type for the UserRepoInterface type
type UserRepoInterface struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepoInterface) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: domain, ctx
func (_m *UserRepoInterface) Register(domain *users.Domain, ctx context.Context) error {
	ret := _m.Called(domain, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*users.Domain, context.Context) error); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}