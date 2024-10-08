// Code generated by github.com/kasaikou/goacc, DO NOT EDIT.
// defaultTag=-
package entity

import "encoding/json"

type CheckConfigBuilder interface {
	Build() *CheckConfig
}

// checkConfigBuilderImpl is an instance for generating an instance of CheckConfig.
type checkConfigBuilderImpl struct {
	__cc *CheckConfig
}

// NewCheckConfigBuilder creates an CheckConfigBuilder instance.
func NewCheckConfigBuilder(
	workingDir string,
	includePattern string,
) CheckConfigBuilder {
	__cc := &CheckConfig{}

	__cc.workingDir = workingDir
	__cc.includePattern = includePattern

	return &checkConfigBuilderImpl{__cc: __cc}
}

// Build purges CheckConfig instance from checkConfigBuilderImpl.
//
// If calls other method in checkConfigBuilderImpl after Purge called, it will be panic.
func (__ccb *checkConfigBuilderImpl) Build() *CheckConfig {
	if __ccb == nil {
		panic("checkConfigBuilderImpl is nil")
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
