package evaluators

import (
	"os"
)

type ErrorType uint32

const (
	NoCompilerFound = ErrorType(iota)
	CompileError
	RunTimeError
)

type ErrorHandler struct {
	errorMessage   string
	originialError *error
	errorType      ErrorType
	errorOutput    string
}

func NewErrorHandler(errorType ErrorType, originalError *error, errorOutput string) *ErrorHandler {
	var error *ErrorHandler = new(ErrorHandler)

	error.errorMessage = func(e ErrorType) string {
		switch e {
		case NoCompilerFound:
			return "No available compilers were found"
		case CompileError:
			return "Error while compiling sources"
		case RunTimeError:
			return "Error at runtime"
		default:
			return "None"
		}
	}(errorType)

	error.errorType = errorType
	error.originialError = originalError
	error.errorOutput = errorOutput

	return error
}

func (error ErrorHandler) What() string {
	return error.errorMessage
}

func (error ErrorHandler) WhatType() ErrorType {
	return error.errorType
}

func (error ErrorHandler) Error() string {
	return error.errorOutput
}

func (error *ErrorHandler) WriteToStderr() {
	os.Stderr.WriteString(error.What() + "\n" + error.Error() + "\n")
}
