package check

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	ValidTime = regexp.MustCompile(`^[0-9]{1,2}\/[0-9]{1,2}\/[0-9]{4}$`)
)

func ConvertTime(value string) (time.Time, error) {
	shortFormat := "01/02/2006"
	var newValue strings.Builder
	newValue.Grow(32)
	splitValue := strings.Split(value, "/")
	if utf8.RuneCountInString(splitValue[0]) == 1 {
		splitValue[0] = "0" + splitValue[0]
	}
	fmt.Fprintf(&newValue, "%s/", splitValue[0])
	if utf8.RuneCountInString(splitValue[1]) == 1 {
		splitValue[1] = "0" + splitValue[1]
	}
	fmt.Fprintf(&newValue, "%s/", splitValue[1])
	fmt.Fprintf(&newValue, "%s", splitValue[2])
	date, err := time.Parse(shortFormat, newValue.String())
	return date, err

}

func CheckQueryTypes(IsString map[string]bool, query []string, Querylength int) (err error) {
	for i := 0; i < Querylength; i += 4 {
		IsHeaderString := IsString[query[i]]
		value := i + 2
		_, err := strconv.ParseFloat(query[value], 32)
		if err != nil && !IsHeaderString {
			return fmt.Errorf("%s field has a type other than %s, not comparable", query[i], query[value])
		} else if err == nil && IsHeaderString {
			return fmt.Errorf("%s field has a type other than %s, not comparable", query[i], query[value])
		} else if ValidTime.MatchString(strings.Trim(query[value], " ")) {
			_, err := ConvertTime(query[value])
			if err != nil {
				return err
			}
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
	return -1, fmt.Errorf("such column not found: %s", column)

}
func CheckQuerySyntax(header []string, query []string, Querylength int, FieldPos map[string]int) (err error) {
	HeaderLength := len(header)
	LogOpMax := HeaderLength - 1
	QearyMax := (HeaderLength * 3) + LogOpMax

	if Querylength > QearyMax {
		return fmt.Errorf("query is too long, max length: %d", QearyMax)

	}
	for i := 0; i < Querylength; i += 4 {
		index, err := CheckHeader(header, query[i])
		if err != nil {
			return err
		}
		FieldPos[query[i]] = index
	}
	for i := 1; i < Querylength; i += 4 {
		op := query[i]
		if op != ">" && op != "<" && op != "=" {

			return fmt.Errorf("%s is inappropriate operator(it should be ><=)", op)
		}
		for i := 3; i < Querylength; i += 4 {
			op := strings.ToLower(query[i])
			if op != "and" && op != "or" {
				return fmt.Errorf("%s is inappropriate logical operator(it should be and/or)", op)
			}
		}
	}
	return nil
}
