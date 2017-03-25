package monitoring

import (
	"time"
	"encoding/json"
)

type BmtMon struct {
	bmt_ts time.Time
	bmt_name string
	bmt_data string
}

func (m BmtMon) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ts time.Time
		data string
	}{
		m.bmt_ts,
		m.bmt_data,
	})
}

