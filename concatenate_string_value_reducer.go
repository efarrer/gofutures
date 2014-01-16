// A simple implementation of futures in go
package gofutures

import (
	"fmt"
)

type stringValueReducer struct {
	addChannel     chan interface{}
	resultsChannel chan interface{}
}

// Creates a ValueReducer that converts values to strings and concatenates them.
func NewConcatenateStringValueReducer(sep string) ValueReducer {
	addChannel := make(chan interface{})
	resultsChannel := make(chan interface{})

	go func() {
		str := ""
		for {
			select {
			case val := <-addChannel:
				vstr := ""
				switch sval := val.(type) {
				case string:
					vstr = sval
				case []byte:
					vstr = string(sval)
				default:
					vstr += fmt.Sprintf("%v", sval)
				}
				if vstr != "" {
					str += vstr + sep
				}
			case resultsChannel <- str:
				return
			}
		}
	}()

	return &stringValueReducer{addChannel, resultsChannel}
}

func (ar *stringValueReducer) AddValue(val interface{}) {
	ar.addChannel <- val
}

func (ar *stringValueReducer) Values() interface{} {
	return <-ar.resultsChannel
}
