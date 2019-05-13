package operations

import (
	"fmt"
)

// Multiply defines number multiplication operation struct
type Multiply struct {
	a, b int32
}

// GetInput will get input for multiplication operation
func (m *Multiply) GetInput() error {
	fmt.Println("Multiply X & Y variables. Input format: x,y")
	fmt.Print("Input: ")
	fmt.Scanf("%d,%d\n", &m.a, &m.b)
	return nil
}

// Execute will execute operation for multiplication
func (m *Multiply) Execute() error {
	result := m.a * m.b
	fmt.Printf("Output: %d\n\n", result)
	return nil
}
