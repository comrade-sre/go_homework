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
const (
	CONN_TYPE = "tcp"
	CONN_HOST = "localhost"
	CONN_PORT = "8000"
)
func main() {
	listener, err := net.Listen(CONN_TYPE, fmt.Sprintf("s%:s%",CONN_HOST,CONN_PORT)
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
