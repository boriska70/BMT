package monitoring

import (
	"testing"
	"fmt"
	"encoding/json"
	"strconv"
)

func TestMonitorMarshalling(t *testing.T) {
	var foo BmtMon
	foo.bmt_name = "abc"
	foo.Bmt_data = "{{'size':0','aggs':{'authors':{'terms':{'field':'author.keyword'}}}}}"
	foo_marshalled, _ := json.Marshal(foo)
	fmt.Println(string(foo_marshalled))

}

func TestSomething(t *testing.T)  {
	var m map[string]interface{};
	json.Unmarshal([]byte("{\"took\":3341,\"timed_out\":false,\"_shards\":{\"total\":5,\"successful\":5,\"failed\":0},\"hits\":{\"total\":169324328,\"max_score\":0.0,\"hits\":[]},\"aggregations\":{\"count_by_type\":{\"doc_count_error_upper_bound\":0,\"sum_other_doc_count\":0,\"buckets\":[{\"key\":\"consolelog\",\"doc_count\":169175299},{\"key\":\"build\",\"doc_count\":140852},{\"key\":\"pipelines\",\"doc_count\":8176},{\"key\":\"commitfeaturemapping\",\"doc_count\":1}]}}}"), &m);
	h := m["hits"]
	f := h.(map[string]interface{})["total"].(float64)
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 64))


}
