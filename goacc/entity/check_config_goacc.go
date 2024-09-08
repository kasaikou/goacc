// Code generated by github.com/kasaikou/goacc, DO NOT EDIT.
package entity

import "encoding/json"

// CheckConfigBuilder is an instance for generating an instance of CheckConfig.
type CheckConfigBuilder struct {
	__cc *CheckConfig
}

// NewCheckConfigBuilder creates an CheckConfigBuilder instance.
func NewCheckConfigBuilder(
	workingDir string,
	includePattern string,
) *CheckConfigBuilder {
	__cc := &CheckConfig{}

	__cc.workingDir = workingDir
	__cc.includePattern = includePattern

	return &CheckConfigBuilder{__cc: __cc}
}

// Purge purges CheckConfig instance from CheckConfigBuilder.
//
// If calls other method in CheckConfigBuilder after Purge called, it will be panic.
func (__ccb *CheckConfigBuilder) Purge() *CheckConfig {
	if __ccb == nil {
		panic("CheckConfigBuilder is nil")
	} else if __ccb.__cc != nil {
		__cc := __ccb.__cc
		__ccb.__cc = nil

		return __cc
	}

	panic("CheckConfig has been already purged")
}

func (__cc *CheckConfig) WorkingDir() string {
	if __cc != nil {
		return __cc.workingDir
	}

	panic("CheckConfig is nil")
}

func (__cc *CheckConfig) IncludePattern() string {
	if __cc != nil {
		return __cc.includePattern
	}

	panic("CheckConfig is nil")
}

func (__cc *CheckConfig) MarshalJSON() ([]byte, error) {

	type CheckConfigJSONContent struct {
		WorkingDir     string `json:"workingDir"`
		IncludePattern string `json:"includePattern"`
	}

	return json.Marshal(CheckConfigJSONContent{
		WorkingDir:     __cc.workingDir,
		IncludePattern: __cc.includePattern,
	})
}
