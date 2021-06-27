package record

import (
	"github.com/mimose/gcosy/lib"
)

// local storage
// ../record/record_[encode(key)]  ---- value: encode(Record)
type Record struct {
	Key        string    `json:key`
	Title      string    `json:title`
	RecordType int       `json:recordType`
	Content    string    `json:content`
	CoverPic   string    `json:coverPic`
	CreateTime lib.CTime `json:createTime`
	UpdateTime lib.CTime `json.updateTime`
}

func Lists(groupKey, recordType int) []Record {
	return nil
}
