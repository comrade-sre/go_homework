package parse

import (
	"gopkg.in/yaml.v2"
	"io"
)

type ConfType struct {
	LOGPATH       string `yaml:"logpath"`
	CSVPATH       string `yaml:"csvpath"`
	SEARCHTIMEOUT int    `yaml:"timeout"`
}

func Parse(file io.Reader) (appConf ConfType, err error) {
	appConf = ConfType{}
	err = yaml.NewDecoder(file).Decode(&appConf)
	return appConf, err
}
