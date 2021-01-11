package api

import "time"

type Notify struct {
	Date     time.Time
	DateType string
	Event    string
}
