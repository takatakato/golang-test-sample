package main

import (
	"errors"
	"fmt"
)

func example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}

func exampleDuplicate(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}

func lintNG(bar int) int {
	if bar > 0 {
		return 123
	} else {
		return 456
	}
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")
}
