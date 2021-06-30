package mo

import (
	"github.com/mimose/gcosy/lib"
	"github.com/wailsapp/wails"
	. "mo-for-desktop/model/record"
	. "mo-for-desktop/model/space"
	"mo-for-desktop/services/record"
	"mo-for-desktop/services/space"
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

// ================ space
// list space
func (mo *Mo) ListSpaces() SpacesList {
	return space.ListAll()
}

// ================ space

// ================ record
func (mo *Mo) ListRecord(spaceKey string, recordType int) RecordsList {
	return record.Lists(spaceKey, recordType)
}

// ================ record
