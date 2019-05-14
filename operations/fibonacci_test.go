package operations

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo_WrongInput(t *testing.T) {
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
	f := new(Fibo)
	f.GetInput()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestFibo_InputLessThan1(t *testing.T) {
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
	f := new(Fibo)
	f.GetInput()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestFibo_Input5(t *testing.T) {
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
	f := new(Fibo)
	f.GetInput()
	f.Execute()
	f.PrintOutput()

	result := map[int]int{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
		4: 3,
	}
	assert.Equal(t, result, f.memoize)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestFibo_Input13(t *testing.T) {
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
	f := new(Fibo)
	f.GetInput()
	f.Execute()
	f.PrintOutput()

	result := map[int]int{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  3,
		5:  5,
		6:  8,
		7:  13,
		8:  21,
		9:  34,
		10: 55,
		11: 89,
		12: 144,
	}
	assert.Equal(t, result, f.memoize)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
