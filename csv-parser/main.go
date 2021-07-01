package main

import (
	"bufio"
	//"errors"
	"flag"
	"fmt"
	"github.com/comrade-sre/go_homework/csv-parser/log"
	"github.com/comrade-sre/go_homework/csv-parser/parse"
	"io"
	"os"
	"strconv"
	"strings"
	//"reflect"
)

var (
	config    = flag.String("c", "./config.yaml", "path to the configuration file")
	sigChan   = make(chan os.Signal, 1)
	Header    []string
	IsString []bool
)

func main() {
	flag.Parse()
	Expression := flag.Args()
	configFile, err := os.OpenFile(*config, os.O_RDONLY, 0666)
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

	csv, err := os.OpenFile(config.CSVPATH, os.O_RDONLY, 0666)
	defer csv.Close()
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
	LineChan := make(chan string, 100)
	reader := bufio.NewReader(csv)
	RawHeader, _, err := reader.ReadLine()
	FirstDataLineRaw, _, err := reader.ReadLine()
    FirstDataLine := strings.Split(string(FirstDataLineRaw), ",")
	Header = (strings.Split(string(RawHeader), ","))


    IsString = GetTypes(Header, FirstDataLine)
	fmt.Println(IsString)
	go ReadCsv(*reader, LineChan)
	err = parse.ParseQuerySyntax(Header, Expression)
	if err != nil {
		panic(err.Error())
	}
	col := "Country/Region"
	index, err := parse.CheckHeader(Header, col)
	pattern := "China"
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < 10; i++ {
		parse.ParseLine(index, pattern, LineChan)
	}

}

func GetTypes(header []string, FirstDataLine []string) (IsString []bool) {
    for _, field := range FirstDataLine {
	    _, err := strconv.ParseFloat(field, 32)
		if err != nil {
			IsString = append(IsString,true)
		} else {
			IsString = append(IsString, false)
		}
	}
	return IsString
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
