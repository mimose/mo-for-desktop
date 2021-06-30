package space

import (
	"github.com/mimose/gcosy/lib"
	"time"
)

type Space struct {
	Key        string    `json:key`
	Name       string    `json:name`
	CreateTime lib.CTime `json:createTime`
	UpdateTime lib.CTime `json.updateTime`
}

type SpacesList []Space

func (spaces SpacesList) Len() int {
	return len(spaces)
}

func (spaces SpacesList) Less(i, j int) bool {
	return time.Time(spaces[i].CreateTime).Before(time.Time(spaces[j].CreateTime))
}

func (spaces SpacesList) Swap(i, j int) {
	spaces[i], spaces[j] = spaces[j], spaces[i]
}
