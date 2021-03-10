package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	err     error
	count   int
	message string
	mutex   sync.Mutex
)

func main() {

	message = "there is no appropiate number for goroutines"
	if len(os.Args) != 2 {
		println(message)
		os.Exit(1)
	}
	count, err = strconv.Atoi(os.Args[1])
	if err != nil {
		println(message)
		os.Exit(1)
	}
	ch := make(chan struct{}, count)
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			ch <- struct{}{}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	var i int
	for range ch {
		i += 1
	}
	fmt.Printf("there were %d goroutines\n", i)
	mutextTest(&mutex)
}

func mutextTest(m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	println("do something and unlock mutext")
}
