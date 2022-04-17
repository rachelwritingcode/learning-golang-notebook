// main.go
package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello " + name
}

func AddCool(str string) string {
	return str + " is cool"
}

func main() {
	fmt.Println(Hello("World"))
	fmt.Println(AddCool("Code"))
}
