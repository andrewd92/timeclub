// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	club_dao "github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao"
	mock "github.com/stretchr/testify/mock"
)

// ClubDao is an autogenerated mock type for the ClubDao type
type ClubDao struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *ClubDao) GetAll() ([]*club_dao.ClubModel, error) {
	ret := _m.Called()

	var r0 []*club_dao.ClubModel
	if rf, ok := ret.Get(0).(func() []*club_dao.ClubModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*club_dao.ClubModel)
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

// GetById provides a mock function with given fields: id
func (_m *ClubDao) GetById(id int64) (*club_dao.ClubModel, error) {
	ret := _m.Called(id)

	var r0 *club_dao.ClubModel
	if rf, ok := ret.Get(0).(func(int64) *club_dao.ClubModel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*club_dao.ClubModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}