package main

import (
	"flag"
	"fmt"
	"github.com/comrade-sre/go_homework/parse"
	"log"
	"os"
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

	fmt.Println(parse.Parse(filename))
}
