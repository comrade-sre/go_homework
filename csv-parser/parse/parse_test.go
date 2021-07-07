package parse
import (
    "testing"
    "fmt"

)

var (
     QueryEqual = []string{"us","uS","3242","3242","23.423","23.423","China","cHina","234", "234.0","1/8/2020","01/08/2020"}

)

func TestCompareValues(t *testing.T) {
    for i := 0; i == len(QueryEqual)-3;i +=2 {
        fmt.Println(QueryEqual[i], QueryEqual[i+1])
        result := CompareValues(QueryEqual[i], QueryEqual[i+1], "=")
        if result != true {
            t.Errorf("gets false, wants true, values: %s = %s", QueryEqual[i], QueryEqual[i+1])
        }
    }

}
