package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigYaml(t *testing.T)  {

	config := ReadConfigYaml()
	assert.NotNil(t, config)
	assert.Equal(t, 2, len(config))

}

func createMonitors() []Monitor {
	m := [] Monitor{}
	var m1 Monitor
	m1.Url="sad"
	m1.Name="asd"
	m1.Kind="es"
	m1.Interval=60
	m1.Body="{}"
	m = append(m, m1)
	return m
}