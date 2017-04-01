package monitoring

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestMonitorMarshalling(t *testing.T) {
	var foo BmtMon
	foo.bmt_name = "abc"
	foo.Bmt_data = "{{'size':0','aggs':{'authors':{'terms':{'field':'author.keyword'}}}}}"
	foo_marshalled, _ := json.Marshal(foo)
	fmt.Println(string(foo_marshalled))

}

