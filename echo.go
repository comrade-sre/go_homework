package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var sep string = ";"
var s string

func simpleJoin() {
	// short versiont with the join function
	s1 := os.Args[1:]
	fmt.Println(strings.Join(s1, sep))
}
func simpleLoop() {
	// long verion with the loop
	arglength := len(os.Args) - 1
	for i := 1; i <= arglength; i++ {
		if i != arglength {
			s += os.Args[i] + sep
		} else {
			s += os.Args[i]
		}
	}
	fmt.Println(s)
}
func simpleRange() {
	// version with range, index printing  and time measure
	start := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
}

func main() {
	simpleJoin()
	simpleLoop()
	simpleRange()
}
