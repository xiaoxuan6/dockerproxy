package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/xiaoxuan6/dockerproxy/services"
	"os"
)

func Start() {
	c := cron.New(cron.WithSeconds())
	_, _ = c.AddFunc(os.Getenv("CRON_SPEC"), func() {
		services.IndexService.FetchUrls()
	})
	c.Start()
}
