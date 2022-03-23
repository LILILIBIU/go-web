package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

//允许跨域请求
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketB(c *gin.Context) {
	c.HTML(http.StatusOK, "WebSocketB.html", nil)
}

func WebSocket(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	defer ws.Close()
	go func() {
		for {
			//读取ws中的数据
			mt, message, err := ws.ReadMessage()
			if err != nil {
				c.Writer.Write([]byte(err.Error()))
				break
			}
			fmt.Println("client message " + string(message))
			//写入ws数据
			err = ws.WriteMessage(mt, []byte(time.Now().String()))
			if err != nil {
				break
			}
			fmt.Println("system message " + time.Now().String())
		}
	}()

	select {}
}
