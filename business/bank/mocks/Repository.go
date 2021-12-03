// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	bank "investaBackend/business/bank"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAllDataBank provides a mock function with given fields: ctx
func (_m *Repository) GetAllDataBank(ctx context.Context) ([]bank.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []bank.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []bank.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bank.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBank provides a mock function with given fields: domain, ctx
func (_m *Repository) InsertBank(domain bank.Domain, ctx context.Context) (bank.Domain, error) {
	ret := _m.Called(domain, ctx)

	var r0 bank.Domain
	if rf, ok := ret.Get(0).(func(bank.Domain, context.Context) bank.Domain); ok {
		r0 = rf(domain, ctx)
	} else {
		r0 = ret.Get(0).(bank.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bank.Domain, context.Context) error); ok {
		r1 = rf(domain, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
