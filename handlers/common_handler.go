package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, fmt.Sprintf("method [%s] not allowed", c.Request.Method))
	c.Abort()
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, fmt.Sprintf("route [%s] not found", c.Request.URL.Path))
	c.Abort()
}
