package operations

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrime_WrongInput(t *testing.T) {
	content := []byte("qwe")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	os.Stdin = tmpfile

	// Test operation
	sum := new(Prime)
	sum.GetInput()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestPrime_InputLessThan1(t *testing.T) {
	content := []byte("0")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	os.Stdin = tmpfile

	// Test operation
	sum := new(Prime)
	sum.GetInput()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestPrime_Input5(t *testing.T) {
	content := []byte("5")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	os.Stdin = tmpfile

	// Test operation
	p := new(Prime)
	p.GetInput()
	p.Execute()
	p.PrintOutput()

	result := []int{2, 3, 5, 7, 9}
	assert.Equal(t, result, p.result)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestPrime_Input13(t *testing.T) {
	content := []byte("13")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	os.Stdin = tmpfile

	// Test operation
	p := new(Prime)
	p.GetInput()
	p.Execute()
	p.PrintOutput()

	result := []int{2, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25}
	assert.Equal(t, result, p.result)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
