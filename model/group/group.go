package group

import (
	"github.com/mimose/gcosy/lib"
	"time"
)

type Group struct {
	Key        string    `json:key`
	Name       string    `json:name`
	CreateTime lib.CTime `json:createTime`
	UpdateTime lib.CTime `json.updateTime`
}

type GroupsList []Group

func (groups GroupsList) Len() int {
	return len(groups)
}

func (groups GroupsList) Less(i, j int) bool {
	return time.Time(groups[i].CreateTime).Before(time.Time(groups[j].CreateTime))
}

func (groups GroupsList) Swap(i, j int) {
	groups[i], groups[j] = groups[j], groups[i]
}
