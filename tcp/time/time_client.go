package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 256)
		_, err = conn.Read(buf)
		if err == io.EOF {
			break
		}
		io.WriteString(os.Stdout, fmt.Sprintf("message from server: %s", string(buf)))
	}
}