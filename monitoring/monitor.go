package monitoring

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/boriska70/bmt/util"
	//client "github.com/boriska70/bmt/elasticsearch"
	"github.com/Sirupsen/logrus"
)

var inputSource = [] rune("abcdefghijklmnopqrstuvwxyz")
var inputLength = 3

func FetchData(ch chan string, monitor util.Monitor)  {
	fmt.Printf("My monitor is %s\n", monitor)
	for true {
		outputStart := rand.Intn(len(inputSource)-inputLength)
		logrus.Infof("Sending data for monitor %s (%s)", monitor.Name, monitor.Index)
		ch <- string(inputSource[outputStart:outputStart+inputLength])
		time.Sleep( time.Duration(monitor.Interval) * time.Second)

//		client.ClientIn.RunSearch(monitor)

	}
}

func SendData(ch chan string)  {
	for true {
		var data = <- ch
		fmt.Printf("Data received: %s\n", data)


	}
}
