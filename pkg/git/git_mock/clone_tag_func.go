// Code generated by mockery 2.9.0. DO NOT EDIT.

package git_mock

import (
	git "github.com/gopaytech/go-commons/pkg/git"
	mock "github.com/stretchr/testify/mock"

	transport "github.com/go-git/go-git/v5/plumbing/transport"
)

// CloneTagFunc is an autogenerated mock type for the CloneTagFunc type
type CloneTagFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: url, destination, auth, tag
func (_m *CloneTagFunc) Execute(url string, destination string, auth transport.AuthMethod, tag string) (git.Repository, error) {
	ret := _m.Called(url, destination, auth, tag)

	var r0 git.Repository
	if rf, ok := ret.Get(0).(func(string, string, transport.AuthMethod, string) git.Repository); ok {
		r0 = rf(url, destination, auth, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(git.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, transport.AuthMethod, string) error); ok {
		r1 = rf(url, destination, auth, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
