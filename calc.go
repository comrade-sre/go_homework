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
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
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
