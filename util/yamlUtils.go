package util

import (
	"gopkg.in/yaml.v2"
//	"io/ioutil"
	"os"
	"github.com/Sirupsen/logrus"
)

const configFileName = "queries.yml"

type Monitor struct {
	Kind     string
	Name     string
	Index      string
	Type      string
	Body     string
	Interval int
}

func ReadConfigYaml(cfgBytes []byte) []Monitor{

	config := [] Monitor{}
//	cfgFile, err := ioutil.ReadFile("../"+configFileName)
/*	cfgFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		logrus.Fatal("Failed to read " + configFileName)
		os.Exit(1)
	}*/
	err := yaml.Unmarshal(cfgBytes, &config)
	if err != nil {
		logrus.Fatal("Failed to parse " + configFileName)
		os.Exit(1)
	}

	logrus.Infof("Monitoring %d queries...", len(config) )

	return config
}
