package record

import (
	"github.com/mimose/gcosy/lib"
	"time"
)

type Record struct {
	Key        string    `json:key`
	Title      string    `json:title`
	SpaceKey   string    `json:spaceKey`
	RecordType int       `json:recordType`
	Content    string    `json:content`
	CoverPic   string    `json:coverPic`
	CreateTime lib.CTime `json:createTime`
	UpdateTime lib.CTime `json.updateTime`
}

type RecordsList []Record

func (records RecordsList) Len() int {
	return len(records)
}

func (records RecordsList) Less(i, j int) bool {
	return time.Time(records[i].CreateTime).Before(time.Time(records[j].CreateTime))
}

func (records RecordsList) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}
