package operations

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Fibo defines struct for Fibonacci operation
type Fibo struct {
	memoize map[int]int
	input   int
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

	f.memoize = map[int]int{}
	return nil
}

// Execute will execute operation for fibonacci
func (f *Fibo) Execute() error {
	for i := 0; i < f.input; i++ {
		f.getFibo(i)
	}
	return nil
}

// getFibo is a recursion function to get fibonacci number
func (f *Fibo) getFibo(n int) int {
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
	ls := make([]int, 0, len(f.memoize))
	for _, val := range f.memoize {
		ls = append(ls, val)
	}

	output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ls)), ", "), "[]")
	fmt.Printf("Output: %s\n\n", output)
	return nil
}
