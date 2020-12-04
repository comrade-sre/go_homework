package main

import (
	"fmt"
	"math/rand"
	"os"
    "time"
)

func bubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
func insertSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		z := s[i]
		j := i
		for ; j >= 1 && s[j-1] > z; j-- {
			s[j] = s[j-1]
		}
		s[j] = z
	}
	return s
}

func main() {
	var length int
	fmt.Print("Enter the length of sequence to sort:")
	fmt.Fscan(os.Stdin, &length)
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = rand.Intn(1000)
	}
	fmt.Println(s)

	start := time.Now()
	fmt.Println(bubbleSort(s))
	duration := time.Since(start)
	fmt.Println("execution time fo bubble sorting", duration.Nanoseconds())

	start = time.Now()
	fmt.Println(insertSort(s))
	duration  = time.Since(start)
	fmt.Println("execution time for insert sort", duration.Nanoseconds())
}
