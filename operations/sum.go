package operations

import (
	"fmt"
)

// Sum defines Sum operation struct
type Sum struct {
	a, b int32
}

// GetInput will get input for Sum operation
func (sum *Sum) GetInput() error {
	fmt.Println("Sum X & Y variables. Input format: x,y")
	fmt.Print("Input: ")
	fmt.Scanf("%d,%d\n", &sum.a, &sum.b)
	return nil
}

// Execute will execute operation for Sum
func (sum *Sum) Execute() error {
	result := sum.a + sum.b
	fmt.Printf("Output: %d\n\n", result)
	return nil
}
