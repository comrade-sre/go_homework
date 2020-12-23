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
			isFile = printLines(counts, true)
			if isFile == false {
				fmt.Printf("in file %s was found duplicates\n", arg)
			}
			f.Close()

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
}
