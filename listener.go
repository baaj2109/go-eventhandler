package main

type EventListener func(Event)

func NewEventListener(listener EventListener) EventListener {
	return listener
}
