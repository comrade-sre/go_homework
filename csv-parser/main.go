package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/comrade-sre/go_homework/csv-parser/log"
	"github.com/comrade-sre/go_homework/csv-parser/check"
	"github.com/comrade-sre/go_homework/csv-parser/parse"
	"io"
	"os"
	"strconv"
	"strings"
	"go.uber.org/zap"

)

var (
	config   = flag.String("c", "./config.yaml", "path to the configuration file")
	sigChan  = make(chan os.Signal, 1)
	Header   []string
	IsString = make(map[string]bool)

)


func main() {
	flag.Parse()
	Query := flag.Args()
	Querylength := len(Query)
	configFile, err := os.Open(*config)
	if err != nil {
		panic(err.Error())
	}
	defer configFile.Close()
	config, err := parse.Parse(configFile)
	if err != nil {
		panic(err.Error())
	}
	logger, _ := log.NewLogger(config.LOGPATH)
	if err != nil {
		panic(err.Error())
	}
	defer logger.Sync()
	fmt.Println(os.Args[0])

	csv, err := os.Open(config.CSVPATH)
	defer csv.Close()
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
	LineChan := make(chan string, 100)
	reader := bufio.NewReader(csv)
	RawHeader, _, err := reader.ReadLine()
	 if err != nil {
	    logger.With(zap.String("Cannot retrieve  header from csv %s", config.CSVPATH)).Error(err.Error())
	    fmt.Fprintln(os.Stderr, "Cannot retrieve  header from csv")
	    os.Exit(1)


	}
	FirstDataLineRaw, _, err := reader.ReadLine()
    if err != nil {
	    logger.With(zap.String("Cannot retrieve first data line from csv %s", config.CSVPATH)).Error(err.Error())
	    fmt.Fprintln(os.Stderr, "Cannot retrieve first data line from csv")
	    os.Exit(1)
	}
	FirstDataLine := strings.Split(string(FirstDataLineRaw), ",")
	Header = (strings.Split(string(RawHeader), ","))
    GetFieldTypes(Header, FirstDataLine, IsString)
	if err = check.CheckQuerySyntax(Header, Query, Querylength); err != nil {
		logger.Error(err.Error())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err = check.CheckQueryTypes(IsString, Query, Querylength); err != nil {
	    logger.Error(err.Error())
	    fmt.Fprintln(os.Stderr, err.Error())
	    os.Exit(1)
	}
	go ReadCsv(*reader, LineChan)

	col := "Country/Region"
	index, err := check.CheckHeader(Header, col)
	pattern := "China"
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < 10; i++ {
		parse.ParseLine(index, pattern, LineChan)
	}

}

func GetFieldTypes(header []string, FirstDataLine []string, IsString map[string]bool) () {
	for index, field := range FirstDataLine {
		_, err := strconv.ParseFloat(field, 32)
		if err != nil {
			IsString[header[index]] = true
		} else {
			IsString[header[index]] = false
		}
	}
	return
}

func ReadCsv(reader bufio.Reader, ch chan string) (err error) {
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- string(line)
	}
	close(ch)
	return nil
}
