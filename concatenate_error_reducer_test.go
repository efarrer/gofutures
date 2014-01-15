// A simple implementation of futures in go
package gofutures

import (
	"errors"
	"testing"
)

func Test_ConcatenateErrorReducer_generic_behavior(t *testing.T) {
	erf := func() ErrorReducer {
		return NewConcatenateErrorReducer("\n")
	}
	test_generic_ErrorReducer_behavior(erf, t)
}

func Test_ConcatenateErrorReducer_concatenate_errors(t *testing.T) {
	sep := "\n"
	errSt0 := "Hello"
	errSt1 := "World"
	errSt2 := "!"
	err0 := errors.New(errSt0)
	err1 := errors.New(errSt1)
	err2 := errors.New(errSt2)

	cer := NewConcatenateErrorReducer(sep)
	cer.AddError(err0)
	cer.AddError(err1)
	cer.AddError(err2)

	expectedStr := errSt0 + sep + errSt1 + sep + errSt2
	actualStr := cer.Errors().Error()
	if actualStr != expectedStr {
		t.Fatalf("Expected error with string\n\"%v\", got\n\"%v\".\n",
			actualStr, expectedStr)
	}
}

func Test_ConcatenateErrorReducer_handles_nil_errors(t *testing.T) {
	cer := NewConcatenateErrorReducer("\n")
	cer.AddError(nil)
	cer.AddError(nil)
	if cer.Errors() != nil {
		t.Fatalf("Expected nil error from concatenating nil errors\n")
	}
}
