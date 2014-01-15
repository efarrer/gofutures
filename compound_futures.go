// A simple implementation of futures in go
package gofutures

/*
 * Creates a future containing the results of executing the passed futures
 * serially. Results and errors are combined using the ValueReducer and
 * ErrorReducer respectively.
 */
func SerializeFutures(valueReducer ValueReducer, errorReducer ErrorReducer, futures ...Future) Future {
	fn := func() (interface{}, error) {
		for i := 0; i != len(futures); i++ {
			err := futures[i].Start()
			errorReducer.AddError(err)
			val, err := futures[i].Results()
			valueReducer.AddValue(val)
			errorReducer.AddError(err)
		}
		return valueReducer.Values(), errorReducer.Errors()
	}
	return NewFuncFuture(fn)
}

/*
 * Creates a future containing the results of executing the passed futures
 * serially. Results and errors are combined using the ValueReducer and
 * ErrorReducer respectively.
 */
func ParallelFutures(valueReducer ValueReducer, errorReducer ErrorReducer, futures ...Future) Future {
	fn := func() (interface{}, error) {
		for i := 0; i != len(futures); i++ {
			err := futures[i].Start()
			errorReducer.AddError(err)
		}
		for i := 0; i != len(futures); i++ {
			val, err := futures[i].Results()
			valueReducer.AddValue(val)
			errorReducer.AddError(err)
		}
		return valueReducer.Values(), errorReducer.Errors()
	}
	return NewFuncFuture(fn)
}
