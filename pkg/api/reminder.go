package api

import "time"

type Reminder struct {
	UserName     string       `json:"userName"`
	Date         *time.Time   `json:"date"`
	DateType     string       `json:"dateType"` // 阴历，阳历
	ReminderType string       `json:"dateType"` // daily,monthly,yearly
	Notification Notification `json:"notification"`
	// 确认收到;
	OpenCallBack string `json:"openCallBack"`
	Content      string `json:"content"` // 提醒内容
}

type Notification struct {
	Email string `json:"email"`

	DingTalkRebotToken string `json:"dingTalkRebotToken"`
	WechatGroup        string `json:"wechatGroup"`
	NotificationNum    int32  `json:"notificationNum"`
	// 提前通知;直到超过约定时间；通知频率程序中指定
	BeforeNofification bool
	BeforeVaule        int32
}
