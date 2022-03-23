package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WsVideos(c *gin.Context) {
	c.HTML(http.StatusOK, "WsVideos.html", nil)
}
func WsVideos2(c *gin.Context) {

}
