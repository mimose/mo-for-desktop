package group

import (
	"fmt"
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/services/storage"
	"time"
)

// local storage
// ../group/group_[encode(key)]  ---- value: encode(Group)
// ../group/recordRel/group_[encode(key) --- value: encode(List: record_local_storage_key)
type Group struct {
	Key        string    `json:key`
	Name       string    `json:name`
	CreateTime lib.CTime `json:createTime`
	UpdateTime lib.CTime `json.updateTime`
}

func ListAll() []Group {
	return nil
}

func AddOne(name string) error {
	now := time.Now()
	key := storage.RandStringBytesMaskImprSrcUnsafe()
	group := Group{
		Key:        key,
		Name:       name,
		CreateTime: lib.CTime(now),
		UpdateTime: lib.CTime(now),
	}
	data, err := lib.Encode(group)
	if err != nil {
		fmt.Printf("[error] Group AddOne. %s", err)
		return err
	}
	dirPath, fileName := storage.LocalGroupDirByKey(key)
	err = lib.Write(dirPath, fileName, data, false)
	if err != nil {
		fmt.Printf("[error] Group AdddOne. %s", err)
	}
	return nil
}
