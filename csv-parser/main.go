package main

import (
	"bufio"
	"context"
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
	"time"
)

const (
	ReadBuff = 1000
)

var (
	config        = flag.String("c", "./config.yaml", "path to the configuration file")
	Header        []string
	IsString      = make(map[string]bool)
	FirstDataLine []uint8
	LineChan      = make(chan string, ReadBuff)
	FieldPos      = make(map[string]int)
	wg            = sync.WaitGroup{}
	ctx           = context.Background()
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

	config, err := parse.ConfigParse(configFile)
	if err != nil {
		panic(err.Error())
	}
	logger, err := log.NewLogger(config.LOGPATH)
	defer logger.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
	loggerErr, err := log.NewLogger(config.LOGERR)
	if err != nil {
		panic(err.Error())
	}
	defer loggerErr.Sync()
	logger.Info("Query " + strings.Join(Query, " ") + " received")
	fmt.Println(os.Args[0])
	csv, err := os.Open(config.CSVPATH)
	defer csv.Close()
	if err != nil {
		loggerErr.Error(err.Error())
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(config.SEARCHTIMEOUT)*time.Second)
	defer cancel()
	reader := bufio.NewReader(csv)
	RawHeader, _, err := reader.ReadLine()
	if err != nil {
		loggerErr.With(zap.String("Cannot retrieve  Header from csv %s", config.CSVPATH)).Error(err.Error())
		fmt.Fprintln(os.Stderr, "Cannot retrieve  Header from csv")
		os.Exit(1)
	}
	FirstDataLineRaw, _, err := reader.ReadLine()
	if err != nil {
		loggerErr.With(zap.String("Cannot retrieve first data line from csv %s", config.CSVPATH)).Error(err.Error())
		fmt.Fprintln(os.Stderr, "Cannot retrieve first data line from csv")
		os.Exit(1)
	}
	Header = (strings.Split(string(RawHeader), ","))
	FirstDataLine := strings.Split(string(FirstDataLineRaw), ",")
	GetFieldTypes(Header, FirstDataLine, IsString)
	if err = check.CheckQuerySyntax(Header, Query, Querylength, FieldPos); err != nil {
		loggerErr.Error(err.Error())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err = check.CheckQueryTypes(IsString, Query, Querylength); err != nil {
		loggerErr.Error(err.Error())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	logger.Info("Query " + strings.Join(Query, " ") + " checked, start  searching")
	LineChan <- string(FirstDataLineRaw)
	go ReadCsv(*reader, LineChan, logger, ctx)
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go worker(Header, Query, LineChan, Querylength, FieldPos)
	}
	wg.Wait()
}
func worker(Header []string, Query []string, ch <-chan string, Querylength int, FieldPos map[string]int) {
	parse.LineParse(Header, Query, LineChan, Querylength, FieldPos)
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

func ReadCsv(reader bufio.Reader, ch chan string, logger *zap.Logger, ctx context.Context) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():

			return
		default:
		}
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		} else if err != nil {
			logger.Error(err.Error())
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		ch <- string(line)
	}
}
