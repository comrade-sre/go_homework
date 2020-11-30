package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var sep string = ";"
var s string

func simple_join() {
	// short versiont with the join function
	s1 := os.Args[1:]
	fmt.Println(strings.Join(s1, sep))
}
func simple_loop() {
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
func simple_range() {
	// version with range, index printing  and time measure
	start := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
}

func main() {
	simple_join()
	simple_loop()
	simple_range()
}
