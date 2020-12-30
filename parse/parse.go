package parse

import (
	"gopkg.in/yaml.v2"
	"log"
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

func Parse(file *os.File) (config, error) {
	appConf := config{}
	err := yaml.NewDecoder(file).Decode(&appConf)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	} else {
		log.Println("Have read configuration file successfully")
	}
	return appConf, err
}
