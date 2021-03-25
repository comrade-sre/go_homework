package main

import (
	"fmt"
	"os"
	"os/signal"
	//"context"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	//ctx := context.Background()

	go func() {
	    time.Sleep(5 * time.Second)
		signal.Notify(c)

		s := <-c
		fmt.Println("Got signal:", s)

	}()

}
