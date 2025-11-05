package merr

import "github.com/zeromicro/go-zero/core/logx"

func NormalErr(err error) logx.LogField {
	return logx.LogField{
		Key:   "merr",
		Value: err.Error(),
	}
}
