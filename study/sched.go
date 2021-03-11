package main

import (
	"fmt"
	"runtime"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	fmt.Println("hello")
	runtime.Gosched()
	fmt.Println("world")
}
