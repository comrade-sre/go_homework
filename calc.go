package main

import (
	"fmt"
	"os"
	"strconv"
)

func sub(x int, y int) int {
	return x - y
}
func mult(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}

func add(x int, y int) int {
	return x + y
}

func main() {
	a, aErr := strconv.Atoi(os.Args[1])
	b, bErr := strconv.Atoi(os.Args[2])
	if aErr != nil || bErr != nil {
		fmt.Println(bErr, aErr)
		os.Exit(2)
	}
	switch op := os.Args[3]; op {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(sub(a, b))
	case "*":
		fmt.Println(mult(a, b))
	case "/":
		fmt.Println(div(a, b))
	}
}
