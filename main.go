package main

import "fmt"

func main() {
	dispatcher := NewDispatcher()

	// register a listener
	listener := NewEventListener(func(event Event) {
		// do something
		fmt.Printf("event: %s received \n", event.Name)
	})
	dispatcher.RegisterListener("testEvent", listener)


	// trigger the event
	dispatcher.Dispatch(Event{Name: "test event"})
}
