package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"github.com/Sirupsen/logrus"
)

const configFileName = "queries.yml"

type Monitor struct {
	Kind     string
	Name     string
	Url      string
	Body     string
	Interval int
}

func ReadConfigYaml() []Monitor{

	config := [] Monitor{}
	cfgFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		logrus.Fatal("Failed to read " + configFileName)
		os.Exit(1)
	}
	err = yaml.Unmarshal(cfgFile, &config)
	if err != nil {
		logrus.Fatal("Failed to parse " + configFileName)
		os.Exit(1)
	}

	logrus.Infof("Monitoring %d queries...", len(config) )

	return config
}
