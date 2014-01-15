// A simple implementation of futures in go
package gofutures

// A ValueReducer is a interface for combining values from multiple futures
type ValueReducer interface {
	// Adds a value to the ValueReducer
	AddValue(interface{})

	// Retrieves the combined values
	Values() interface{}
}
