package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/xiaoxuan6/dockerproxy/global"
	"github.com/xiaoxuan6/dockerproxy/services"
	"log"
	"net/http"
	"sync"
	"time"
)

var data []services.Item

func Index(c *gin.Context) {
	if len(data) < 1 {
		data = services.IndexService.FetchUrlsWithVerify()
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"data":             data,
		"sub_data":         int(global.NextUpdateTime.Sub(time.Now()).Seconds()),
		"last_update_time": global.LastUpdateTime.Format(time.DateTime),
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return false
		}

		return true
	},
}

func Ws(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade failed: %v", err)
		return
	}
	defer ws.Close()

	var mu sync.Mutex
	nextUpdateDuration := global.NextUpdateTime.Sub(time.Now())
	if nextUpdateDuration <= 0 {
		nextUpdateDuration = time.Duration(1) * time.Second
	}

	ticker := time.NewTicker(nextUpdateDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			mu.Lock()
			urls := services.IndexService.FetchUrlsWithVerify()
			mu.Unlock()

			global.LastUpdateTime = time.Now()
			global.NextUpdateTime = global.LastUpdateTime.Add(global.AutoUpdateTime)

			if err = ws.WriteJSON(gin.H{
				"data":             urls,
				"sub_data":         int(global.NextUpdateTime.Sub(time.Now()).Seconds()),
				"last_update_time": global.LastUpdateTime.Format(time.DateTime),
			}); err != nil {
				log.Printf("Error sending message or Connection closed: %v", err)
				return
			}

			data = data[:0]
			data = append(data, urls...)

			newDuration := time.Duration(global.NextUpdateTime.Sub(time.Now()).Seconds()+1) * time.Second
			ticker.Reset(newDuration)
		}
	}
}
