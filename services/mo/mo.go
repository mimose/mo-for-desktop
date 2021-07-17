package mo

import (
	"github.com/mimose/gcosy/lib"
	"github.com/wailsapp/wails"
	. "mo-for-desktop/model/resp_info"
	"mo-for-desktop/services/record"
)

type Mo struct {
	Version string
	Program *lib.Program
	Runtime *wails.Runtime
}

// new moProgram
func New(p *lib.Program) (*Mo, error) {
	// get version
	//version, errStr, errCode, err := p.Command("version")
	//if err != nil {
	//	return nil, lib.NewError(errCode, errStr, err)
	//}
	return &Mo{
		Version: "0.0.1",
		Program: p,
	}, nil
}

// wails init
func (mo *Mo) WailsInit(runtime *wails.Runtime) error {
	mo.Runtime = runtime
	runtime.Log.New("").Debug(mo.Version)
	return nil
}

// ================ record
// 新增记录
func (mo *Mo) NewRecord(body string) RespInfo {
	err := record.AddOne(body)
	if err != nil {
		return Error(err)
	}
	return Success(nil)
}

// 删除记录
func (mo *Mo) RemoveRecord(recordKey string) RespInfo {
	err := record.RemoveOne(recordKey)
	if err != nil {
		return Error(err)
	}
	return Success(nil)
}

// 获取记录列表
func (mo *Mo) ListRecord(recordType int) RespInfo {
	return Success(record.Lists(recordType))
}

// 完成一条记录
func (mo *Mo) DoneRecord(recordKey string) RespInfo {
	err := record.DoneOne(recordKey)
	if err != nil {
		return Error(err)
	}
	return Success(nil)
}

// 重设一条记录
func (mo *Mo) UndoneRecord(recordKey string) RespInfo {
	err := record.UndoneOne(recordKey)
	if err != nil {
		return Error(err)
	}
	return Success(nil)
}

// ================ record
