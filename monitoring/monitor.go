package monitoring

import (
	"time"
	"fmt"
	_"math/rand"
	"github.com/boriska70/bmt/config"
	"github.com/boriska70/bmt/util"
	clients "github.com/boriska70/bmt/clients"
	es "github.com/boriska70/bmt/elasticsearch"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
	"bytes"
	"io/ioutil"
	json "encoding/json"
	"context"
)

func FetchData(ch chan BmtMon, monitor util.Monitor) {
	fmt.Printf("My monitor is %s\n", monitor)
	for true {
		if strings.EqualFold(monitor.Kind, "http") {
			url := config.IES + "/" + monitor.Index
			if len(monitor.Type) > 0 {
				url += "/" + monitor.Type
			}
			url += "/_search"
			res := fetchDataOverHttp(monitor, url)

			//https://eager.io/blog/go-and-json/
			var parsed map[string]interface{}
			err := json.Unmarshal([]byte(res), &parsed)
			if err != nil {
				log.Errorln("Error while parshing response: " + res)
			} else {
				hits := parsed["hits"]
				if hits != nil {
					f := hits.(map[string]interface{})["total"].(float64)
					totalMap := make(map[string]interface{})
					totalMap["ts"] = time.Now().UTC().Round(time.Second)
					totalMap["total_count"] = f
					totalMap["metric_name"] = monitor.Name
					var totalBytes []byte;
					totalBytes,_ = json.Marshal(totalMap)
					var data BmtMon
					data.bmt_name = monitor.Name + "_total"
					data.Bmt_data = string(totalBytes)
					ch <- data

					hitsMap, _ := hits.(map[string]interface{})
					if len(hitsMap["hits"].([]interface{})) > 0 {
						var hitsBytes []byte
						hitsMap["ts"] = time.Now().UTC().Round(time.Second)
						hitsMap["metric_name"] = monitor.Name
						hitsBytes, _ = json.Marshal(hitsMap)
						//log.Println("Sending hits: " + string(hitsBytes))
						var data BmtMon
						data.bmt_name = monitor.Name + "_hits"
						data.Bmt_data = string(hitsBytes)
						ch <- data
					}
				} else {
					log.Warn("No hits in the response: %s", res)
				}
				aggs := parsed["aggregations"]
				if aggs != nil {
					aggsMap, _ := aggs.(map[string]interface{})
					if len(aggsMap) > 0 {
						var aggsBytes []byte
						aggsMap["ts"] = time.Now().UTC().Round(time.Second)
						aggsMap["metric_name"] = monitor.Name
						aggsBytes, _ = json.Marshal(aggsMap)
						//log.Println("Sending aggregations: " + string(aggsBytes))
						var data BmtMon
						data.bmt_name = monitor.Name + "_aggregations"
						data.Bmt_data = string(aggsBytes)
						ch <- data
					}
				} else {
					log.Warn("No aggregations in the response: %s", res)
				}

			}


		}

		time.Sleep(time.Duration(monitor.Interval) * time.Second)

	}
}

func fetchDataOverHttp(monitor util.Monitor, url string) string {
	client := clients.CreateHttpClient();
	req, errReq := http.NewRequest(monitor.Method, url, bytes.NewBuffer([]byte(monitor.Body)))
	if errReq != nil {
		log.Error("Cannot create http req for monitor %s", monitor.Name)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	resp, errResp := client.Do(req)
	if errResp != nil {
		log.Error("Error occurred when fetching data for monitor %s. Error message: %s", monitor.Name, errResp.Error())
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()




	log.Infof("Status for fetching data for monitor %s is %d", monitor.Name, resp.StatusCode)
	log.Infof("Response headers are %v", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func SendData(ch chan BmtMon) {
	for true {
		var data = <-ch
		//dataStr,_ := json.Marshal(data);
		fmt.Printf("Data received: %v\n", data)
		_, err := es.ClientOut.Index().Index("bmt").Type(data.bmt_name).BodyJson(data.Bmt_data).Do(context.Background())
		if err != nil {
			fmt.Println("Error: %s", err)
		}
	}
}
