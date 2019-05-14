package operations

import (
	"fmt"
)

// Sum defines Sum operation struct
type Sum struct {
	a, b, result int64
}

// GetInput will get input for Sum operation
func (sum *Sum) GetInput() error {
	fmt.Println("Sum X & Y variables. Input format: x,y")
	fmt.Print("Input: ")

	_, err := fmt.Scanf("%d,%d\n", &sum.a, &sum.b)
	return err
}

// Execute will execute operation for Sum
func (sum *Sum) Execute() error {
	sum.result = sum.a + sum.b
	return nil
}

// PrintOutput will print result in console
func (sum *Sum) PrintOutput() error {
	fmt.Printf("Output: %d\n\n", sum.result)
	return nil
}
