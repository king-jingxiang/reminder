package cache

import "github.com/ruanxingbaozi/reminder/pkg/api"

type ReminderCache struct {
	Path       string
	Sync       bool
	SyncPeriod int32
	// [user][]api.Reminder
	Reminders map[string][]api.Reminder
}

func New() *ReminderCache {
	return &ReminderCache{
		Path:       "/data",
		Sync:       false,
		SyncPeriod: 10,
	}
}

func SyncToDataPath() {

}
