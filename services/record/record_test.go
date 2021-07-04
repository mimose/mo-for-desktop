package record

import (
	"fmt"
	"testing"
)

func TestAddOne(t *testing.T) {
	body := `{"title":"测试标题555","spaceKey":"pyqCJMRebP", "recordType":0, "content":"测试内容122122333456987654yvubhinjomk"}`
	err := AddOne(body)
	if err != nil {
		t.Errorf(err.Error())
	}
	file := ListAll()
	for _, record := range file {
		fmt.Println(record)
	}
}

func TestLists(t *testing.T) {
	lists := Lists("pyqCJMRebP", 1)
	for _, list := range lists {
		fmt.Println(list)
	}
	fmt.Println("=================")
	lists = Lists("pyqCJMRebP", 0)
	for _, list := range lists {
		fmt.Println(list)
	}
}

func TestRemoveOne(t *testing.T) {
	err := RemoveOne("brodbYwfie")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGroup(t *testing.T) {
	all := ListAll()
	group := all.Group()
	t.Log(group)
}
