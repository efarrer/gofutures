// A simple implementation of futures in go
package gofutures

import (
	"testing"
)

func Test_ConcatenateStringValueReducer_returns_empty_string_when_given_no_values(t *testing.T) {
	ar := NewConcatenateStringValueReducer("\n")
	v := ar.Values()
	vals := v.(string)
	if vals != "" {
		t.Fatalf("Expected \"\" values got %v\n", vals)
	}
}

func Test_ConcatenateStringValueReducer_appends_values_as_strings(t *testing.T) {
	val0 := 0
	val1 := struct{ x int }{1}
	val2 := "2"

	ar := NewConcatenateStringValueReducer(";")
	ar.AddValue(val0)
	ar.AddValue(val1)
	ar.AddValue(val2)

	v := ar.Values()
	val := v.(string)

	expectedStr := "0;{1};2;"
	if val != expectedStr {
		t.Fatalf("Expected appended string value \"%v\". Got \"%v\"\n", expectedStr, val)
	}
}

func Test_ConcatenateStringValueReducer_ignores_empty_strings(t *testing.T) {
	val0 := ""
	val1 := struct{ x int }{1}
	val2 := ""

	ar := NewConcatenateStringValueReducer(";")
	ar.AddValue(val0)
	ar.AddValue(val1)
	ar.AddValue(val2)

	v := ar.Values()
	val := v.(string)

	expectedStr := "{1};"
	if val != expectedStr {
		t.Fatalf("Expected appended string value \"%v\". Got \"%v\"\n", expectedStr, val)
	}
}
