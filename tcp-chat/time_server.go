package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var messageChan = make(chan string)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			if message != "" {
				messageChan <- message
			}
		}
	}()
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, messageChan)
		fmt.Println("conn processing finished")
	}
}

func handleConn(c net.Conn, ch <-chan string) {
	defer c.Close()
	defer close(messageChan)
	for {
		select {
		case input := <-ch:
			_, err := io.WriteString(c, input)
			if err != nil {
				return
			}
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}

}
