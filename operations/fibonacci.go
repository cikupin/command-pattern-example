package operations

import (
	"errors"
	"fmt"
	"os"
)

// Fibo defines struct for Fibonacci operation
type Fibo struct {
	memoize map[int64]int64
	input   int64
}

// GetInput will get input for fibonacci operation
func (f *Fibo) GetInput() error {
	fmt.Println("Find first N Fibonacci sequence. Input format: x")
	fmt.Print("Input: ")

	_, err := fmt.Fscanf(os.Stdin, "%d\n", &f.input)
	if err != nil {
		return err
	}

	if f.input < 1 {
		return errors.New("input cannot be less than 1")
	}

	f.memoize = map[int64]int64{}
	return nil
}

// Execute will execute operation for fibonacci
func (f *Fibo) Execute() error {
	for i := int64(0); i < f.input; i++ {
		f.getFibo(i)
	}
	return nil
}

// getFibo is a recursion function to get fibonacci number
func (f *Fibo) getFibo(n int64) int64 {
	if n == 0 {
		f.memoize[0] = 0
		return f.memoize[0]
	} else if n == 1 {
		f.memoize[1] = 1
		return f.memoize[1]
	} else {
		if _, ok := f.memoize[n]; ok {
			return f.memoize[n]
		}

		val := f.getFibo(n-1) + f.getFibo(n-2)
		f.memoize[n] = val
		return val
	}
}

// PrintOutput will print result in console
func (f *Fibo) PrintOutput() error {
	fmt.Print("Output: ")
	for i := int64(0); i < f.input; i++ {
		fmt.Printf("%d ", f.memoize[i])
	}
	fmt.Print("\n\n")
	return nil
}
