package parse

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type config struct {
	port         int    `yaml:port`
	db_url       string `yaml:db_url`
	jaeger_url   string `yaml:jaeger_url`
	sentry_url   string `yaml:sentry_url`
	kafka_broker string `yaml:kafka_broker`
	some_app_id  string `yaml:some_app_id`
	some_app_key string `yaml:some_app_key`
}

func Parse(filename string) (config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %v\n", err)
	}
	appConf := config{}
	err = yaml.Unmarshal(yamlFile, &appConf)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
	return appConf, err
}
