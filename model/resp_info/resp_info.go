package resp_info

import (
	"github.com/mimose/gcosy/lib"
	. "mo-for-desktop/model/errs"
)

type RespInfo struct {
	Code int         `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

func Error(err error) RespInfo {
	switch err.(type) {
	case *lib.CError:
		cErr := err.(*lib.CError)
		return RespInfo{
			Code: cErr.Code,
			Desc: cErr.Desc,
			Data: nil,
		}
	default:
		return RespInfo{
			Code: UnknownError,
			Desc: err.Error(),
			Data: nil,
		}
	}
}

func Success(data interface{}) RespInfo {
	return RespInfo{
		Code: SuccessCode,
		Desc: "success",
		Data: data,
	}
}

func Custom(code int, desc string, data interface{}) RespInfo {
	return RespInfo{
		Code: code,
		Desc: desc,
		Data: data,
	}
}
