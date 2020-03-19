package web

// 响应code
type RESCODE int32

const (
	RESCODE__unkonwn RESCODE = 0
	RESCODE__200     RESCODE = 200
	RESCODE__500     RESCODE = 500
	RESCODE__400     RESCODE = 400
)

var RESCODE_name = map[int32]string{
	0:   "_unkonwn",
	200: "_200",
	500: "_500",
	400: "_400",
}
var RESCODE_value = map[string]int32{
	"_unkonwn": 0,
	"_200":     200,
	"_500":     500,
	"_400":     400,
}
