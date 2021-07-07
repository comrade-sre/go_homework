package parse
import (
    "testing"
    "math/rand"
    "time"
)

const (
    length = 15
)

var (
     QueryEqual = map[string]string{"us": "us","3242": "3242","23.423": "23.423","China": "China","234": "234.0","1/8/2020": "01/08/2020"}
     QueryBigger = map[string]string{"324242": "432", "55000": "4.0", "43234.423": "4324.43", "1/8/2021": "1/8/2020"}
     letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// thanks to https://stackoverflow.com/users/1705598/icza
func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func TestCompareValues(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    first := RandStringRunes(length)
    second := RandStringRunes(length)
    result := CompareValues(first, second, "=")
    if result == true {
            t.Errorf("gets true, wants false, values: %s = %s", first, second)
        }
    for key, val := range QueryEqual {
        result := CompareValues(key, val, "=")
        if result != true {
            t.Errorf("gets false, wants true, values: %s = %s", key, val)
        }
    }

}
