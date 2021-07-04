package space_record_rel

import (
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/services/storage"
	"strings"
	"sync"
)

// local storage
// ../space/recordRel/space_key --- value: Encrypted(List: record_local_storage_key)
var relMutex sync.RWMutex

func AppendRel(spaceKey, recordKey string) error {
	relMutex.Lock()
	defer relMutex.Unlock()

	spaceRelFilePath := lib.CompletePath(storage.LocalSpaceRecordRelDir(), spaceKey)
	relExists := lib.FileExists(spaceRelFilePath)
	// 若存在，则追加，否则新建
	_, err := lib.Write(storage.LocalSpaceRecordRelDir(), spaceKey, []byte(recordKey+"\n"), relExists)
	return err
}

func GetRel(spaceKey string) ([]string, error) {
	relMutex.RLock()
	defer relMutex.RUnlock()

	spaceRelFilePath := lib.CompletePath(storage.LocalSpaceRecordRelDir(), spaceKey)
	byte, err := lib.Read(spaceRelFilePath)
	if err != nil {
		return nil, err
	}
	recordKeyLineStr := string(byte)
	recordKeys := strings.Split(recordKeyLineStr, "\n")
	if len(recordKeys) > 0 {
		index := 0
		for _, recordKey := range recordKeys {
			recordKey = strings.Replace(recordKey, " ", "", -1)
			if recordKey != "" {
				recordKeys[index] = recordKey
				index++
			}
		}
		recordKeys = recordKeys[:index]
	}
	return recordKeys, nil
}

func RemoveRel(spaceKey, recordKey string) error {
	relMutex.Lock()
	relMutex.Unlock()

	spaceRelFilePath := lib.CompletePath(storage.LocalSpaceRecordRelDir(), spaceKey)
	bytes, err := lib.Read(spaceRelFilePath)
	if err != nil {
		return err
	}
	recordKeyLineStr := string(bytes)
	if recordKeyLineStr != "" {
		recordKeyLineStr = strings.Replace(recordKeyLineStr, recordKey, "", -1)
		if _, err := lib.Write(storage.LocalSpaceRecordRelDir(), spaceKey, []byte(recordKeyLineStr), false); err != nil {
			return err
		}
	}
	return nil
}
