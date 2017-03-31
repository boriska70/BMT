package monitoring

import (
	//"time"
)

type BmtMon struct {
	bmt_name string
	Bmt_data interface{} `json:"data"`
}

/*type BmtMonEnriched struct {
	Bmt_ts  time.Time `json:"ts"`
	Bmt_mon BmtMon `json:"result"`
}*/
