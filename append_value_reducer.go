// A simple implementation of futures in go
package gofutures

type appendValueReducer struct {
	addChannel     chan interface{}
	resultsChannel chan interface{}
}

// Creates a ValueReducer that appends all values in a slice
func NewAppendValueReducer(capacity int) ValueReducer {
	addChannel := make(chan interface{})
	resultsChannel := make(chan interface{})

	go func() {
		interfaces := make([]interface{}, 0, capacity)
		for {
			select {
			case val := <-addChannel:
				interfaces = append(interfaces, val)
			case resultsChannel <- interfaces:
				return
			}
		}
	}()

	return &appendValueReducer{addChannel, resultsChannel}
}

func (ar *appendValueReducer) AddValue(val interface{}) {
	ar.addChannel <- val
}

func (ar *appendValueReducer) Values() interface{} {
	return <-ar.resultsChannel
}
