package parse

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type config struct {
	Port        int    `yaml:"port"`
	DBURL       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broke"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func Parse(filename string) (config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
//	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %v\n", err)
	}
	appConf := config{}
	//	err = yaml.NewDecoder(file).Decode(&appConf)
	err = yaml.Unmarshal(yamlFile, &appConf)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
	return appConf, err
}
