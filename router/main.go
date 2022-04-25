package router

import (
	"Common/SQL"
	"Common/middle"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func InitRouter() {
	//初始化路由
	r := gin.Default()
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板
	r.LoadHTMLGlob("templates/*")
	r.Use(middle.CORSMiddleware())
	r.GET("/", root)
	r.GET("/Radio", radio)
	r.GET("/videos2", videos2)
	r.GET("/Layout", layout)
	r.GET("/Login", login)
	r.POST("/Login", loginIn)
	r.GET("/Register", registerGet)
	r.POST("/Register", register)

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/WebSocketB", WebSocketB)
		wsGroup.GET("/WebSocket", WebSocket)
		wsGroup.GET("/WsVideos", WsVideos)
	}
	//用于执行用户操作的API

	AuthGroup := r.Group("/auth")
	{
		//Register 用户注册
		AuthGroup.GET("/register", registerGet)
		AuthGroup.POST("/register", register)
		//用户登录
		AuthGroup.GET("/Login", login)
		AuthGroup.POST("/Login", loginIn)
		AuthGroup.POST("/Info", middle.Authmiddleware(), Info)
		//查看所有待办事项
		//AuthGroup.POST("/login", login)
		//查看某一个待办事项
		AuthGroup.PUT("/todo/:id", func(c *gin.Context) {
			var todo SQL.User
			err := c.BindJSON(&todo)
			if err != nil {
				log.Printf("c.BindJSON faild!")
				return
			}
			c.JSON(http.StatusOK, todo)

		})
		//删除
		AuthGroup.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
			//DB.DeleteDB(DB, id)
		})
	}

	//默认端口8080

	err := r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		log.Printf("r.run faild!")
		return
	}
}
