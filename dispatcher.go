package main

type Dispatcher struct {
	listeners map[string][]EventListener
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[string][]EventListener),
	}
}

func (d *Dispatcher) RegisterListener(name string, listener EventListener) {
	d.listeners[name] = append(d.listeners[name], listener)
}

func (d *Dispatcher) Dispatch(event Event) {
	for _, listener := range d.listeners[event.Name] {
		listener(event)
	}
	// for _, listener := range d.listeners[event.Name] {
	// 	if event.Cancelled {
	// 		break
	// 	}
	// 	if err := listener(event); err != nil {
	// 		break
	// 	}
	// }
}

// async dispatch
func (d *Dispatcher) DispatchAsync(event Event) {
	for _, listener := range d.listeners[event.Name] {
		go listener(event)
	}
}
