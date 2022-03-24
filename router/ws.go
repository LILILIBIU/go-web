package router

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
	chatServer.Server.Handler(ws, c)
	//go func() {
	//	for {
	//		//读取ws中的数据
	//		buf := make([]byte, 1024)
	//		mt, buf, err := ws.ReadMessage()
	//		//fmt.Printf("%s", buf)
	//		if err != nil {
	//			c.Writer.Write([]byte(err.Error()))
	//			break
	//		}
	//		//fmt.Println("client message " + string(message))
	//		//写入ws数据[]byte(time.Now().String())
	//		err = ws.WriteMessage(mt, buf)
	//		if err != nil {
	//			break
	//		}
	//		//fmt.Println("system message " + time.Now().String())
	//	}
	//}()

	select {}
}
