package httpdo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// 自定义响应(10xxx, 20xxx, 30xxx, ...)
//
// 例子:
//
// 10200 => 正常, Msg
// 非10200 => 错误, Msg

const (
	Code_10200 = 10200

	Code_10400 = 10400
	Code_10401 = 10401
	Code_10403 = 10403
	Code_10404 = 10404
	Code_10405 = 10405
	Code_10406 = 10406
	Code_10408 = 10408

	Code_10500 = 10500
	Code_10503 = 10503
	Code_10505 = 10505
)

var codeMsgMap = NewCodeMsgMapDefault()

func NewCodeMsgMapDefault() map[int]string {
	return map[int]string{
		Code_10200: "请求成功",

		Code_10400: "请求参数错误",
		Code_10401: "认证失败",
		Code_10403: "无权访问",
		Code_10404: "未找到",
		Code_10405: "请求方法不支持",
		Code_10406: "请求不接受",
		Code_10408: "请求超时",

		Code_10500: "服务器错误",
		Code_10503: "服务器暂不可用",
		Code_10505: "HTTP版本不支持",
	}
}

func ResetCodeMsgMap(newCodeMsgMap map[int]string) {
	codeMsgMap = newCodeMsgMap
}

type ResData struct {
	Code     int         `json:"code"`      // 自定义code
	Msg      string      `json:"msg"`       // 提示信息
	Data     interface{} `json:"data"`      // 返回数据
	MoreInfo string      `json:"more_info"` // 一般返回失败时，捕捉的错误信息，用于开发者查看
}

func NewResData(code int, msg string, data interface{}, err error) *ResData {
	if msg == "" {
		msg = codeMsgMap[code]
	}
	minfo := ""
	if err != nil {
		minfo = err.Error()
	}
	return &ResData{
		Code:     code,
		Msg:      msg,
		Data:     data,
		MoreInfo: minfo,
	}
}

func (r *ResData) Info() string {
	return fmt.Sprintf("[%d] %s | %s", r.Code, r.Msg, r.MoreInfo)
}

func (r *ResData) JsonUnmarshal(body io.ReadCloser) error {
	if body == nil {
		return fmt.Errorf("resp body is empty")
	}
	b, err := ioutil.ReadAll(body)
	if err != nil {
		err = fmt.Errorf("read resp body err %s", err)
		return err
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		err = fmt.Errorf("resp body json unmarshal err %s", err)
		return err
	}
	return nil
}
