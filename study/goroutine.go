package main

import (
	"fmt"
	"runtime/trace"
	"time"
	"os"
)

const count = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var ch = make(chan struct{}, count)
	var num int
	for i := 0; i < count; i++ {
		go func() {
			num++
			ch <- struct{}{}
		}()
	}
	time.Sleep(2 * time.Second)
	close(ch)
	var i int
	for range ch {
		i++
	}
	fmt.Printf("%d\n%d\n", num, i)
}
