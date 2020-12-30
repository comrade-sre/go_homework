package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {

		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz ", i)
		case i%5 == 0:
			fmt.Println("Buzz ", i)
		case i%3 == 0:
			fmt.Println("Fizz ", i)
		default:
			fmt.Println(i)
		}
	}
}
