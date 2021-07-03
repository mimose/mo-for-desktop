package space

import (
	"fmt"
	"io/ioutil"
	. "mo-for-desktop/services/errs"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAddOne(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("recover: %s", r)
		}
	}()
	err := AddOne("测试一下新增空间")
	if err != NilError {
		t.Errorf("error append, %s", err.Error())
	}
}

func TestFile(t *testing.T) {
	dirPath := "/Users/mimose/AppData"
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = os.RemoveAll(dirPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	for i, info := range dir {
		fmt.Println(i, info)
	}
	//_, err = os.ReadDir(dirPath)
	//if err != nil {
	//	t.Errorf(err.Error())
	//}
}

func TestListAll(t *testing.T) {
	spaces := ListAll()
	if spaces == nil || len(spaces) == 0 {
		t.Error("empty space")
	}
	t.Fatalf("space detail: %s\n", spaces)
}

func TestAddAndList(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("recover: %s", r)
		}
	}()
	fmt.Println(ListAll())
	for i := 0; i < 5; i++ {
		if err := AddOne(strings.Join([]string{"空间", strconv.Itoa(i)}, "")); err != NilError {
			panic(err)
		}
	}
	fmt.Println("after add... ")
	fmt.Println(ListAll())
}
