package main

import (
	"Common/SQL"
	"Common/common/chatServer"
	"Common/router"
	"github.com/spf13/viper"
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
	InitConfig()
	//初始化数据库
	SQL.Init()
	//初始化chatServer
	chatServer.InitChatServer()
	//初始化路由
	router.InitRouter()

}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置出错！")
	}
}
