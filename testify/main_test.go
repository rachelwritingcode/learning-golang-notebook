// main_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "Hello World", Hello("World"))
	assert.NotEqual(t, "Hello Cats", Hello("World"))
	assert.NotEqual(t, "Hello Peaches", Hello("World"))
}

func TestCool(t *testing.T) {
	assert.Equal(t, "Code is cool", AddCool("Code"))
	assert.Equal(t, "Tech is cool", AddCool("Tech"))
	assert.NotEqual(t, "Malware is cool", AddCool("Software"))
}
