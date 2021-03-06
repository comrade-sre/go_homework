// package contains struct and function for work with the yaml configuration
package parse

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/url"
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
// function Parse returns struct with all data for  imaginable service configuration
func Parse(file *os.File) (config, error) {
	appConf := config{}
	err := yaml.NewDecoder(file).Decode(&appConf)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return appConf, err
	} else {
		fmt.Println("Have read configuration file successfully")
	}
	u, err := url.ParseRequestURI(appConf.SentryURL)
	if err != nil {
		fmt.Printf("incorrect url in configuration: %s", u)
		return appConf, err
	}
	u, err = url.ParseRequestURI(appConf.JaegerURL)
	if err != nil {
		fmt.Println("incorrect url in configuration: %s", u)
		return appConf, err
	}
	return appConf, err
}
