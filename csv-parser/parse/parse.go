package parse

import (
    "gopkg.in/yaml.v2"
    "strings"
    "fmt"
    "io"
)

type ConfType struct {
	LOGPATH       string `yaml:"logpath"`
	CSVPATH       string `yaml:"csvpath"`
	SEARCHTIMEOUT int    `yaml:"timeout"`
}

func Parse(file io.Reader) (appConf ConfType, err error) {
	appConf = ConfType{}
	err = yaml.NewDecoder(file).Decode(&appConf)
	return appConf, err
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