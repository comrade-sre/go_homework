// Package implements simple exmaple for opening files, also parsing yaml file configuration
// in addition it's example for logging process in golang, also error processing during file operations
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
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.Println("web server starting")
	flag.Parse()
	filename := strings.TrimSpace(*config)
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("Can't check file existing: %v", err)
	}
	configFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open file: %v\n", err)
	}

	config, err := parse.Parse(configFile)
	if err != nil {
		log.Fatalf("parsing failed: %v", err)
	}
	fmt.Printf("%+v\n", config)

	err = logFile.Close()
	if err != nil {
		log.Printf("Can't close file: %v", err)
	}
	err = configFile.Close()
	if err != nil {
		log.Printf("Can't close file: %v", err)
	}
}
