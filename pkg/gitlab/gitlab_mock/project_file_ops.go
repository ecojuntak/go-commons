// Code generated by mockery v2.9.4. DO NOT EDIT.

package gitlab_mock

import (
	gitlab "github.com/gopaytech/go-commons/pkg/gitlab"
	gitlab_file "github.com/gopaytech/go-commons/pkg/gitlab/gitlab_file"

	go_gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// ProjectFileOps is an autogenerated mock type for the ProjectFileOps type
type ProjectFileOps struct {
	mock.Mock
}

// SearchFileWithContent provides a mock function with given fields: id, ref, searchFunc, fileNames
func (_m *ProjectFileOps) SearchFileWithContent(id gitlab.NameOrId, ref string, searchFunc gitlab_file.SearchFunc, fileNames ...string) ([]go_gitlab.File, error) {
	_va := make([]interface{}, len(fileNames))
	for _i := range fileNames {
		_va[_i] = fileNames[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, ref, searchFunc)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []go_gitlab.File
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, string, gitlab_file.SearchFunc, ...string) []go_gitlab.File); ok {
		r0 = rf(id, ref, searchFunc, fileNames...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_gitlab.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId, string, gitlab_file.SearchFunc, ...string) error); ok {
		r1 = rf(id, ref, searchFunc, fileNames...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchFiles provides a mock function with given fields: id, ref, fileNames
func (_m *ProjectFileOps) SearchFiles(id gitlab.NameOrId, ref string, fileNames ...string) ([]go_gitlab.File, error) {
	_va := make([]interface{}, len(fileNames))
	for _i := range fileNames {
		_va[_i] = fileNames[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, ref)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []go_gitlab.File
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, string, ...string) []go_gitlab.File); ok {
		r0 = rf(id, ref, fileNames...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_gitlab.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId, string, ...string) error); ok {
		r1 = rf(id, ref, fileNames...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchFilesMetadata provides a mock function with given fields: id, ref, fileNames
func (_m *ProjectFileOps) SearchFilesMetadata(id gitlab.NameOrId, ref string, fileNames ...string) ([]go_gitlab.File, error) {
	_va := make([]interface{}, len(fileNames))
	for _i := range fileNames {
		_va[_i] = fileNames[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, ref)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []go_gitlab.File
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, string, ...string) []go_gitlab.File); ok {
		r0 = rf(id, ref, fileNames...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]go_gitlab.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId, string, ...string) error); ok {
		r1 = rf(id, ref, fileNames...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
