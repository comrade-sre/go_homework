package check

import (
	"fmt"
	"strings"
	"strconv"
)

func CheckQueryTypes(IsString map[string]bool, query []string, Querylength int) (err error) {
	for i := 0; i < Querylength; i += 4 {
	    IsHeaderString := IsString[query[i]]
	    value := i+2
	    _, err := strconv.ParseFloat(query[value], 32)
	    if err != nil && !IsHeaderString {
	        return fmt.Errorf("%s field has a type other than %s, not comparable, beacuse %s", query[i], query[value], IsHeaderString )
	    } else if err == nil && IsHeaderString {
	        return fmt.Errorf("%s field has a type other than %s, not comparable",query[i], query[value] )
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
func CheckQuerySyntax(header []string, query []string, Querylength int) (err error) {
	HeaderLength := len(header)
	LogOpMax := HeaderLength - 1
	QearyMax := (HeaderLength * 3) + LogOpMax

	if Querylength > QearyMax {
		return fmt.Errorf("query is too long, max length: %d", QearyMax)

	}
	for i := 0; i < Querylength; i += 4 {
		if _, err = CheckHeader(header, query[i]); err != nil {
			return err
		}
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
