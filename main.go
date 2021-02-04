package main

import (
	"fmt"
	"log"
)

func main() {

	e, err := initializeEvent("Welcome to Sensedia")

	if err != nil {
		log.Fatal("failed to initialize event")
	}

	e.start()
}

func newMessage(msg string) message {
	return message{value: msg}
}

type message struct {
	value string
}

func (m message) String() string {
	return fmt.Sprintf("[OFFICIAL MESSAGE]: %s", m.value)
}

func newGreeter(m message) greeter {
	return greeter{message: m}
}

type greeter struct {
	message message
}

func (g greeter) greet() string {
	return g.message.String()
}

func newEvent(g greeter) (event, error) {
	return event{greeter: g}, nil
}

type event struct {
	greeter greeter
}

func (e event) start() {
	fmt.Println(e.greeter.greet())
}
