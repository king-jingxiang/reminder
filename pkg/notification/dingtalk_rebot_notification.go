package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/ruanxingbaozi/reminder/pkg/api"
	"net/http"
)

type DingTalkMarkdownMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type At struct {
	IsAtAll bool `json:"isAtAll"`
}

func NotificationDingTalkRebot(token string, title, content string) {
	url := fmt.Sprintf("%s%s", api.DingTalkRebotApi, token)

	message := &DingTalkMarkdownMessage{
		Msgtype: "markdown",
		Markdown: Markdown{
			Title: title,
			Text:  content,
		},
		At: At{
			IsAtAll: true,
		},
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		glog.Errorf("json marshal failed %v", err)
	}
	response, err := http.Post(url, api.DingTalkRebotContentType, bytes.NewReader(marshal))
	if err != nil {
		glog.Errorf("Call DingTalk Rebot API Failed %v", err)
	}
	glog.Infof("Notification Ding Talk Rebot Successed [%v]", response.Status)
}
