package space

import (
	"encoding/json"
	"fmt"
	"github.com/mimose/gcosy/lib"
	. "mo-for-desktop/model/space"
	. "mo-for-desktop/services/errs"
	"mo-for-desktop/services/storage"
	"sort"
	"sync"
	"time"
)

// local storage
// ../space/space_key  ---- value: Encrypted(Space)
var (
	cipherKey     = "&dapIKE$dkIp1keP"
	cipherBuilder *lib.CipherBuilder
	cipherMutex   sync.Mutex
)
var (
	allSpacesList      SpacesList
	allSpacesListMutex sync.RWMutex
)

// 获取空间数据的加密对象
func getCipherBuilder() *lib.CipherBuilder {
	if cipherBuilder == nil {
		cipherMutex.Lock()
		defer cipherMutex.Unlock()
		newCipherBuilder, err := lib.NewCipherBuilder(cipherKey)
		if err != nil {
			panic("panic when new cipher builder about Space")
		}
		cipherBuilder = newCipherBuilder
	}
	return cipherBuilder
}

// 获取所有空间
func ListAll() SpacesList {
	if allSpacesList != nil {
		return allSpacesList
	}

	allSpacesListMutex.Lock()
	defer allSpacesListMutex.Unlock()

	var spaces SpacesList
	spaceDir := storage.LocalSpaceDir()
	bytes, err := lib.ReadDirAllFiles(spaceDir)
	if err != nil {
		// TODO log
		fmt.Printf("[error] Space ListAll. %s\n", err)
		return spaces
	}
	if bytes == nil {
		// TODO log
		fmt.Printf("[info] Space ListAll. byte empty\n")
		return spaces
	}
	spaces = make(SpacesList, 0, len(bytes))
	for _, byte := range bytes {
		var space Space
		decrypted, err := getCipherBuilder().AesDecrypt(byte)
		if err != nil {
			// TODO log
			fmt.Printf("[error] Space ListAll. %s\n", err)
			return SpacesList{}
		}
		json.Unmarshal(decrypted, &space)
		spaces = append(spaces, space)
	}
	if len(spaces) > 0 {
		sort.Sort(spaces)
		allSpacesList = spaces
	}
	return allSpacesList
}

// 新增空间
func AddOne(name string) error {
	if err := uniqueSpaceName(name); err != nil {
		// TODO log
		fmt.Printf("[error] Space AddOne. uniqueSpaceName. %s\n", err.Error())
		return err
	}

	now := time.Now()
	key := storage.RandStringBytesMaskImprSrcUnsafe()
	space := Space{
		Key:        key,
		Name:       name,
		CreateTime: lib.CTime(now),
		UpdateTime: lib.CTime(now),
	}

	data, err := storage.GenerateToStorageData(space, *getCipherBuilder())
	if err != nil {
		// TODO log
		fmt.Printf("[error] Space AddOne. generateSpaceStorageData. %s\n", err.Error())
		return err
	}

	dirPath, fileName := storage.LocalSpaceDirByKey(key)
	_, wErr := lib.Write(dirPath, fileName, data, false)
	if wErr != nil {
		// TODO log
		fmt.Printf("[error] Space AdddOne. %s\n", wErr)
		return lib.NewError(WriteFile, WriteFileDesc, wErr)
	}

	allSpacesListMutex.Lock()
	defer allSpacesListMutex.Unlock()
	allSpacesList = append(allSpacesList, space)
	sort.Sort(allSpacesList)

	return nil
}

// 空间名唯一判断
func uniqueSpaceName(name string) error {
	spaces := ListAll()
	if spaces == nil {
		return nil
	}
	for _, space := range spaces {
		if space.Name == name {
			return lib.NewError(UniqueSpaceName, UniqueSpaceNameDesc, nil)
		}
	}
	return nil
}

// 生成空间内容存储字节流
