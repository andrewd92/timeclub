// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	api "github.com/andrewd92/timeclub/club_service/api"
	mock "github.com/stretchr/testify/mock"

	time "time"

	visit "github.com/andrewd92/timeclub/visit_service/domain/visit"
)

// Marshaller is an autogenerated mock type for the Marshaller type
type Marshaller struct {
	mock.Mock
}

// Marshal provides a mock function with given fields: _a0, now, club
func (_m *Marshaller) Marshal(_a0 *visit.Visit, now time.Time, club *api.Club) (interface{}, error) {
	ret := _m.Called(_a0, now, club)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*visit.Visit, time.Time, *api.Club) interface{}); ok {
		r0 = rf(_a0, now, club)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*visit.Visit, time.Time, *api.Club) error); ok {
		r1 = rf(_a0, now, club)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarshalAll provides a mock function with given fields: visits, club, now
func (_m *Marshaller) MarshalAll(visits []*visit.Visit, club *api.Club, now time.Time) ([]interface{}, error) {
	ret := _m.Called(visits, club, now)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func([]*visit.Visit, *api.Club, time.Time) []interface{}); ok {
		r0 = rf(visits, club, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*visit.Visit, *api.Club, time.Time) error); ok {
		r1 = rf(visits, club, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
