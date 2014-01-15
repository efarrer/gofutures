// A simple implementation of futures in go
package gofutures

// An ErrorReducer is a interface for combining errors from multiple Futures
type ErrorReducer interface {
	// Adds an error to the ErrorReducer
	AddError(error)

	// Retrieves the combined error
	Errors() error
}
