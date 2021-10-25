// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/alfathaulia/ca_ecommerce_api/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ProductRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ProductRepository) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Product, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.Product); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ProductRepository) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *ProductRepository) Store(ctx context.Context, a *domain.Product) error {
	ret := _m.Called(ctx, a)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Product) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *ProductRepository) Update(ctx context.Context, ar *domain.Product) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Product) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}