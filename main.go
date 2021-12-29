package main

import (
	"errors"
	"fmt"
	"log"
)

type server struct {
	connection    string
	connected     bool
	eventsEmitter EventEmitter
}

func (server server) init() {
	// Register all listeners
	server.eventsEmitter.on("connect", handleConnection)
	server.eventsEmitter.on("connect", logMessage)
	server.eventsEmitter.on("data", logMessage)
	server.eventsEmitter.on("error", errorHandler)
	server.eventsEmitter.on("disconnect", handleDisonnection)
}

// Function for connecting to server
func (server *server) connect(connectionString string) {
	server.connection = connectionString
	server.connected = true
	server.eventsEmitter.emit("connect", server.connection, "Connected successfully")
}

func (server *server) data(userData string) {
	if server.connected {
		server.eventsEmitter.emit("data", server.connection, userData)
	}
}

// Function for disconnecting from server
func (server *server) disconnect() {
	if server.connected {
		server.connected = false
		server.eventsEmitter.emit("disconnect", server.connection, "Disconnected successfully")
	} else {
		fmt.Print("Already disconnected")
	}
}

var handleConnection = func(args ...interface{}) (interface{}, error) {
	values := args[0].([]interface{})
	fmt.Println(values[0], values[1])
	return nil, nil
}

var handleDisonnection = func(args ...interface{}) (interface{}, error) {
	return nil, errors.New("failed to disconnect")
}

var errorHandler = func(args ...interface{}) (interface{}, error) {
	errors := args[0].([]interface{})
	for _, error := range errors {
		fmt.Println("ERROR!", error)
	}
	return nil, nil
}

var logMessage = func(args ...interface{}) (interface{}, error) {
	values := args[0].([]interface{})
	log.Print(values[1])
	return nil, nil
}

func main() {
	var s = server{
		eventsEmitter: EventEmitter{
			events: make(map[string]*event),
		},
	}
	s.init()

	s.connect("http/connection/string")

	s.data("Message from user 1")
	s.data("Message from user 2")
	s.data("Message from user 3")
	s.data("Message from user 4")

	s.disconnect()
}
