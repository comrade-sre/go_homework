package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("test", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	intentPanic()

}

type ErrorDate struct {
	message string
	date    string
}

func NewError(text string) error {
	return &ErrorDate{
		message: text,
		date:    time.Now().Format(time.RFC850),
	}
}

func (e *ErrorDate) Error() string {
	return fmt.Sprintf("error: %s\t%s", e.date, e.message)
}

func intentPanic() {
	defer func() {
		if v := recover(); v != nil {
			log.Println("capture a panic:", v)
		}
	}()
	panic(NewError("custom error"))
}
