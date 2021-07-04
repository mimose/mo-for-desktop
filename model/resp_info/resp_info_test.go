package resp_info

import (
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/model/errs"
	"mo-for-desktop/model/space"
	"testing"
	"time"
)

func TestError(t *testing.T) {
	info := Error(lib.NewError(errs.AddRecordBodyError, "未知异常", nil))
	//info := Error(errors.New("未知异常"))
	t.Log(info)
}

func TestSuccess(t *testing.T) {
	success := Success(space.Space{
		Key:        "abc",
		Name:       "name",
		CreateTime: lib.CTime(time.Now()),
		UpdateTime: lib.CTime(time.Now()),
	})
	t.Log(success)
}

func TestCustom(t *testing.T) {
	custom := Custom(123456, "test", space.Space{
		Key:        "abc",
		Name:       "name",
		CreateTime: lib.CTime(time.Now()),
		UpdateTime: lib.CTime(time.Now()),
	})
	t.Log(custom)
}
