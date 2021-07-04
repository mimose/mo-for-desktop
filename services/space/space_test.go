package space

import (
	"fmt"
	. "mo-for-desktop/services/errs"
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
	if err != nil {
		t.Errorf("error append, %s", err.Error())
	}
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
		if err := AddOne(strings.Join([]string{"空间", strconv.Itoa(i)}, "")); err != nil {
			panic(err)
		}
	}
	fmt.Println("after add... ")
	fmt.Println(ListAll())
}
