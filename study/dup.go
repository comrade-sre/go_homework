package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var isFile bool
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		printLines(counts, false)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// isFile flag for understnading, that file was passed to the fuction
			isFile = printLines(counts, true)
			if !isFile  {
				// we set isFile to false again, when we found duplicates in file
				fmt.Printf("in file %s was found duplicates\n", arg)
			}
			err = f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "there is an error occured while closing file:%v\n")
			}

		}
	}
}
func printLines(counts map[string]int, isFile bool) bool {
	for line, n := range counts {
		if n > 1 {
			isFile = false
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	for key := range counts {
		delete(counts, key)
	}
	return isFile
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standart input: %v\n", err)
	}
}
