package controllers

import (
	live "Common/live/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Name struct {
	name string `json:"name"`
}

func LiveRegister(c *gin.Context) {
	channel, room := live.GetPassword()
	if channel == "-1" && room == "-1" {
		c.JSON(http.StatusOK, gin.H{"errMsg": "获取失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"channel": channel, "room": room})
}
