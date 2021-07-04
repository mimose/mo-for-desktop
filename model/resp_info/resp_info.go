package resp_info

import (
	"github.com/mimose/gcosy/lib"
	. "mo-for-desktop/model/errs"
)

type RespInfo struct {
	code int         `json:code`
	desc string      `json:desc`
	data interface{} `json:data`
}

func Error(err error) *RespInfo {
	switch err.(type) {
	case *lib.CError:
		cErr := err.(*lib.CError)
		return &RespInfo{
			code: cErr.Code,
			desc: cErr.Desc,
			data: nil,
		}
	default:
		return &RespInfo{
			code: UnknownError,
			desc: err.Error(),
			data: nil,
		}
	}
}

func Success(data interface{}) *RespInfo {
	return &RespInfo{
		code: SuccessCode,
		desc: "success",
		data: data,
	}
}

func Custom(code int, desc string, data interface{}) *RespInfo {
	return &RespInfo{
		code: code,
		desc: desc,
		data: data,
	}
}
