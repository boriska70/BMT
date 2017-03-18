package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"io/ioutil"
	"os"
)

func TestReadConfigYaml(t *testing.T)  {

	cfgBytes, err := ioutil.ReadFile("queries_test.yml")
	if err!=nil{
		fmt.Errorf("Problem occured: %v", err)
		os.Exit(-1)
	}
	config := ReadConfigYaml(cfgBytes)
	assert.NotNil(t, config)
	assert.Equal(t, 2, len(config))
	assert.False(t, config[0].Disable)
	assert.True(t, config[1].Disable)

	assert.Equal(t, config[0], createMonitors()[0])


	fmt.Printf("Monitor: %b", config[0].Index)

}

func createMonitors() []Monitor {
	m := [] Monitor{}
	var m1 Monitor
	m1.Kind="http"
	m1.Name="mon1"
	m1.Index="tweets"
	m1.Method="POST"
	m1.Body="{{'size':0','aggs':{'authors':{'terms':{'field':'author.keyword'}}}}}"
	m1.Interval=20
	m1.Disable=false
	m = append(m, m1)
	return m
}