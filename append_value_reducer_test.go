// A simple implementation of futures in go
package gofutures

import (
	"testing"
)

func Test_AppendValueReducer_returns_empty_value_when_given_no_values(t *testing.T) {
	ar := NewAppendValueReducer(2)
	v := ar.Values()
	vals := v.([]interface{})
	if len(vals) != 0 {
		t.Fatalf("Expected [] values got %v\n", vals)
	}
}

func Test_AppendValueReducer_appends_values(t *testing.T) {
	val0 := 0
	val1 := 1
	val2 := 2

	ar := NewAppendValueReducer(2)
	ar.AddValue(val0)
	ar.AddValue(val1)
	ar.AddValue(val2)

	v := ar.Values()
	vals := v.([]interface{})

	if vals[0] != val0 || vals[1] != val1 || vals[2] != val2 {
		t.Fatalf("Expected appended values got %v\n", vals)
	}
}
