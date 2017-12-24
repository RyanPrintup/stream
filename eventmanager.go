package stream

type EventManager struct {
	listeners map[interface{}][]Listener
	async bool

}

func NewEventManager(async bool) *EventManager {
	return &EventManager{ make(map[string][]Listener), async }
}

// Alias to Register()
func (e *EventManger) OnEvent(eventType string, listener Listener) {
	e.Register(eventType, listener)

	return
}

func (e *EventManager) Register(eventType string, listener Listener) {
	if _, ok := e.listeners[event]; !ok {
		e.listeners[event] = []Listener{}
	}

	e.listeners[event] = append(e.listeners[event], callback)

	return
}

func (e *EventManager) Unregister(eventType string, listener Listener) {
	if _, ok := e.listeners[eventType]; !ok {
		return
	}

	for k, v := range(e.listeners[eventType]) {
		if &v == listener {
			e.listeners[eventType] = append(e.listeners[eventType][:k], e.listeners[eventType][k + 1]...)

			break
		} 
	}

	if (len(e.listeners[eventType]) == 0) {
		delete(e.listeners, eventType)
	}

	return
}

func (e *EventManager) Dispatch(eventType string, args ..interface{}) {
	if _, ok := e.listeners[eventType]; !ok {
		return
	}

	for _, v := range(e.listeners[eventtType]) {
		if e.async {
			v.Callback(args...)
		} else {
			go v.Callback(args...)
		}
	}

	return
}