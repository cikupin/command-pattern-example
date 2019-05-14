package operations

import (
	"fmt"
	"os"
)

// Multiply defines number multiplication operation struct
type Multiply struct {
	a, b, result int64
}

// GetInput will get input for multiplication operation
func (m *Multiply) GetInput() error {
	fmt.Println("Multiply X & Y variables. Input format: x,y")
	fmt.Print("Input: ")

	_, err := fmt.Fscanf(os.Stdin, "%d,%d\n", &m.a, &m.b)
	return err
}

// Execute will execute operation for multiplication
func (m *Multiply) Execute() error {
	m.result = m.a * m.b
	return nil
}

// PrintOutput will print result in console
func (m *Multiply) PrintOutput() error {
	fmt.Printf("Output: %d\n\n", m.result)
	return nil
}
