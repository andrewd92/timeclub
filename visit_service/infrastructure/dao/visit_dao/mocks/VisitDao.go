// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	visit "github.com/andrewd92/timeclub/visit_service/domain/visit"
	visit_dao "github.com/andrewd92/timeclub/visit_service/infrastructure/dao/visit_dao"
	mock "github.com/stretchr/testify/mock"
)

// VisitDao is an autogenerated mock type for the VisitDao type
type VisitDao struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *VisitDao) GetAll() ([]visit_dao.VisitModel, error) {
	ret := _m.Called()

	var r0 []visit_dao.VisitModel
	if rf, ok := ret.Get(0).(func() []visit_dao.VisitModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]visit_dao.VisitModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *VisitDao) Insert(_a0 *visit.Visit) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*visit.Visit) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*visit.Visit) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
