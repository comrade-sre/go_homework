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
		os.Exit(1)
	}
	fmt.Printf("%+v\n", config)

	defer func() {
		err = logFile.Close()
		if err != nil {
			log.Printf("Can't close file: %v", err)
		}
		err = configFile.Close()
		if err != nil {
			log.Printf("Can't close file: %v", err)
		}
	}()
}
