package parse

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type config struct {
	port         int
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string
}

func Parse(f *os.File) {
	appConf := config{}
	buf := bytes.NewBuffer(make([]byte, 2000))
	data, err := buf.ReadFrom(f)
	if err != nil {
		Fprintf(os.Stderror, "there is an error occured: %v\n", err)
	}
	err = yaml.Unmarshal([]byte(data), &appConf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", appConf)
	data, err := yaml.Marshal(&appConf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- config data dump:\n%s\n\n", string(data))

}
