package evaluators

type ErrorType uint32

const (
	NoCompilerFound = ErrorType(iota)
)

type ErrorHandler struct {
	errorMessage   string
	originialError error
	errorType      ErrorType
}

func NewErrorHandler(et ErrorType, oe error) *ErrorHandler {
	var error *ErrorHandler = new(ErrorHandler)

	error.errorMessage = func(e ErrorType) string {
		switch e {
		case NoCompilerFound:
			return "No available compilers were found"
		default:
			return "None"
		}
	}(et)

	error.errorType = et
	error.originialError = oe

	return error
}

func (error ErrorHandler) What() string {
	return error.errorMessage
}

func (error ErrorHandler) WhatType() ErrorType {
	return error.errorType
}

func (error ErrorHandler) Error() error {
	return error.originialError
}
