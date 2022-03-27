package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//https://open.feishu.cn/open-apis/im/v1/messages
func SendMessage(empID string, msg_type string, content *MsgContent) error {
	url := "https://open.feishu.cn/open-apis/im/v1/messages" + "?receive_id_type=user_id"
	msg := struct {
		EmpID    string      `json:"receive_id"`
		Content  *MsgContent `json:"content"`
		Msg_type string      `json:"msg_type"`
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
	content string
}

func (slf *TextMsg) NewMsg(messages ...interface{}) interface{} {
	message := "{\"text\":\" "
	for i, item := range messages {
		message = message + item.(string) + " \\n "
		if i == len(messages)-1 {
			message = message + item.(string) + " \"}"
		}
	}
	return TextMsg{
		content: message,
	}
}
