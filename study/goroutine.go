package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var num int
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			num++
			mutex.Unlock()
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println(num)
}
