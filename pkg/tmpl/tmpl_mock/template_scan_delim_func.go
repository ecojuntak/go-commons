// Code generated by mockery 2.9.0. DO NOT EDIT.

package tmpl_mock

import (
	tmpl "github.com/gopaytech/go-commons/pkg/tmpl"
	mock "github.com/stretchr/testify/mock"
)

// TemplateScanDelimFunc is an autogenerated mock type for the TemplateScanDelimFunc type
type TemplateScanDelimFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: scanPath, filter, tmplExt, startDelims, endDelims
func (_m *TemplateScanDelimFunc) Execute(scanPath string, filter tmpl.FileFilter, tmplExt string, startDelims string, endDelims string) (tmpl.ScanResult, error) {
	ret := _m.Called(scanPath, filter, tmplExt, startDelims, endDelims)

	var r0 tmpl.ScanResult
	if rf, ok := ret.Get(0).(func(string, tmpl.FileFilter, string, string, string) tmpl.ScanResult); ok {
		r0 = rf(scanPath, filter, tmplExt, startDelims, endDelims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tmpl.ScanResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, tmpl.FileFilter, string, string, string) error); ok {
		r1 = rf(scanPath, filter, tmplExt, startDelims, endDelims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}