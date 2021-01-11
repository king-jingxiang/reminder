package chinese_calendar

import (
	"time"

	"github.com/ruanxingbaozi/chinese-calendar-golang/calendar"
)

func LunarConvertSolar(t time.Time) time.Time {
	isLeapMonth := false
	if t.Year()%4 == 0 {
		isLeapMonth = true
	}
	// 农历(最后一个参数表示是否闰月)
	c := calendar.ByLunar(int64(t.Year()), int64(t.Month()),
		int64(t.Day()), int64(t.Hour()), int64(t.Minute()), int64(t.Second()), isLeapMonth)

	return *c.Time
}
