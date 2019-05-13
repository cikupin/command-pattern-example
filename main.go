package main

import (
	"github.com/cikupin/command-pattern-example/operations"
)

// IEvent defines interface for every event
type IEvent interface {
	GetInput() error
	Execute() error
}

// EventStore defines storage to store assigned events
type EventStore struct {
	events []IEvent
}

// InitAndExecute will initialize event and execute event command
func (e *EventStore) InitAndExecute(event IEvent) {
	e.events = append(e.events, event)
	event.GetInput()
	event.Execute()
}

func main() {
	sumOperation := new(operations.Sum)
	multiplicationOperation := new(operations.Multiply)

	eventStorage := new(EventStore)
	eventStorage.InitAndExecute(sumOperation)
	eventStorage.InitAndExecute(multiplicationOperation)
}
