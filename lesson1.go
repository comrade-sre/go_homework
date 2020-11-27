package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sep string = ";"
	var s string
	s1 := os.Args[1:]
	arglength := len(os.Args) - 1
	// short variant with the join function
	fmt.Println(strings.Join(s1, sep))
	// long variant with the loop
	for i := 1; i <= arglength; i++ {
		if i != arglength {
			s += os.Args[i] + sep
		} else {
			s += os.Args[i]
		}
	}
	fmt.Println(s)
}
