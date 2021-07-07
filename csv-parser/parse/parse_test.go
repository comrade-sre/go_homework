package parse
import (
    "testing"

)

var (
     QueryEqual = map[string]string{"us": "us","3242": "3242","23.423": "23.423","China": "China","234": "234.0","1/8/2020": "01/08/2020"}

)

func TestCompareValues(t *testing.T) {
    for key, val := range QueryEqual {
        result := CompareValues(key, val, "=")
        if result != true {
            t.Errorf("gets false, wants true, values: %s = %s", key, val)
        }
    }

}
