package main

import (
	"Common/SQL"
	"Common/common/chatServer"
	"Common/router"
	"log"
	"os"
)

func init() {
	log.SetPrefix("项目：")
	f, err := os.OpenFile("./XX.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
}
func main() {
	//初始化数据库
	SQL.Init()
	//初始化路由
	chatServer.InitChatServer()
	router.InitRouter()
	//初始化chatServer

	//chatServer.Server.ListenMessage()
}
