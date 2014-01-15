// A simple implementation of futures in go
package gofutures

import (
	"testing"
	"time"
)

func test_compound_future_Results_error_is_nil_if_has_no_futures(ff func(...Future) Future, t *testing.T) {
	f := ff()
	if err := f.Start(); err != nil {
		t.Fatalf("Start returned unexpected error %v\n", err)
	}
	if _, err := f.Results(); err != nil {
		t.Fatalf("Expected nil error but got %v\n", err)
	}
}

func test_compound_future_passes_Start_errors_onto_Results(ff func(...Future) Future, t *testing.T) {
	before := createSuccessFuture(1)
	after := createSuccessFuture(1)
	serial := ff(before, after)

	before.Start()

	if err := serial.Start(); err != nil {
		t.Fatalf("Didn't expect error on Future.Start but got %v.\n", err)
	}
	if _, err := serial.Results(); err == nil {
		t.Fatalf("Compound futures Start error should be reflected in the results\n")
	}
}

func test_compound_future_errors_from_futures_passed_onto_Results(ff func(...Future) Future, t *testing.T) {
	serial := ff(createSuccessFuture(1), createFailureFuture())

	if err := serial.Start(); err != nil {
		t.Fatalf("Didn't expect error on Future.Start but got %v.\n", err)
	}
	if _, err := serial.Results(); err == nil {
		t.Fatalf("Compound futures Results error should be reflected in the results\n")
	}
}

func test_compound_future_values_are_included_in_Results(ff func(...Future) Future, t *testing.T) {
	after := createSuccessFuture(1)
	before := createSuccessFuture(2)
	serial := ff(after, before)

	serial.Start()

	val, _ := serial.Results()
	vals := val.([]interface{})
	if vals[0].(int) != 1 || vals[1].(int) != 2 {
		t.Fatalf("Compound futures values should be reflected in the results %v\n", vals)
	}
}

func test_generic_compound_future_behavior(ff func(...Future) Future, t *testing.T) {
	test_compound_future_Results_error_is_nil_if_has_no_futures(ff, t)
	test_compound_future_passes_Start_errors_onto_Results(ff, t)
	test_compound_future_errors_from_futures_passed_onto_Results(ff, t)
	test_compound_future_values_are_included_in_Results(ff, t)

	f := func() Future {
		return ff(createSuccessFuture(1))
	}
	test_generic_future_behavior(f, t)
}

/*
 * Test SerializeFutures
 */
func Test_generic_SerialFuture_behavior(t *testing.T) {
	ff := func(futures ...Future) Future {
		vr := NewAppendValueReducer(0)
		er := NewConcatenateErrorReducer("\n")
		return SerializeFutures(vr, er, futures...)
	}

	test_generic_compound_future_behavior(ff, t)
}

func Test_SerializeFutures_serializes_futures(t *testing.T) {
	FIRST := 1
	SECOND := 2
	THIRD := 3
	ch := make(chan int, 3)
	first := NewFuncFuture(func() (interface{}, error) {
		<-time.After(30 * time.Millisecond)
		ch <- FIRST
		return nil, nil
	})
	second := NewFuncFuture(func() (interface{}, error) {
		<-time.After(20 * time.Millisecond)
		ch <- SECOND
		return nil, nil
	})
	third := NewFuncFuture(func() (interface{}, error) {
		ch <- THIRD
		return nil, nil
	})

	vr := NewAppendValueReducer(0)
	er := NewConcatenateErrorReducer("\n")
	serial := SerializeFutures(vr, er, first, second, third)
	serial.Start()
	serial.Results()

	if val := <-ch; FIRST != val {
		t.Fatalf("Expected first value but got %v\n", val)
	}
	if val := <-ch; SECOND != val {
		t.Fatalf("Expected second value but got %v\n", val)
	}
	if val := <-ch; THIRD != val {
		t.Fatalf("Expected third value but got %v\n", val)
	}
}

/*
 * Test ParallelFutures
 */
func Test_generic_ParallelFuture_behavior(t *testing.T) {
	ff := func(futures ...Future) Future {
		vr := NewAppendValueReducer(0)
		er := NewConcatenateErrorReducer("\n")
		return ParallelFutures(vr, er, futures...)
	}
	test_generic_compound_future_behavior(ff, t)
}

func Test_ParallelFutures_parallelizes_futures(t *testing.T) {
	FIRST := 1
	SECOND := 2
	THIRD := 3
	ch := make(chan int, 3)
	first := NewFuncFuture(func() (interface{}, error) {
		<-time.After(30 * time.Millisecond)
		ch <- FIRST
		return nil, nil
	})
	second := NewFuncFuture(func() (interface{}, error) {
		<-time.After(20 * time.Millisecond)
		ch <- SECOND
		return nil, nil
	})
	third := NewFuncFuture(func() (interface{}, error) {
		ch <- THIRD
		return nil, nil
	})

	vr := NewAppendValueReducer(0)
	er := NewConcatenateErrorReducer("\n")
	serial := ParallelFutures(vr, er, first, second, third)
	serial.Start()
	serial.Results()

	if val := <-ch; THIRD != val {
		t.Fatalf("Expected third value but got %v\n", val)
	}
	if val := <-ch; SECOND != val {
		t.Fatalf("Expected second value but got %v\n", val)
	}
	if val := <-ch; FIRST != val {
		t.Fatalf("Expected first value but got %v\n", val)
	}
}
