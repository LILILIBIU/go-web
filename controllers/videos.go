package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WsVideos(c *gin.Context) {
	c.HTML(http.StatusOK, "WsVideos.html", nil)
}
