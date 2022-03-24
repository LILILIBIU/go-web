package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func radio(c *gin.Context) {
	c.HTML(http.StatusOK, "Radio.html", nil)
}

func layout(c *gin.Context) {
	c.HTML(http.StatusOK, "Layout.html", nil)
}
