package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/comrade-sre/go_homework/csv-parser/check"
	"github.com/comrade-sre/go_homework/csv-parser/log"
	"github.com/comrade-sre/go_homework/csv-parser/parse"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	ReadBuff = 1000
)

var (
	config        = flag.String("c", "./config.yaml", "path to the configuration file")
	sigChan       = make(chan os.Signal, 1)
	Header        []string
	IsString      = make(map[string]bool)
	FirstDataLine []uint8
	LineChan      = make(chan string, ReadBuff)
	FieldPos      = make(map[string]int)
	wg            = sync.WaitGroup{}
)

func main() {
	flag.Parse()
	Query := flag.Args()
	Querylength := len(Query)
	configFile, err := os.Open(*config)
	defer configFile.Close()
	if err != nil {
		panic(err.Error())
	}

	config, err := parse.Parse(configFile)
	if err != nil {
		panic(err.Error())
	}
	logger, err := log.NewLogger(config.LOGPATH)
	defer logger.Sync()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(os.Args[0])
	csv, err := os.Open(config.CSVPATH)
	defer csv.Close()
	if err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
	reader := bufio.NewReader(csv)
	RawHeader, _, err := reader.ReadLine()
	if err != nil {
		logger.With(zap.String("Cannot retrieve  Header from csv %s", config.CSVPATH)).Error(err.Error())
		fmt.Fprintln(os.Stderr, "Cannot retrieve  Header from csv")
		os.Exit(1)
	}
	FirstDataLineRaw, _, err := reader.ReadLine()
	if err != nil {
		logger.With(zap.String("Cannot retrieve first data line from csv %s", config.CSVPATH)).Error(err.Error())
		fmt.Fprintln(os.Stderr, "Cannot retrieve first data line from csv")
		os.Exit(1)
	}
	Header = (strings.Split(string(RawHeader), ","))
	FirstDataLine := strings.Split(string(FirstDataLineRaw), ",")
	GetFieldTypes(Header, FirstDataLine, IsString)
	if err = check.CheckQuerySyntax(Header, Query, Querylength, FieldPos); err != nil {
		logger.Error(err.Error())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err = check.CheckQueryTypes(IsString, Query, Querylength); err != nil {
		logger.Error(err.Error())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	go ReadCsv(*reader, LineChan, logger)
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go worker(Header, Query, LineChan, Querylength, FieldPos)
	}
	wg.Wait()
}
func worker(Header []string, Query []string, ch <-chan string, Querylength int, FieldPos map[string]int) {
	parse.ParseLine(Header, Query, LineChan, Querylength, FieldPos)
	wg.Done()
}
func GetFieldTypes(Header []string, FirstDataLine []string, IsString map[string]bool) {
	for index, field := range FirstDataLine {
		_, err := strconv.ParseFloat(field, 32)
		if err != nil {
			IsString[Header[index]] = true
		} else {
			IsString[Header[index]] = false
		}
	}
	return
}

func ReadCsv(reader bufio.Reader, ch chan string, logger *zap.Logger) {
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			close(ch)
			break
		} else if err != nil {
			logger.Error(err.Error())
			fmt.Fprintln(os.Stderr, err.Error())
			close(ch)
			break
		}
		ch <- string(line)
	}
}
