package elasticsearch

import (
	elastic "gopkg.in/olivere/elastic.v5"
	log "github.com/Sirupsen/logrus"
)

var ClientIn *elastic.Client
var ClientOut *elastic.Client

func CreateClient(url string) (*elastic.Client) {
	//https://github.com/olivere/elastic/wiki/Configuration
	if true {
		log.Infof("Creating Elasticsearch client for %v", url)
		client, err := elastic.NewClient(elastic.SetURL(url))
		if err != nil {
			log.Infof("Failed to create Elasticsearch client for %s", url)
		} else {
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