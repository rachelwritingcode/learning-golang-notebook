package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputUserInput(t *testing.T) {

	var stdin bytes.Buffer
	stdin.Write([]byte("hello\n"))
	result, err := ReadInput(&stdin)
	assert.NoError(t, err)
	assert.Equal(t, "hello\n", result)
}
