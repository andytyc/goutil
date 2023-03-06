package httpdo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type ResData struct {
	Code     int         `json:"code"`      // 自定义code
	Message  string      `json:"message"`   // 提示信息
	Data     interface{} `json:"data"`      // 返回数据
	MoreInfo string      `json:"more_info"` // 一般返回失败时，捕捉的错误信息，用于开发者查看
}

func (r *ResData) Info() string {
	return fmt.Sprintf("[%d] %s | %s", r.Code, r.Message, r.MoreInfo)
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
