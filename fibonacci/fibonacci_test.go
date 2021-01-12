package test
import (
	"math/rand"
	"fmt"
)
func TestFibonacci(t *testing.T) {
	
	fib := [30]{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040}

      var number int
      number = rand.Intn(30)
      if numErr != nil {
              fmt.Printf("%s is incorrect number\n", os.Args[1])
              os.Exit(1)
      }
	result := fibonacci(number)
	if result != fib[number] {
		t.Errorf("fibonacci value %s, wants %s", result,fib[number])	
	}
	
}	
