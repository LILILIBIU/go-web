package router

import (
	"Common/SQL"
	"Common/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", nil)
}
func register(c *gin.Context) {
	//前端页面填写一个待办事项 点击 提交方式 请求到这里
	//1，从请求中把数据拿出来
	user := SQL.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("c.BindJSON faild!")
		return
	}
	fmt.Println("%#v", user)
	//判断user信息是否符合格式
	isOk, errMsg := SQL.TodoIsOK(&user, true)
	if !isOk {
		c.JSON(http.StatusOK, gin.H{"code": 500, "errMsg": errMsg})
		return
	}
	//判断用户是否建立成功 建立失败返回150 用户存在返回100 成功建立回200
	OkValue := SQL.CreatAccount(&user)
	if OkValue != 200 {
		if OkValue == 150 {
			c.JSON(500, gin.H{"code": 500, "errMsg": "系统错误", "errCode": string(OkValue)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 500, "errMsg": "用户已存在"})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token err:%v", err)
		return
	}
	log.Println("%v", user)
	c.JSON(http.StatusOK, gin.H{"code": 200, "token": token})
}
