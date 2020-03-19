package web

type Response struct {
	Code RESCODE     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResp(data interface{}) Response {
	return Response{Code: RESCODE__200, Data: data}
}

func FailResp(msg string, data interface{}) Response {
	return Response{Code: RESCODE__500, Msg: msg, Data: data}
}

func DenyResp(msg string, data interface{}) Response {
	return Response{Code: RESCODE__400, Msg: msg, Data: data}
}
