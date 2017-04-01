package elasticsearch

import (
	elastic "gopkg.in/olivere/elastic.v5"
	log "github.com/Sirupsen/logrus"
	"time"
)

var ClientOut *elastic.Client

const CONNECTION_RETRIES  =  5
const CONNECTION_TIMEOUT  = 10

func CreateClient(url string) (*elastic.Client) {
	//https://github.com/olivere/elastic/wiki/Configuration
	var i=0
	for i<CONNECTION_RETRIES {
		log.Infof("Creating Elasticsearch client for %v", url)
		client, err := elastic.NewClient(elastic.SetURL(url))
		if err != nil {
			log.Errorf("Failed to create Elasticsearch client for %s", url)
			i++
			time.Sleep(CONNECTION_TIMEOUT * time.Second)
		} else {
			log.Infof("Elasticsearch client for %s is created", url)
			return client
		}
	}
	return nil
}

func  RunSearch(client *elastic.Client, index string, datatype string, body string) *elastic.SearchResult {
	var searchService = client.Search();
	searchService.Index(index);
	if len(datatype)!=0 {
		searchService.Type(datatype)
	}
	return nil
}