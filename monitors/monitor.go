package monitors

import (
	"time"
	"fmt"
	"math/rand"
)

var inputSource = [] rune("abcdefghijklmnopqrstuvwxyz")
var inputLength = 3

func FetchData(ch chan string)  {
	for true {
		time.Sleep( 1 * time.Second)
		outputStart := rand.Intn(len(inputSource)-inputLength)
		ch <- string(inputSource[outputStart:outputStart+inputLength])
	}
}

func SendData(ch chan string)  {
	for true {
		var data = <- ch
		fmt.Printf("Data received: %s\n", data)
	}
}