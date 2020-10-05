// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	models "github.com/smartcontractkit/chainlink/core/store/models"
	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// BridgeResponseURL provides a mock function with given fields:
func (_m *Config) BridgeResponseURL() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}

// DefaultHTTPLimit provides a mock function with given fields:
func (_m *Config) DefaultHTTPLimit() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// DefaultHTTPTimeout provides a mock function with given fields:
func (_m *Config) DefaultHTTPTimeout() models.Duration {
	ret := _m.Called()

	var r0 models.Duration
	if rf, ok := ret.Get(0).(func() models.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Duration)
	}

	return r0
}

// DefaultMaxHTTPAttempts provides a mock function with given fields:
func (_m *Config) DefaultMaxHTTPAttempts() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// PipelineRunnerParallelism provides a mock function with given fields:
func (_m *Config) PipelineRunnerParallelism() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}