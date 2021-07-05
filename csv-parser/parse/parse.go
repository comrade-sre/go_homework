package parse

import (
	"fmt"
	"github.com/comrade-sre/go_homework/csv-parser/check"
	"gopkg.in/yaml.v2"
	"io"
	"strconv"
	"strings"
)

type ConfType struct {
	LOGPATH       string `yaml:"logpath"`
	LOGERR        string `yaml:"logerrpath"`
	CSVPATH       string `yaml:"csvpath"`
	SEARCHTIMEOUT int    `yaml:"timeout"`
}

func Parse(file io.Reader) (appConf ConfType, err error) {
	appConf = ConfType{}
	err = yaml.NewDecoder(file).Decode(&appConf)
	return appConf, err
}
func CompareValues(first string, second string, op string) (result bool) {
	if check.ValidTime.MatchString(strings.Trim(first, " ")) {
		ftime, err := check.ConvertTime(first)
		if err == nil {
			stime, _ := check.ConvertTime(second)
			switch op {
			case "=":
				result = ftime == stime
				return result
			case ">":
				result = ftime.After(stime)
				return result
			case "<":
				result = ftime.Before(stime)
				return result
			}
		}
	}
	ffloat, err := strconv.ParseFloat(first, 32)
	if err == nil {
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
	} else {
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

	}
	return false
}

func ParseLine(header []string, Query []string, ch <-chan string, Querylength int, FieldPos map[string]int) error {

	for line := range ch {
	    values := strings.Split(line, ",")
		for BeginExpression := 0; BeginExpression <= Querylength-3; BeginExpression += 4 {
			FIELD := Query[BeginExpression]
			OP := Query[BeginExpression+1]
			VALUE := Query[BeginExpression+2]
			res := CompareValues(values[FieldPos[FIELD]], VALUE, OP)
			if res {
				fmt.Println(line)
			}
		}
	}
	return nil
}
