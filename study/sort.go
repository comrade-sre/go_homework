package main

import (
	"fmt"
	"math/rand" //nolint
	"os"
	"time"
)

func bubbleSort(sliceToSort []int) []int {
	for index := 0; index < len(sliceToSort); index++ {
		for tmpIndex := index + 1; tmpIndex < len(sliceToSort); tmpIndex++ {
			if sliceToSort[index] > sliceToSort[tmpIndex] {
				sliceToSort[index], sliceToSort[tmpIndex] = sliceToSort[tmpIndex], sliceToSort[index]
			}
		}
	}
	return sliceToSort
}
func insertSort(sliceToSort []int) []int {
	for index := 1; index < len(sliceToSort); index++ {
		compareValue := sliceToSort[index]
		tmpIndex := index
		for ; tmpIndex >= 1 && sliceToSort[tmpIndex-1] > compareValue; tmpIndex-- {
			sliceToSort[tmpIndex] = sliceToSort[tmpIndex-1]
		}
		sliceToSort[tmpIndex] = compareValue
	}
	return sliceToSort
}

func main() {
	var length int
	fmt.Print("Enter the length of sequence to sort:")
	fmt.Fscan(os.Stdin, &length)
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = rand.Intn(1000) //nolint
	}
	fmt.Println(s)

	start := time.Now()
	fmt.Println(bubbleSort(s))
	duration := time.Since(start)
	fmt.Println("execution time fo bubble sorting", duration.Nanoseconds())

	start = time.Now()
	fmt.Println(insertSort(s))
	duration = time.Since(start)
	fmt.Println("execution time for insert sorting", duration.Nanoseconds())
}
