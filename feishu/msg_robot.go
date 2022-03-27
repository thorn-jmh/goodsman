//消息发送器

package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//若要新增消息类型,请新建结构体
//然后为其实现 NewMsg

//https://open.feishu.cn/open-apis/im/v1/messages
func SendMessage(empID string, msg_type string, content MsgContent) error {
	url := "https://open.feishu.cn/open-apis/im/v1/messages" + "?receive_id_type=user_id"
	msg := struct {
		EmpID    string     `json:"receive_id"`
		Content  MsgContent `json:"content"`
		Msg_type string     `json:"msg_type"`
	}{
		EmpID:    empID,
		Content:  content,
		Msg_type: msg_type,
	}

	accessToken, err := TenantTokenManager.GetAccessToken()
	if err != nil {
		return err
	}
	reqbody, _ := json.Marshal(msg)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(reqbody))
	resp, err := CommonClient.Do(req, accessToken)
	if err != nil {
		return err
	}
	result := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}
	json.Unmarshal(resp, &result)
	if result.Code != 0 {
		return errors.New(result.Msg)
	}
	return nil
}

type MsgContent interface {
	NewMsg(messages ...interface{}) interface{}
}

type TextMsg struct {
	Content string
}

func (slf *TextMsg) NewMsg(messages ...interface{}) interface{} {
	items, _ := messages[0].([]string)
	message := "{\"text\":\" "
	for i, item := range items {
		message = message + item + " \\n "
		if i == len(messages)-1 {
			message = message + item + " \"}"
		}
	}
	return message
}
