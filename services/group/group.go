package group

import (
	"encoding/json"
	"fmt"
	"github.com/mimose/gcosy/lib"
	. "mo-for-desktop/model/group"
	"mo-for-desktop/services/storage"
	"sort"
	"sync"
	"time"
)

// local storage
// ../group/group_[encode(key)]  ---- value: encode(Group)
// ../group/recordRel/group_[encode(key) --- value: encode(List: record_local_storage_key)

var (
	cipherKey     = "&dapIKE$dkIp1keP"
	cipherBuilder *lib.CipherBuilder
	cipherMutex   sync.Mutex
)
var (
	allGroupsList      GroupsList
	allGroupsListMutex sync.RWMutex
)

func getCipherBuilder() *lib.CipherBuilder {
	if cipherBuilder == nil {
		cipherMutex.Lock()
		defer cipherMutex.Unlock()
		newCipherBuilder, err := lib.NewCipherBuilder(cipherKey)
		if err != nil {
			panic("panic when new cipher builder about Group")
		}
		cipherBuilder = newCipherBuilder
	}
	return cipherBuilder
}

func ListAll() GroupsList {
	if allGroupsList != nil {
		return allGroupsList
	}

	allGroupsListMutex.Lock()
	defer allGroupsListMutex.Unlock()

	var groups GroupsList
	groupDir := storage.LocalGroupDir()
	bytes, err := lib.ReadDirAllFiles(groupDir)
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
	groups = make([]Group, 0, len(bytes))
	for _, byte := range bytes {
		var group Group
		decrypted, err := getCipherBuilder().AesDecrypt(byte)
		if err != nil {
			// TODO log
			fmt.Printf("[error] Group ListAll. %s", err)
			return []Group{}
		}
		json.Unmarshal(decrypted, &group)
		groups = append(groups, group)
	}
	if len(groups) > 0 {
		sort.Sort(groups)
		allGroupsList = groups
	}
	return allGroupsList
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
	groupJson, marshalErr := json.Marshal(group)
	if marshalErr != nil {
		// TODO log
		fmt.Printf("[error] Group AddOne. %s", marshalErr)
		return marshalErr
	}
	data, err := getCipherBuilder().AesEncrypt(groupJson)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Group AddOne. %s", err)
		return err
	}
	dirPath, fileName := storage.LocalGroupDirByKey(key)
	err = lib.Write(dirPath, fileName, data, false)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Group AdddOne. %s", err)
	}

	allGroupsListMutex.Lock()
	defer allGroupsListMutex.Unlock()
	allGroupsList = append(allGroupsList, group)
	sort.Sort(allGroupsList)

	return nil
}
