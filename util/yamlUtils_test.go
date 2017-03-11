package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"io/ioutil"
	"os"
)

func TestReadConfigYaml(t *testing.T)  {

	cfgBytes, err := ioutil.ReadFile("queries.yml")
	if err!=nil{
		fmt.Errorf("Problem occured: %v", err)
		os.Exit(-1)
	}
	config := ReadConfigYaml(cfgBytes)
	assert.NotNil(t, config)
	assert.Equal(t, 2, len(config))
	fmt.Printf("Monitor: %b", config[1].Index)

}

func createMonitors() []Monitor {
	m := [] Monitor{}
	var m1 Monitor
	m1.Kind="es"
	m1.Name="asd"
	m1.Index="tweet"
	m1.Body="{}"
	m1.Interval=60
	m = append(m, m1)
	return m
}