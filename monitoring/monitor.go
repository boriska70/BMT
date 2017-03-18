package monitoring

import (
	"time"
	"fmt"
	_"math/rand"
	"github.com/boriska70/bmt/config"
	"github.com/boriska70/bmt/util"
	clients "github.com/boriska70/bmt/clients"
	"github.com/Sirupsen/logrus"
	"net/http"
	"strings"
	"bytes"
	"io/ioutil"
)

var inputSource = [] rune("abcdefghijklmnopqrstuvwxyz")
var inputLength = 3

func FetchData(ch chan string, monitor util.Monitor) {
	fmt.Printf("My monitor is %s\n", monitor)
	for true {

		if strings.EqualFold(monitor.Kind, "http") {
			url := config.IES+"/"+monitor.Index
			if len(monitor.Type)>0 {
				url += url+"/"+monitor.Type
			}
			res := fetchDataOverHttp(monitor, url)
			ch <- res

		}

/*		outputStart := rand.Intn(len(inputSource) - inputLength)
		logrus.Infof("Sending data for monitor %s (%s)", monitor.Name, monitor.Index)
		ch <- string(inputSource[outputStart:outputStart+inputLength])*/

		time.Sleep(time.Duration(monitor.Interval) * time.Second)

	}
}

func fetchDataOverHttp(monitor util.Monitor, url string) string {
	client := clients.CreateHttpClient();
	req, errReq := http.NewRequest(monitor.Method, url, bytes.NewBuffer([]byte(monitor.Body)))
	if errReq != nil {
		logrus.Error("Cannot create http req for monitor %s", monitor.Name)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	resp, errResp := client.Do(req)
	if errResp != nil {
		logrus.Error("Error occurred when fetching data for monitor %s. Error message: %s", monitor.Name, errResp.Error())
	}
	defer resp.Body.Close()

	logrus.Info("Status for fetching data for monitor %s is %d", monitor.Name, resp.StatusCode)
	logrus.Info("Response headers are %v", resp.Header)
	body,_ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func SendData(ch chan string) {
	for true {
		var data = <-ch
		fmt.Printf("Data received: %s\n", data)

	}
}
