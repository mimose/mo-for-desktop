package record

import (
	"fmt"
	"testing"
)

func TestAddOne(t *testing.T) {
	body := `{"title":"测试标题4", "recordType":0, "content":"测试一个Notice", "noticeTime":"2021-07-10 10:00:01", "done": false}`
	//body := `{"title":"测试标题2","spaceKey":"QBvSZiJupK", "recordType":1, "content":"测试一个Note"， "done": true}`
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
	lists := Lists(1)
	for _, list := range lists {
		fmt.Println(list)
	}
	fmt.Println("=================")
	lists = Lists(0)
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

func TestDoneOne(t *testing.T) {
	var recordKey = "BmMbSOZxoQ"
	record, fileFullPath, err := getOne(recordKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("record path: %s\n", fileFullPath)
	fmt.Println(record)

	err = DoneOne(recordKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	record, fileFullPath, err = getOne(recordKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("record path: %s\n", fileFullPath)
	fmt.Println(record)

}

func TestUndoneOne(t *testing.T) {
	var recordKey = "BmMbSOZxoQ"
	record, fileFullPath, err := getOne(recordKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("record path: %s\n", fileFullPath)
	fmt.Println(record)

	err = UndoneOne(recordKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	record, fileFullPath, err = getOne(recordKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("record path: %s\n", fileFullPath)
	fmt.Println(record)

}
