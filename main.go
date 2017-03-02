package main

import (
	"fmt"
	"io/ioutil"
	elastic "gopkg.in/olivere/elastic.v5"
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/BMT/monitors"
	"github.com/boriska70/BMT/web"
)

func main() {

	flag.String("ies", "http://localhost:9200", "Elasticsearch URL to query for data")
	flag.String("oes", "http://localhost:9200", "Elasticsearch URL to save data")

	fmt.Printf("Hello\n")
	qBytes, _ := ioutil.ReadFile("./queries.yml")
	fmt.Printf("File: %s", string(qBytes[:len(qBytes)]))

	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
	}

	state := client.ClusterState();
	log.Infof("Cluster state: %", state);
	dataChannel := make(chan string, 100)

	go monitors.FetchData(dataChannel)
	go monitors.SendData(dataChannel)

	//time.Sleep(15 * time.Second)

	web.StartHttpServer()

}
