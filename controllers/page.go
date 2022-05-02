package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Radio(c *gin.Context) {
	c.HTML(http.StatusOK, "Radio.html", nil)
}
func Videos2(c *gin.Context) {
	c.HTML(http.StatusOK, "videos2.html", nil)
}

func Layout(c *gin.Context) {
	c.HTML(http.StatusOK, "Layout.html", nil)
}
