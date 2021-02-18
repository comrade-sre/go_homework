package main

import (
	"log"
	"math/rand"
	"time"
	"fmt"
	"os"
)

func main() {
	var err error
	err = NewError("custom error")
	fmt.Println(err)
	f, err := os.OpenFile("test", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	defer func() {
		if v := recover(); v != nil {
			log.Println("capture a panic:", v)
			log.Println("avoid crashing the program")
		}
	}()
	intentPanic()
	
}

type ErrorDate struct {
	message string
	date   string
}

func NewError(text string) error {
	return &ErrorDate {
		message: text,
		date:    time.Now().Format(time.RFC850),
	}
}
func (e *ErrorDate) Error() string {
	return fmt.Sprintf("error: %s\t%s", e.date, e.message)
}
func intentPanic() {
	var i int
	for i = 1000; i > -1; i-- {
		println(rand.Intn(1000) / i)

	}
}
