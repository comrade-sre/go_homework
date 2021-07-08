package parse

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	TestLength = 100
)

var (
	QueryEqual  = map[string]string{"us": "us", "3242": "3242", "23.423": "23.423", "China": "China", "234": "234.0", "1/8/2020": "01/08/2020"}
	QueryBigger = map[string]string{"324242": "432", "55000": "4.0", "43234.423": "4324.43", "1/8/2021": "1/8/2020", "abcd": "ABCD"}
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	wg          = sync.WaitGroup{}
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
	for step := 0; step <= TestLength; step++ {
		wg.Add(1)
		go func() {
			b := make([]rune, rand.Intn(100))
			a := make([]rune, rand.Intn(100))
			c := fmt.Sprint(rand.Intn(10000))
			d := fmt.Sprint(rand.Intn(10000))
			for index := range b {
				b[index] = letterRunes[rand.Intn(len(letterRunes))]
			}
			for index := range a {
				a[index] = letterRunes[rand.Intn(len(letterRunes))]
			}

			resString := CompareValues(string(a), string(b), "=")
			t.Logf("testing string values: %s = %s", string(a), string(b))
			resInt := CompareValues(c, d, "=")
			t.Logf("testing string values: %s = %s", c, d)
			if resString == true {
				t.Errorf("gets true, wants false, values: %s = %s", string(a), string(b))
			}
			if resInt == true {
				t.Errorf("gets true, wants false, values: %s = %s", c, d)
			}
			wg.Done()
		}()

	}
	wg.Wait()

	for key, val := range QueryEqual {
		result := CompareValues(key, val, "=")
		if result != true {
			t.Errorf("gets false, wants true, values: %s = %s", key, val)
		}
	}
	for key, val := range QueryBigger {
		result := CompareValues(key, val, ">")
		if result != true {
			t.Errorf("gets false, wants true, values: %s > %s", key, val)
		}
	}

}
