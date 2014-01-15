// A simple implementation of futures in go
package gofutures

import (
	"errors"
	"testing"
)

func createSuccessFuture(val int) Future {
	return NewFuncFuture(func() (interface{}, error) {
		return val, nil
	})
}

func createFailureFuture() Future {
	return NewFuncFuture(func() (interface{}, error) {
		return nil, errors.New("fail")
	})
}

func Test_generic_FuncFuture_behavior(t *testing.T) {
	ff := func() Future {
		return createSuccessFuture(1)
	}
	test_generic_future_behavior(ff, t)
}

func Test_FuncFuture_Results_returns_value_on_success(t *testing.T) {
	value := 1
	f := createSuccessFuture(value)
	if err := f.Start(); err != nil {
		t.Fatalf("Start returned unexpected error %v\n", err)
	}
	if val, err := f.Results(); value != val.(int) || err != nil {
		t.Fatalf("Expected value and nil error but got %v, %v\n", val, err)
	}
}

func Test_FuncFuture_Results_returns_error_on_failure(t *testing.T) {
	f := createFailureFuture()
	if err := f.Start(); err != nil {
		t.Fatalf("Start returned unexpected error %v\n", err)
	}
	if val, err := f.Results(); val != nil || err == nil {
		t.Fatalf("Expected nil value and error but got %v, %v\n", val, err)
	}
}
