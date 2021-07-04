package main
import (
	"time"
	"fmt"
)

func main () {
	//loc, _ := time.LoadLocation("Europe/Moscow")
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09")
	fmt.Println(t)	
	

}

