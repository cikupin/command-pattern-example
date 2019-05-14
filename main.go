package main

import (
	"github.com/cikupin/command-pattern-example/operations"
)

// IEvent defines interface for every event
type IEvent interface {
	GetInput() error
	Execute() error
	PrintOutput() error
}

// EventStore defines storage to store assigned events
type EventStore struct {
	events []IEvent
}

// InitAndExecute will initialize event and execute event command
func (e *EventStore) InitAndExecute(event IEvent) {
	e.events = append(e.events, event)
	err := event.GetInput()
	if err != nil {
		panic(err)
	}

	err = event.Execute()
	if err != nil {
		panic(err)
	}
	event.PrintOutput()
}

func main() {
	sumOperation := new(operations.Sum)
	multiplicationOperation := new(operations.Multiply)
	fibonacciOperations := new(operations.Fibo)

	eventStorage := new(EventStore)
	eventStorage.InitAndExecute(sumOperation)
	eventStorage.InitAndExecute(multiplicationOperation)
	eventStorage.InitAndExecute(fibonacciOperations)
}
