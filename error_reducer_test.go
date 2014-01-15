// A simple implementation of futures in go
package gofutures

import (
	"testing"
)

func test_ErrorReducer_returns_nil_when_given_no_errors(er ErrorReducer, t *testing.T) {
	if err := er.Errors(); err != nil {
		t.Fatalf("Expected nil error but got %v.\n", err)
	}
}

func test_generic_ErrorReducer_behavior(erf func() ErrorReducer, t *testing.T) {
	test_ErrorReducer_returns_nil_when_given_no_errors(erf(), t)
}
