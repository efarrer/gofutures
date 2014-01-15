// A simple implementation of futures in go
package gofutures

import (
	"errors"
)

type concatenatedError struct {
	addChannel    chan error
	errorsChannel chan error
}

// An ErrorReducer that concatenates all error messages into a single error
func NewConcatenateErrorReducer(sep string) ErrorReducer {
	addChannel := make(chan error)
	errorsChannel := make(chan error)
	go func() {
		var allErr error = nil
		for {
			select {
			case err := <-addChannel:
				if err == nil {
					continue
				}
				if allErr == nil {
					allErr = err
				} else {
					allErr = errors.New(allErr.Error() + sep + err.Error())
				}
			case errorsChannel <- allErr:
				return
			}
		}
	}()
	return &concatenatedError{addChannel, errorsChannel}
}

func (ce *concatenatedError) AddError(err error) {
	ce.addChannel <- err
}

func (ce *concatenatedError) Errors() error {
	return <-ce.errorsChannel
}
