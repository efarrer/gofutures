// A simple implementation of futures in go
package gofutures

type Future interface {

	/*
	 * Starts the calculation.
	 * Subsequent calls return an error.
	 */
	Start() error

	/*
	 * The calculated result.
	 * Subsequent calls return an error.
	 */
	Results() (interface{}, error)
}
