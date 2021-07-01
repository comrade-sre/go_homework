package main

import (
	"fmt"
	"os"
	"flag"
    "github.com/comrade-sre/go_homework/csv-parser/parse"
    "github.com/comrade-sre/go_homework/csv-parser/log"
	"io"
    "strings"
	"bufio"
	"errors"

)
type Expression []string
func (exp *Expression) String() string {
    var buff strings.Builder
    buff.Grow(32)
    for _, val := range *exp {
        fmt.Fprintf(&buff,"%s", val)
    }
    return buff.String()

}
func (exp *Expression) Set(value string) error{
    *exp = append(*exp, value)
    return nil

}
var (
	config = flag.String("c", "./config.yaml", "path to the configuration file")
	sigChan = make(chan os.Signal, 1)
    query Expression
)

func main() {
	flag.Var(&query, "", "query string")
	flag.Parse()
	fmt.Println(*query)

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

	csv, err := os.OpenFile(config.CSVPATH,os.O_RDONLY, 0666)
	defer csv.Close()
	if err != nil {
        logger.Error(err.Error())
        panic(err.Error())
	}
	LineChan := make(chan string, 100)
	reader := bufio.NewReader(csv)
	RawHeader, _, err := reader.ReadLine()
	Header := (strings.Split(string(RawHeader), ","))
	go ReadCsv(*reader, LineChan)
	column := "Country/Region"
	pattern := "US"
    index, err := CheckHeader(Header, column)
    if err != nil {
        panic(err.Error())
    }
	for i := 0; i < 10; i++ {
	    ParseLine(index, pattern, LineChan)
	}


}
func CheckHeader(header []string, column string) (index int, err error) {
	    for index, value := range header {
	        if value == column {
	            return index, nil
	        }
	    }
	    err = errors.New("such column not found")
	    return 0, err
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
func ParseLine(index int, pattern string, ch <-chan string) (err error) {
    for line := range ch {
        values := strings.Split(line, ",")
        if values[index] == pattern {
        fmt.Println(line)
        }
    }
    return nil
}
