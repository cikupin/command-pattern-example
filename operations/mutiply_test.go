package operations

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	content := []byte("5,4")
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
	sum := new(Multiply)
	sum.GetInput()
	sum.Execute()
	sum.PrintOutput()

	assert.Equal(t, int64(20), sum.result)

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
