package record

import (
	"sort"
	"sync"
)

var (
	recordLock = sync.Mutex{}
	records    = []Record{}
)

type Record struct {
	ID           string
	Timestamp    int64
	Result       string
	RemainingQty int64
}

func InsertRecord(record Record) {
	recordLock.Lock()
	defer recordLock.Unlock()
	records = append(records, record)
}

func GetRecords() []Record {
	recordLock.Lock()
	defer recordLock.Unlock()
	sort.Slice(records, func(i, j int) bool {
		if records[i].RemainingQty != records[j].RemainingQty {
			return records[i].RemainingQty > records[j].RemainingQty
		}
		return records[i].Timestamp < records[j].Timestamp
	})
	return records
}
