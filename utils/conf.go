package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

//定义接收配置文件的结构体
type DataBaseConnection struct {
	IpAddress    string
	Port         int
	UserName     string
	Password     int
	DataBaseName string
}

func main() {
	config := viper.New()
	//配置文件名（不带扩展名）
	config.SetConfigName("conf")
	//在项目中查找配置文件的路径，可以使用相对路径，也可以使用绝对路径
	config.AddConfigPath("/config")
	//设置文件类型，这里是yaml文件
	config.SetConfigType("yaml")
	//定义用于接收配置文件的变量
	var configData DataBaseConnection
	//查找并读取配置文件
	err := config.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := config.Unmarshal(&configData); err != nil { // 读取配置文件转化成对应的结构体错误
		panic(fmt.Errorf("read config file to struct err: %s \n", err))
	}
	//控制台打印输出配置文件读取的值
	fmt.Println(configData) //{127.0.0.1 3306 root 123456 go_test}
}
