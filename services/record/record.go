package record

import (
	"encoding/json"
	"fmt"
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/model/errs"
	. "mo-for-desktop/model/record"
	"mo-for-desktop/services/space_record_rel"
	"mo-for-desktop/services/storage"
	"sort"
	"sync"
	"time"
)

// local storage
// ../record/record_key  ---- value: Encrypted(Record)
var (
	cipherKey     = "A&ju7iUE=Tep9j6o"
	cipherBuilder *lib.CipherBuilder
	cipherMutex   sync.Mutex
)

// 获取记录数据的加密对象
func getCipherBuilder() *lib.CipherBuilder {
	if cipherBuilder == nil {
		cipherMutex.Lock()
		defer cipherMutex.Unlock()
		newCipherBuilder, err := lib.NewCipherBuilder(cipherKey)
		if err != nil {
			panic("panic when new cipher builder about Record")
		}
		cipherBuilder = newCipherBuilder
	}
	return cipherBuilder
}

// 根据空间key + 记录类型 获取所有记录
func Lists(spaceKey string, recordType int) RecordsList {
	var records RecordsList
	if spaceKey == "" {
		records = ListAll()
	} else {
		records = ListsBySpaceKey(spaceKey)
	}
	if len(records) == 0 {
		return records
	}
	if recordType >= 0 {
		records = records.FilterRecordListByRecordType(recordType)
	}
	return records
}

// 根据空间key 获取所有记录
func ListsBySpaceKey(spaceKey string) RecordsList {
	var records RecordsList
	recordKeys, err := space_record_rel.GetRel(spaceKey)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Record ListsBySpaceKey. spaceKey: %s, err: %s\n", spaceKey, err)
		return records
	}
	if len(recordKeys) == 0 {
		return records
	}

	records = make(RecordsList, 0, len(recordKeys))
	for _, recordKey := range recordKeys {
		record, _, err := getOne(recordKey)
		if err != nil {
			// TODO log
			fmt.Printf("[error] Record ListsBySpaceKey. spaceKey: %s, err: %s\n", spaceKey, err)
			return RecordsList{}
		}
		records = append(records, record)
	}
	return records
}

// 获取所有记录
func ListAll() RecordsList {
	var records RecordsList
	recordDir := storage.LocalRecordDir()
	bytes, err := lib.ReadDirAllFiles(recordDir)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Record ListAll. %s\n", err)
		return records
	}
	if bytes == nil {
		// TODO log
		fmt.Printf("[info] Record ListAll. byte empty\n")
		return records
	}
	records = make(RecordsList, 0, len(bytes))
	for _, byte := range bytes {
		var record Record
		decrypted, err := getCipherBuilder().AesDecrypt(byte)
		if err != nil {
			// TODO log
			fmt.Printf("[error] Record ListAll. %s\n", err)
			return RecordsList{}
		}
		json.Unmarshal(decrypted, &record)
		records = append(records, record)
	}
	if len(records) > 0 {
		sort.Sort(records)
	}
	return records
}

// 获取某条记录
func getOne(recordKey string) (Record, string, error) {
	dirPath, fileName := storage.LocalRecordDirByKey(recordKey)
	fileFullPath := lib.CompletePath(dirPath, fileName)
	byte, err := lib.Read(fileFullPath)
	if err != nil {
		return Record{}, fileFullPath, err
	}
	var record Record
	decrypted, err := getCipherBuilder().AesDecrypt(byte)
	if err != nil {
		return Record{}, fileFullPath, err
	}
	json.Unmarshal(decrypted, &record)
	return record, fileFullPath, nil
}

// 新增记录
func AddOne(body string) error {
	record, err := bodyToRecord(body)
	if err != nil {
		return err
	}

	now := time.Now()
	key := storage.RandStringBytesMaskImprSrcUnsafe()
	record.Key = key
	record.CreateTime = lib.CTime(now)
	record.UpdateTime = lib.CTime(now)

	data, err := storage.GenerateToStorageData(record, *getCipherBuilder())
	if err != nil {
		// TODO log
		fmt.Printf("[error] Record AddOne. generateRecordStorageData. %s\n", err.Error())
		return err
	}

	dirPath, fileName := storage.LocalRecordDirByKey(key)
	fileFullPath, wErr := lib.Write(dirPath, fileName, data, false)
	if wErr != nil {
		// TODO log
		fmt.Printf("[error] Record AdddOne. %s\n", wErr)
		return lib.NewError(errs.WriteFile, errs.WriteFileDesc, wErr)
	}

	// add the space_record_rel
	defer func() {
		// delete the record file
		if r := recover(); r != nil {
			// TODO log
			fmt.Printf("[error] panic happen, will delete the record. %s\n", err)
			lib.DeleteFile(fileFullPath)
		}
	}()
	if err := space_record_rel.AppendRel(record.SpaceKey, record.Key); err != nil {
		// TODO log
		fmt.Printf("[error] Record AppendRel. %s\n", err)
		panic(err)
	}
	return nil
}

func bodyToRecord(body string) (Record, error) {
	if body == "" {
		// TODO log
		fmt.Printf("[error] Record AddOne. body empty\n")
		return Record{}, lib.NewError(errs.AddRecordBodyError, "转换前数据为空", nil)
	}
	var record Record
	err := json.Unmarshal([]byte(body), &record)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Record AddOne. to Struct fail. %s\n", err)
		return Record{}, lib.NewError(errs.AddRecordBodyError, "数据转换异常", err)
	}
	if record == (Record{}) {
		// TODO log
		fmt.Printf("[error] Record AddOne. to Struct fail\n")
		return Record{}, lib.NewError(errs.AddRecordBodyError, "转换后数据为空", nil)
	}
	if err = record.Validate(); err != nil {
		// TODO log
		fmt.Printf("[error] Record AddOne. record validate fail, %s\n", err)
		return Record{}, lib.NewError(errs.AddRecordBodyError, "转换后数据校验不通过", err)
	}
	return record, nil
}

// 删除记录
func RemoveOne(recordKey string) error {
	record, fileFullPath, err := getOne(recordKey)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Record RemoveOne. recordKey: %s, err: %s\n", recordKey, err)
		return err
	}
	err = space_record_rel.RemoveRel(record.SpaceKey, recordKey)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Record RemoveOne. recordKey: %s, err: %s\n", recordKey, err)
		return err
	}

	defer func() {
		// resave the rel
		if r := recover(); r != nil {
			if err := space_record_rel.AppendRel(record.SpaceKey, record.Key); err != nil {
				// TODO log
				fmt.Printf("[error] Record RemoveOne --> Resave AppendRel. spaceKye: %s, recordKey: %s, err: %s\n", record.SpaceKey, recordKey, err)
				panic(err)
			}
		}
	}()

	if err := lib.DeleteFile(fileFullPath); err != nil {
		// TODO log
		fmt.Printf("[error] Record RemoveOne. recordKey, %s, err: %s\n", recordKey, err)
		panic(err)
	}
	return nil
}
