package group

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mimose/gcosy/lib"
	"io/ioutil"
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
	groups := []Group{}
	groupDir := storage.LocalGroupDir()
	bytes, err := FilesBytes(groupDir)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Group ListAll. %s", err)
		return groups
	}
	if bytes == nil {
		// TODO log
		fmt.Printf("[error] Group ListAll. byte empty")
		return groups
	}
	for _, byte := range bytes {
		var group Group
		json.Unmarshal(byte, &group)
		groups = append(groups, group)
	}
	return groups
}

func FilesBytes(dirPath string) ([][]byte, error) {
	if !lib.DirExists(dirPath) {
		return nil, errors.New("dir is not exists")
	}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	//var filesContents []string
	var fileBytes [][]byte
	for _, file := range files {
		byte, err := lib.Read(lib.CompletePath(dirPath, file.Name()))
		if err != nil {
			return nil, err
		}
		fileBytes = append(fileBytes, byte)
	}
	return fileBytes, nil
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
