// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello World", Hello("World"))
	assert.NotEqual(t, "Hello Cats", Hello("World"))
	assert.NotEqual(t, "Hello Peaches", Hello("World"))
}

func TestCool(t *testing.T) {
	assert.Equal(t, "Code is cool", AddCool("Code"))
	assert.Equal(t, "Tech is cool", AddCool("Tech"))
	assert.NotEqual(t, "Malware is cool", AddCool("Software"))
}

func TestHello(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello")
	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, GetHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
