package notify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"supervisor-event-listener/event"
)

type WebHook struct{}
type PanicMsg struct {
	MsgType string   `json:"msg_type"`
	Content *Content `json:"content"`
}
type Content struct {
	Text string `json:"text"`
}
// 飞书
func (hook *WebHook) Send(message event.Message) error {
	buf, err := json.Marshal(message)
	if err != nil {
		return err
	}

	content := &Content{Text: "报警"+ string(buf)}
	msg := &PanicMsg{
		MsgType: "text",
		Content: content,
	}

	jsonBuf, _ := json.MarshalIndent(msg, "  ", "  ")

	response, err := http.Post(Conf.WebHook.Url,"application/json",
		strings.NewReader(string(jsonBuf)))
	if err != nil {
		return err
	}

	defer response.Body.Close()
	errorMessage := ""
	if response.StatusCode != 200 {
		errorMessage = fmt.Sprintf("webhook执行失败#HTTP状态码-%d#HTTP-BODY-%s",
			response.StatusCode, response.Body)
		return errors.New(errorMessage)
	}

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return nil
}
