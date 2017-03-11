package main

import (
	//elastic "gopkg.in/olivere/elastic.v5"
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/bmt/web"
	"github.com/boriska70/bmt/util"
	"github.com/boriska70/bmt/monitoring"
	client "github.com/boriska70/bmt/elasticsearch"
	"io/ioutil"
)

func main() {

	ies := flag.String("ies", "http://localhost:9200", "Elasticsearch URL to query for data")
	oes := flag.String("oes", "http://localhost:9200", "Elasticsearch URL to save data")
	cfg := flag.String("queries","queries.yml","Path to queries.yml")

	log.Info("Hello\n")
	cfgFile,cfgErr := ioutil.ReadFile(*cfg)
	if cfgErr!=nil {
		log.Fatalf("Cannot find queries.yml in %v", cfg)
	}
	monitors := util.ReadConfigYaml(cfgFile)

	client.ClientIn= client.CreateClient(*ies)
	client.ClientOut= client.CreateClient(*oes)

	stateIn := client.ClientIn.ClusterState();
	stateOut := client.ClientOut.ClusterState();
	log.Infof("Cluster in state: %", stateIn);
	log.Infof("Cluster out state: %", stateOut);

	dataChannel := make(chan string, 100)

	go monitoring.SendData(dataChannel)
	for _, monitor := range monitors {
		go monitoring.FetchData(dataChannel, monitor)
	}

	//time.Sleep(15 * time.Second)

	web.StartHttpServer()

}
