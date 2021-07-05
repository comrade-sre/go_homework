package main
import (
	"time"
	"fmt"
)

func main () {
	
	const shortForm = "2006-01-02"
	_, err := time.Parse(shortForm, "2001-09-11")
	fmt.Println(err)

}

