package notification

import (
	"fmt"
	"github.com/ruanxingbaozi/reminder/pkg/api"
	"github.com/ruanxingbaozi/reminder/pkg/chinese_calendar"
	"time"
)

func GeneratorNotifyContent(notify api.Notify) string {
	notifyDate := notify.Date
	content := fmt.Sprintf("当前时间: [%v], 日期类型为: [%v]", notify.Date, notify.DateType)
	if notify.DateType == api.Date_Type_Lunar {
		solar := chinese_calendar.LunarConvertSolar(notify.Date)
		notifyDate = solar
		content += fmt.Sprintf(", 对应阳历日期: [%v]", solar)
	}
	content += fmt.Sprintf(", 提醒内容为: [%v]", notify.Event)

	sub := time.Now().Sub(notifyDate)
	content += fmt.Sprintf(", 还剩: [%vH]", sub.Hours())
	return content
}
func GeneratorMarkdownNotifyContent(notify api.Notify) string {
	notifyDate := notify.Date
	content := ""
	if notify.DateType == api.Date_Type_Lunar {
		solar := chinese_calendar.LunarConvertSolar(notify.Date)
		notifyDate = solar
		content += fmt.Sprintf("当前时间: **[%v]**, 对应的阴历时间为: **[%v]**", solar, notify.Date)
	} else {
		content += fmt.Sprintf("当前时间: **[%v]**", notify.Date)
	}
	content += fmt.Sprintf(", 提醒内容为: **[%v]**", notify.Event)

	sub := time.Now().Sub(notifyDate)
	content += fmt.Sprintf(", 还剩: **[%vH]**", sub.Hours())
	return content
}
