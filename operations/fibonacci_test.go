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
	sum := new(Fibo)
	sum.GetInput()

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
	sum := new(Fibo)
	sum.GetInput()

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

	result := map[int64]int64{
		int64(0): int64(0),
		int64(1): int64(1),
		int64(2): int64(1),
		int64(3): int64(2),
		int64(4): int64(3),
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

	result := map[int64]int64{
		int64(0):  int64(0),
		int64(1):  int64(1),
		int64(2):  int64(1),
		int64(3):  int64(2),
		int64(4):  int64(3),
		int64(5):  int64(5),
		int64(6):  int64(8),
		int64(7):  int64(13),
		int64(8):  int64(21),
		int64(9):  int64(34),
		int64(10): int64(55),
		int64(11): int64(89),
		int64(12): int64(144),
	}
	assert.Equal(t, result, f.memoize)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
