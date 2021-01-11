package notification

import (
	"github.com/ruanxingbaozi/reminder/pkg/api"
	"testing"
	"time"
)

func TestNotificationDingTalkRebot(t *testing.T) {
	token := "16968615dd1dc2909f7e6a440ae310462b6eb5e19855e5b5dacdc2613aa08840"
	title := "节日提醒"
	notify := api.Notify{
		Date:     time.Now(),
		DateType: api.Date_Type_Lunar,
		Event:    "节日提醒",
	}
	content := GeneratorMarkdownNotifyContent(notify)
	NotificationDingTalkRebot(token, title, content)
}
