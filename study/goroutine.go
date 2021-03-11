package main

import (
	"fmt"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

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
