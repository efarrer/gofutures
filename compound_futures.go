// A simple implementation of futures in go
package gofutures

// A FutureGroup is a collection of futures that can be processed together
type FutureGroup interface {
	AddFutures(...Future) FutureGroup

	Finalize() Future
}

type futureGroupData struct {
	futures []Future
	fn      func() (interface{}, error)
}

func (fg *futureGroupData) AddFutures(futures ...Future) FutureGroup {
	fg.futures = append(fg.futures, futures...)
	return fg
}

func (fg *futureGroupData) Finalize() Future {
	return NewFuncFuture(fg.fn)
}

/*
 * Creates a FutureGroup that will executing the group of futures serially.
 * Results and errors are combined using the ValueReducer and ErrorReducer
 * respectively.
 */
func SerializeFutureGroup(valueReducer ValueReducer, errorReducer ErrorReducer) FutureGroup {
	fg := &futureGroupData{[]Future{}, nil}
	fg.fn = func() (interface{}, error) {
		for i := 0; i != len(fg.futures); i++ {
			err := fg.futures[i].Start()
			errorReducer.AddError(err)
			val, err := fg.futures[i].Results()
			valueReducer.AddValue(val)
			errorReducer.AddError(err)
		}
		return valueReducer.Values(), errorReducer.Errors()
	}
	return fg
}

/*
 * Creates a FutureGroup that will executing the group of futures in parallel.
 * Results and errors are combined using the ValueReducer and ErrorReducer
 * respectively.
 */
func ParallelFutureGroup(valueReducer ValueReducer, errorReducer ErrorReducer) FutureGroup {
	fg := &futureGroupData{[]Future{}, nil}
	fg.fn = func() (interface{}, error) {
		for i := 0; i != len(fg.futures); i++ {
			err := fg.futures[i].Start()
			errorReducer.AddError(err)
		}
		for i := 0; i != len(fg.futures); i++ {
			val, err := fg.futures[i].Results()
			valueReducer.AddValue(val)
			errorReducer.AddError(err)
		}
		return valueReducer.Values(), errorReducer.Errors()
	}
	return fg
}
