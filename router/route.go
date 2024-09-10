package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoxuan6/dockerproxy/handlers"
)

func Register(g *gin.Engine) {
	g.GET("/", handlers.Index)
	g.GET("/ws", handlers.Ws)

	g.NoMethod(handlers.NotMethod)
	g.NoRoute(handlers.NotFound)
}
