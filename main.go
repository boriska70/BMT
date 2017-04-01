package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/bmt/config"
	"github.com/boriska70/bmt/web"
	"github.com/boriska70/bmt/util"
	"github.com/boriska70/bmt/monitoring"
	client "github.com/boriska70/bmt/elasticsearch"
	"io/ioutil"
	"os"
)

func main() {

	v1 := flag.String("ies", "http://localhost:9200", "Elasticsearch URL to query for data")
	v2 := flag.String("oes", "http://localhost:9200", "Elasticsearch URL to save data")
	v3 := flag.String("queries", "queries.yml", "Path to queries.yml")
	flag.Parse()

	config.IES = *v1
	config.OES = *v2
	config.CfgFile = *v3

	log.Info("Hello")
	log.Infof("Parameters received: input ES - %s, output ES - %s, configuration file path - %s", config.IES, config.OES, config.CfgFile)
	cfgFile, cfgErr := ioutil.ReadFile(config.CfgFile)
	if cfgErr != nil {
		log.Fatalf("Cannot find queries.yml in %v", config.CfgFile)
	}
	monitors := util.ReadConfigYaml(cfgFile)

	client.ClientOut = client.CreateClient(config.OES)

	if client.ClientOut == nil {
		os.Exit(-13)
	} else {
		stateOut := client.ClientOut.ClusterState();
		log.Infof("Cluster out state: %", stateOut);
		dataChannel := make(chan monitoring.BmtMon, 100)

		go monitoring.SendData(dataChannel)
		for _, monitor := range monitors {
			if (monitor.Disable) {
				log.Infof("Monitor %s is disabled in configuration, will not run", monitor.Disable)
			} else {
				go monitoring.FetchData(dataChannel, monitor)
			}
		}

		//time.Sleep(15 * time.Second)

		web.StartHttpServer()
	}
}
