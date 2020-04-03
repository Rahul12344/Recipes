package errors

import (
	"fmt"
)

/* CUSTOMIZED ERROR */

//ErrorTypes types of error
type ErrorTypes int

const (
	// None error
	None ErrorTypes = iota
	//Service corresponds to service being down i.e. Redis or Postgres
	Service
	//Access Invalid access to cache/storage
	Access
	//Association Association creation error
	Association
)

//Errors struct of error
type Errors struct {
	ErrorType         ErrorTypes
	PreformattedError error
}

//Error sets Errors to the Error interface
func (e Errors) Error() string {
	return e.PreformattedError.Error()
}

//Cause sets Errors to the Cause interface
func (e Errors) Cause() error {
	return e.PreformattedError
}

//New creates new error
func New(errorMessage string, args ...interface{}) error {
	return Errors{
		ErrorType:         None,
		PreformattedError: fmt.Errorf(errorMessage, args...),
	}
}
