package main

import (
	"fmt"
	"time"
)

func main() {
	var num int
	for i := 0; i < 1001; i++ {
		go func() {
			num++
		}()
	}
	time.Sleep(10 * time.Second)
	fmt.Println(num)
}
