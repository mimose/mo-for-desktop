package record

import (
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/model/errs"
	"time"
)

type Record struct {
	Key        string    `json:key`
	Title      string    `json:title`
	SpaceKey   string    `json:spaceKey`
	RecordType int       `json:recordType`
	Content    string    `json:content`
	NoticeTime lib.CTime `json:noticeTime`
	CoverPic   string    `json:coverPic`
	CreateTime lib.CTime `json:createTime`
	UpdateTime lib.CTime `json.updateTime`
}

func (r Record) Validate() error {
	if r.RecordType != Notice && r.RecordType != Note {
		return lib.NewError(errs.AddRecordValidateError, "未知的类型", nil)
	}
	if r.Title == "" {
		return lib.NewError(errs.AddRecordValidateError, "主题不可为空", nil)
	}
	if r.SpaceKey == "" {
		return lib.NewError(errs.AddRecordValidateError, "分组不可为空", nil)
	}
	return nil
}

type RecordsList []Record

type RecordsGroup struct {
	NoticesList RecordsList `json:notices`
	NotesList   RecordsList `json:notes`
}

// ======= 排序interface
func (records RecordsList) Len() int {
	return len(records)
}

func (records RecordsList) Less(i, j int) bool {
	return time.Time(records[i].CreateTime).Before(time.Time(records[j].CreateTime))
}

func (records RecordsList) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}

// ======= 排序interface

func (records RecordsList) FilterRecordListByRecordType(recordType int) RecordsList {
	index := 0
	for _, record := range records {
		if record.RecordType == recordType {
			records[index] = record
			index++
		}
	}
	return records[:index]
}

func (records RecordsList) Group() *RecordsGroup {
	var notices, notes RecordsList
	if records != nil && records.Len() > 0 {
		notices = make(RecordsList, 0, records.Len()/2)
		notes = make(RecordsList, 0, records.Len()/2)
		for _, record := range records {
			switch record.RecordType {
			case Notice:
				notices = append(notices, record)
			case Note:
				notes = append(notes, record)
			default:
				continue
			}
		}
	}
	return &RecordsGroup{
		NoticesList: notices,
		NotesList:   notes,
	}
}
