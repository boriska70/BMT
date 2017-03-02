package main

import (
	log "github.com/Sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v5"
)

func queryDataAndReport(cli elastic.Client, ch chan string) {

	for {

		data := <-ch
		log.Infof("Going to send container data to ES: %s", data)
	}
}
