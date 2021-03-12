package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

var (
	buf bytes.Buffer
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "using: script number")
	}

	fmt.Println(comma(os.Args[1]))

	fmt.Println(bufComma(os.Args[1]))
}
func bufComma(s string) string {

	for i, n := range s {
		tmpBuf := make([]byte, 1)
		_ = utf8.EncodeRune(tmpBuf, n)
		val, _ := strconv.Atoi(string(tmpBuf))
		if i%3 == 0 && i > 0 {
			fmt.Fprintf(&buf, ",")
			fmt.Fprintf(&buf, "%d", val)

		} else {
			fmt.Fprintf(&buf, "%d", val)
		}
	}
	return buf.String()
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
