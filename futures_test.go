// A simple implementation of futures in go
package gofutures

import (
	"testing"
)

func test_call_Start_twice_returns_error(f Future, t *testing.T) {
	if err := f.Start(); err != nil {
		t.Fatalf("Start returned unexpected error %v\n", err)
	}
	if err := f.Start(); err == nil {
		t.Fatalf("Start should have returned and error.\n")
	}
	if val, err := f.Results(); val == nil || err != nil {
		t.Fatalf("Expected value and nil error but got %v, %v\n", val, err)
	}
}

func test_calling_Results_before_Start_returns_error(f Future, t *testing.T) {
	if val, err := f.Results(); val != nil || err == nil {
		t.Fatalf("Expected nil value and error but got %v, %v\n", val, err)
	}
}

func test_calling_Results_before_Start_doesnt_break_future(f Future, t *testing.T) {
	if val, err := f.Results(); val != nil || err == nil {
		t.Fatalf("Expected nil value and error but got %v, %v\n", val, err)
	}
	if err := f.Start(); err != nil {
		t.Fatalf("Start returned unexpected error %v\n", err)
	}
	if val, err := f.Results(); val == nil || err != nil {
		t.Fatalf("Expected value and nil error but got %v, %v\n", val, err)
	}
}

func test_calling_Results_twice_returns_error(f Future, t *testing.T) {
	if err := f.Start(); err != nil {
		t.Fatalf("Start returned unexpected error %v\n", err)
	}
	if val, err := f.Results(); val == nil || err != nil {
		t.Fatalf("Expected value and nil error but got %v, %v\n", val, err)
	}
	if val, err := f.Results(); val != nil || err == nil {
		t.Fatalf("Expected error on second call to Results %v, %v\n", val, err)
	}
}

func test_generic_future_behavior(f func() Future, t *testing.T) {
	test_call_Start_twice_returns_error(f(), t)
	test_calling_Results_before_Start_returns_error(f(), t)
	test_calling_Results_before_Start_doesnt_break_future(f(), t)
	test_calling_Results_twice_returns_error(f(), t)
}
