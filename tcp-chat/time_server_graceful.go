package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	messageChan = make(chan string)
	sigChan     = make(chan os.Signal)
)

const (
	CONN_TYPE = "tcp"
	CONN_HOST = "localhost"
	CONN_PORT = "8000"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	listener, err := net.Listen(CONN_TYPE, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				if message != "" {
					messageChan <- message
				}
			}
		}
		close(messageChan)
	}()
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			listener.Close()
			break
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			wg.Add(1)
			go handleConn(wg, ctx, conn, messageChan)
		}
	}
}

func handleConn(wg sync.WaitGroup, ctx context.Context, c net.Conn, ch <-chan string) {
	defer 	c.Close()
	for {
		select {
		case <-ctx.Done():
			_, err := io.WriteString(c, "BYE!")
			if err != nil {
				log.Println("Unable to write to the connection")
			}
			c.Close()
			wg.Done()
			return
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
