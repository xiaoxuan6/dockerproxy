package global

import "time"

var (
	NextUpdateTime time.Time
	LastUpdateTime time.Time
	AutoUpdateTime time.Duration

	Urls = make([]string, 0)
)
