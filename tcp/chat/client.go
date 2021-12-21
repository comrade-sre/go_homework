package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var nickName string

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if len(os.Args) == 1 {
		nickName = "noName"
	} else {
		nickName = os.Args[1]
	}
	fmt.Fprintf(conn, nickName)
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin)
	fmt.Printf("%s: exit", conn.LocalAddr())
}
