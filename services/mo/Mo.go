package mo

import (
	"github.com/mimose/gcosy/lib"
	"github.com/wailsapp/wails"
	"mo-for-desktop/services/group"
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

// ================ group
// list group
func (mo *Mo) ListGroups() []group.Group {
	return group.ListAll()
}

// ================ group

// ================ record
func (mo *Mo) ListRecord(groupKey, recordType int) []record.Record {
	return record.Lists(groupKey, recordType)
}

// ================ record
