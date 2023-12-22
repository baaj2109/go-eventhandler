package main


type Event struct {
	Name string
	Data interface{}
	Cancelled bool
}