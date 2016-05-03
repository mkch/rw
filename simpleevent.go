package rw

type simpleEvent struct {
	sender interface{}
}

func (evt *simpleEvent) Sender() interface{} {
	return evt.sender
}
