package mocks

import (
	"context"
	"function-srv/pkg/models"

	"github.com/stretchr/testify/mock"
)

type FunctionRepository struct {
	mock.Mock
}

func (_m *FunctionRepository) Create(_a0 context.Context, _a1 *models.Function) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Function) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *FunctionRepository) FindOne(_a0 context.Context, _a1 interface{}) (*models.Function, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Function
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *models.Function); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Function)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}