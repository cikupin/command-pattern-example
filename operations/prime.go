package operations

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

// Prime defines struct for prime number operation
type Prime struct {
	result []int
	input  int
}

// GetInput will get input for prime number operation
func (p *Prime) GetInput() error {
	fmt.Println("Find first N prime number. Input format: x")
	fmt.Print("Input: ")

	_, err := fmt.Fscanf(os.Stdin, "%d\n", &p.input)
	if err != nil {
		return err
	}

	if p.input < 1 {
		return errors.New("input cannot be less than 1")
	}
	return nil
}

// Execute will execute operation for prime number
func (p *Prime) Execute() error {
	num := 2
	for {
		primeStatus := p.checkPrime(num)
		if primeStatus {
			p.result = append(p.result, num)
		}

		if len(p.result) == p.input {
			break
		}
		num++
	}
	return nil
}

// checkPrime will check if given number is prime or not
func (p *Prime) checkPrime(n int) bool {
	floatN := float64(n)
	limit := int(math.Sqrt(floatN)) + 1
	for i := 2; i < limit; i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true

}

// PrintOutput will print result in console
func (p *Prime) PrintOutput() error {
	output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(p.result)), ", "), "[]")
	fmt.Printf("Output: %s\n\n", output)
	return nil
}
