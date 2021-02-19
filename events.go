package main

type event struct {
	name      string
	listeners []func(args ...string)
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
func (ee EventEmitter) on(event string, listener func(args ...string)) {
	e := ee.findOrCreateEvent(event)
	e.listeners = append(e.listeners, listener)
}

// fire event
func (ee EventEmitter) emit(event string, args ...string) {
	if e, ok := ee.events[event]; ok {
		for _, listener := range e.listeners {
			listener(args...)
		}
	}
}
