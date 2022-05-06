package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadInput(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	text, err := reader.ReadString('\n')
	return text, err
}

func OutputUserInput(stdin io.Reader) (string, error) {
	str, err := ReadInput(stdin)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func main() {

	fmt.Println("Say something")
	result, err := OutputUserInput(os.Stdin)
	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
		os.Exit(1)
	}
	fmt.Printf("\nYou said: %s\n", result)
}

// Credits to this post on testing command line user input https://petersouter.xyz/testing-and-mocking-stdin-in-golang/
