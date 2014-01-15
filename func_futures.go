// A simple implementation of futures in go
package gofutures

import (
	"errors"
)

// Simple structure for functions wrapped in a Future
type funcFuture struct {
	startCmdCh <-chan bool     // Channel for starting the calculation
	resultsCh  <-chan *results // Blocks until calculation is finished
}

// A calculated results
type results struct {
	value interface{}
	err   error
}

// Creates a new future from a function
func NewFuncFuture(calculationFunc func() (interface{}, error)) Future {
	startCmdCh := make(chan bool)
	resultsCh := make(chan *results)
	res := &results{nil, errors.New("Error: Future.Results called before Future.Start.")}
	go func() {
		started := false
		for {
			select {
			case startCmdCh <- true:
				close(startCmdCh)
				startCmdCh = nil
				value, err := calculationFunc()
				res = &results{value, err}
				started = true
			case resultsCh <- res:
				if started {
					close(resultsCh)
					return
				}
			}
		}
	}()
	return &funcFuture{startCmdCh, resultsCh}
}

func (df *funcFuture) Start() error {
	if !<-df.startCmdCh {
		return errors.New("Future.Start should not be called more than once.")
	}

	return nil
}

func (df *funcFuture) Results() (interface{}, error) {
	results := <-df.resultsCh
	if results == nil {
		return nil, errors.New("Future.Results should not be called more than once.")
	}
	return results.value, results.err
}
