// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/alfathaulia/ca_ecommerce_api/domain"
	mock "github.com/stretchr/testify/mock"
)

// ReviewRepository is an autogenerated mock type for the ReviewRepository type
type ReviewRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ReviewRepository) Fetch(ctx context.Context, cursor string, num int) ([]domain.Review, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.Review
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []domain.Review); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Review)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ReviewRepository) GetByID(ctx context.Context, id int64) (domain.Review, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Review
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Review); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Review)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}