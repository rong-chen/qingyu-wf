package global

type RespMsgStruct struct {
	code uint
	msg  string
}

func RespMsg(code uint, msg string) *RespMsgStruct {
	return &RespMsgStruct{
		code: code,
		msg:  msg,
	}
}
