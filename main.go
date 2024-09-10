package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xiaoxuan6/dockerproxy/cron"
	"github.com/xiaoxuan6/dockerproxy/global"
	"github.com/xiaoxuan6/dockerproxy/router"
	"github.com/xiaoxuan6/dockerproxy/services"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	_ = godotenv.Load()

	setUpdateTime()
	services.IndexService.FetchUrls()
	cron.Start()

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	g.Use(gzip.Gzip(gzip.DefaultCompression))
	g.Static("/static", "./static")
	g.LoadHTMLGlob("templates/*")

	router.Register(g)
	if err := g.Run(":9101"); err != nil {
		log.Panic(err.Error())
	}
}

func setUpdateTime() {
	var autoUpdateTime string
	if autoUpdateTime = os.Getenv("AUTO_UPDATE_TIME"); autoUpdateTime == "" {
		autoUpdateTime = "30"
	}

	autoUpdateTimeInt, _ := strconv.Atoi(autoUpdateTime)
	if autoUpdateTimeInt < 1 {
		autoUpdateTimeInt = 30
	}

	global.AutoUpdateTime = time.Minute * time.Duration(autoUpdateTimeInt)
	global.LastUpdateTime = time.Now()
	global.NextUpdateTime = global.LastUpdateTime.Add(global.AutoUpdateTime)
}
