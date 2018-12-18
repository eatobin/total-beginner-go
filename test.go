package main

import (
	"fmt"
)

type DetailedError struct {
	x, y int
}

func (e DetailedError) Error() string {
	return fmt.Sprintf("Error occured at (%v,%v)", e.x, e.y)
}

func foo(answer, x, y int) (int, error) {
	if answer == 42 {
		return 100, nil
	}
	return 0, DetailedError{x: x, y: y}
}

func main() {
	r, err := foo(42, 10, 20)
	if err != nil {
		fmt.Printf("Error 1 occured: %v\n", err)
		return
	}
	fmt.Println("Answer: ", r)
	r1, err := foo(43, 30, 40)
	if err != nil {
		fmt.Printf("Error 2 occured: %v\n", err)
		return
	}
	fmt.Println(r1)
}
