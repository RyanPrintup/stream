package stream

type Listener struct {
	EventType string
	Callback func(args... interface{})
}