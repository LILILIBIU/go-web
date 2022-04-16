package main

import (
	"Common/SQL"
	"Common/common/chatServer"
	"Common/config"
	"Common/router"
	"fmt"
)

func main() {
	//初始化配置文件
	config.InitConfig()
	//初始化数据库
	SQL.Init()
	//初始化chatServer
	chatServer.InitChatServer()
	//初始化路由
	router.InitRouter()
	fmt.Println("Hello World!")
}
