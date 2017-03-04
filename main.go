package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/bmt/web"
	"github.com/boriska70/bmt/util"
	"github.com/boriska70/bmt/monitoring"
)

func main() {

	flag.String("ies", "http://localhost:9200", "Elasticsearch URL to query for data")
	flag.String("oes", "http://localhost:9200", "Elasticsearch URL to save data")

	fmt.Printf("Hello\n")
	monitors := util.ReadConfigYaml()

	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	state := client.ClusterState();
	log.Infof("Cluster state: %", state);
	dataChannel := make(chan string, 100)

	go monitoring.SendData(dataChannel)
	for _, monitor := range monitors {
		go monitoring.FetchData(dataChannel, monitor)
	}

	//time.Sleep(15 * time.Second)

	web.StartHttpServer()

}
