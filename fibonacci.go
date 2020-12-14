package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var calculated = make(map[int]int)

func fibonacci(number int) int {
	if number <= 1 {
		return number
	} else if val, ok := calculated[number]; ok {
		return val
	} else {
		result := fibonacci(number-1) + fibonacci(number-2)
		calculated[number] = result
		return result
	}
}
func main() {
	_, file, _, _ := runtime.Caller(1)
	if len(os.Args) < 2 {
		scriptName := strings.Split(file, "/")
		fmt.Printf("usage: %s int positive number\n", scriptName[len(scriptName)-1])
		os.Exit(1)
	}
	var number int
	number, numErr := strconv.Atoi(os.Args[1])
	if numErr != nil {
		fmt.Printf("%s is incorrect number\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println(fibonacci(number))
}
