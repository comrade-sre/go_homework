package main

import (
	"fmt"
	"os"
	"strconv"
)

func search(x int) []int {
	stack := []int{2}
	simple := []int{1}
	if x == 1 {
		return simple
	}
	for i := 3; i < x; i++ {
		for j := range stack {
			if i%stack[j] == 0 {
				stack = append(stack, i)
				break
			} else {
				stack = append(stack, i)
				simple = append(simple, i)
				break
			}
		}
	}
	return simple
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter the number")
		os.Exit(1)
	}
	num, numErr := strconv.Atoi(os.Args[1])
	if numErr != nil {
		fmt.Println(os.Args[1], "is not valid number")
		os.Exit(1)
	}
	fmt.Println(search(num))
}
