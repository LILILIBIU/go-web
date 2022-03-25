package main

import (
	"Common/SQL"
	"Common/common/chatServer"
	"Common/config"
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
	//初始化配置文件
	config.InitConfig()
	//初始化数据库
	SQL.Init()
	//初始化chatServer
	chatServer.InitChatServer()
	//初始化路由
	router.InitRouter()

}
