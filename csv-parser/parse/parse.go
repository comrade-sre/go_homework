package parse

import (
	"gopkg.in/yaml.v2"
	"io"
	"fmt"
	"strings"
	//"errors"
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
func CheckHeader(header []string, column string) (index int, err error) {
	for index, value := range header {
		if value == column {
			return index, nil
		}
	}
	return 0, fmt.Errorf("such column not found: %s", column)

}
func ParseQuerySyntax(header []string, query []string) (err error) {
    HeaderLength := len(header)
    LogOpMax := HeaderLength - 1
    QearyMax := (HeaderLength * 3) + LogOpMax
    Querylength := len(query)
    if Querylength > QearyMax {
        return fmt.Errorf("query is too long, max length: %d", QearyMax)

    }
    for i := 0; i < Querylength; i+=4 {
        if _, err = CheckHeader(header, query[i]); err != nil {
        return err
        }
    }
    for i := 1; i < Querylength; i+=4 {
        op := query[i]
        if  op != ">" && op != "<" && op != "=" {

            return fmt.Errorf("%s is inappropriate operator(it should be ><=)", op)
        }
    for i := 3; i < Querylength; i+=4 {
        op := strings.ToLower(query[i])
        if op != "and" && op != "or" {
        return fmt.Errorf("%s is inappropriate logical operator(it should be and/or)", op)
        }
    }
    }



    return nil

}