package xxl

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
)

// Int64ToStr int64 to str
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// StrToInt64 int64 to str
func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

//执行任务回调
func returnCall(req *RunReq, code int64, msg string) []byte {
	data := call{
		&callElement{
			LogID:      req.LogID,
			LogDateTim: req.LogDateTime,
			ExecuteResult: &ExecuteResult{
				Code: code,
				Msg:  msg,
			},
			HandleCode: int(code),
			HandleMsg:  msg,
		},
	}
	str, _ := json.Marshal(data)
	return str
}

//杀死任务返回
func returnKill(req *killReq, code int64) []byte {
	msg := ""
	if code != SuccessCode {
		msg = "Task kill err"
	}
	data := res{
		Code: code,
		Msg:  msg,
	}
	str, _ := json.Marshal(data)
	return str
}

//忙碌返回
func returnIdleBeat(code int64) []byte {
	msg := ""
	if code != SuccessCode {
		msg = "Task is busy"
	}
	data := res{
		Code: code,
		Msg:  msg,
	}
	str, _ := json.Marshal(data)
	return str
}

//通用返回
func returnGeneral() []byte {
	data := &res{
		Code: SuccessCode,
		Msg:  "",
	}
	str, _ := json.Marshal(data)
	return str
}

func md5V(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
