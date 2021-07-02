package parse

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"strconv"
	"strings"

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
func CompareValues(first string, second string, op string) (result bool) {

	ffloat, ferr := strconv.ParseFloat(first, 32)
	if ferr != nil {
		switch op {
		case "=":
			result = first == second
			return result
		case ">":
			result = first > second
			return result
		case "<":
			result = first < second
			return result
		}

	} else {
		sfloat, _ := strconv.ParseFloat(second, 32)
		switch op {
		case "=":
			result = ffloat == sfloat
			return result
		case ">":
			result = ffloat > sfloat
			return result
		case "<":
			result = ffloat < sfloat
			return result
		}

	}
	return false
}

func ParseLine(header []string, query []string, ch <-chan string, Querylength int, FieldPos map[string]int) {
	//BEGIN := 0 //BEGIN < Querylength - 3; BEGIN += 4
	EXPRESSION := query //[BEGIN:BEGIN+2]
	FIELD := EXPRESSION[0]
	OP := EXPRESSION[1]
	VALUE := EXPRESSION[2]
	for line := range ch {
		values := strings.Split(line, ",")
		res := CompareValues(values[FieldPos[FIELD]], VALUE, OP)
		if res {
			fmt.Println(line)
		}
	}
}
