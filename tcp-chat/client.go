package main

import (
	"fmt"
	//"golang.org/x/text/secure/precis"
	"io"
	"log"
	"net"
	"os"
)
var nickName string
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if len(os.Args) == 1 {
		nickName = "noName"
	} else {
		nickName = os.Args[1]
	}

	fmt.Fprintln(conn, nickName)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin)
	fmt.Printf("%s: exit", conn.LocalAddr())
}