package main

import "fmt"

type event struct {
	name      string
	listeners []func(args ...interface{}) (interface{}, error)
}

func newEvent(name string) *event {
	e := event{name: name}
	return &e
}

// EventEmitter is struct for managing events
type EventEmitter struct {
	events map[string]*event
}

func (ee EventEmitter) findOrCreateEvent(event string) *event {
	if _, ok := ee.events[event]; ok == false {
		ee.events[event] = newEvent(event)
	}
	return ee.events[event]
}

// register new listener
func (ee EventEmitter) on(event string, listener func(args ...interface{}) (interface{}, error)) {
	e := ee.findOrCreateEvent(event)
	e.listeners = append(e.listeners, listener)
}

// fire event
func (ee EventEmitter) emit(event string, args ...interface{}) {
	if event == "error" {
		if _, ok := ee.events[event]; ok == false {
			fmt.Println(args...)
			panic("no listener registered")
		}
	}

	if e, ok := ee.events[event]; ok {
		for _, listener := range e.listeners {
			_, err := listener(args)
			if err != nil {
				ee.emit("error", err)
			}
		}
	}
}
