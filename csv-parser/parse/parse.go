package parse

import (
    "gopkg.in/yaml.v2"
    "strings"
    "fmt"
    "io"
    "strconv"
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
func CompareValues(first string, second string, op string) (result bool, err error ){

    ffloat, ferr := strconv.ParseFloat(first, 32)
    sfloat, serr := strconv.ParseFloat(second, 32)

    if ferr !=  serr {
        return false, fmt.Errorf("bad query check, %s and %s not comparable", first, second)
        } else if ferr == nil {
            switch op {
        case "=":
            result = ffloat == sfloat
            return result, nil
        case ">":
            result = ffloat > sfloat
            return result, nil
        case "<":
            result = ffloat < sfloat
            return result, nil
            }
        } else {
        switch op {
        case "=":
            result = first == second
            return result, nil
        case ">":
            result = first > second
            return result, nil
        case "<":
            result = first < second
            return result, nil
        }
    }
    return false, fmt.Errorf("we have never should be here!")
}

func ParseLine(header []string, query []string, ch <-chan string,  cherr chan error, Querylength int, FieldPos map[string]int) (err error) {
         BEGIN := 0 //BEGIN < Querylength - 3; BEGIN += 4
        EXPRESSION := query[BEGIN:BEGIN+3]
        FIELD := EXPRESSION[0]
        OP := EXPRESSION[1]
        VALUE := EXPRESSION[3]
        for line := range ch {
            values := strings.Split(line, ",")
            if res, err := CompareValues(values[FieldPos[FIELD]], VALUE, OP); err != nil {
                return err
                } else if res {
                fmt.Println(line)
                }
            }
        return nil


}
