package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"parse"
	"strings"
)

var (
	config = flag.String("config", "", "configuration file for app")
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println(flag.NArg())
		log.Fatal("There is incorrect number of arguments, need file name")
	}
	filename := strings.TrimSpace(flag.Arg(0))
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("Can't check file existing: %v", err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Can't close file: %v", err)
		}
	}()

	parse.Parse(file)
}
