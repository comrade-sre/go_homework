package searchSimple

import (
	"fmt"
	"os"
	"strconv"
)

func search(x int) (result []int) {
	stack := []int{2}
	var isSimple bool
	if x == 1 {
		result = append(result, 1)
		return
	}
	for i := 3; i < x; i++ {
		for _, j := range stack {
			isSimple = true
			if i%j == 0 {
				stack = append(stack, i)
				isSimple = false
				break
			}
		}
		stack = append(stack, i)
		if isSimple {
			result = append(result, i)
		}
	}
	return result
}

//func main() {
//	if len(os.Args) < 2 {
//		fmt.Println("enter the number")
//		os.Exit(1)
//	}
//	num, numErr := strconv.Atoi(os.Args[1])
//	if numErr != nil || num < 1 {
//		fmt.Println("enter positive number")
//		os.Exit(1)
//	}
//	fmt.Println(search(num))
//}
