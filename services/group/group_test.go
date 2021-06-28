package group

import (
	"encoding/json"
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/services/storage"
	"testing"
	"time"
)

func TestAddOne(t *testing.T) {
	err := AddOne("测试一下新增分组3")
	if err != nil {
		t.Errorf("error append, %s", err)
	}
}

func TestListAll(t *testing.T) {
	groups := ListAll()
	if groups == nil || len(groups) == 0 {
		t.Error("empty group")
	}
	t.Fatalf("group detail: %s\n", groups)
}

func Test1(t *testing.T) {
	now := time.Now()
	key := storage.RandStringBytesMaskImprSrcUnsafe()
	group := Group{
		Key:        key,
		Name:       "测试",
		CreateTime: lib.CTime(now),
		UpdateTime: lib.CTime(now),
	}
	groupJson, err1 := json.Marshal(group)
	if err1 != nil {
		t.Error(err1)
	}
	encrypt, err3 := getCipherBuilder().AesEncrypt(groupJson)
	if err3 != nil {
		t.Error(err1)
	}

	decrypted, err4 := getCipherBuilder().AesDecrypt(encrypt)
	if err4 != nil {
		t.Error(err1)
	}

	var group2 Group
	err2 := json.Unmarshal(decrypted, &group2)
	if err2 != nil {
		t.Error(err2)
	}
}
