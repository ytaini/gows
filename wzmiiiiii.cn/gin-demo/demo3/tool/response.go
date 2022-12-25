// @author: wzmiiiiii
// @since: 2022/12/26 01:05:00
// @desc: TODO

package tool

type Code uint8

const (
	Fail Code = iota
	Success
)

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func NewResponse(code Code, data any, msg string) *Response {
	return &Response{Code: code, Data: data, Msg: msg}
}
