package ws

import (
	"Common/common/chatServer"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

//允许跨域请求
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {
	c.HTML(http.StatusOK, "WebSocket.html", nil)
}

func WebSocketB(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	defer ws.Close()
	chatServer.Server.Handler(ws, c)
}
