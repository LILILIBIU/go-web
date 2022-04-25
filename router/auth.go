package router

import (
	"Common/SQL"
	"Common/common"
	"Common/dto"
	"Common/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	username string
	password string
}

func registerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "Register.html", nil)
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
	//判断user信息是否符合格式
	isOk, errMsg := SQL.ListIsOK(&user, true)
	if !isOk {
		response.Response(c, http.StatusOK, 500, nil, errMsg)
		return
	}
	//判断用户是否建立成功 建立失败返回150 用户存在返回100 成功建立回200
	OkValue := SQL.CreatAccount(&user)
	if OkValue != 200 {
		if OkValue == 150 {
			response.Response(c, http.StatusOK, 500, nil, "用户建立失败")
			return
		}
		response.Response(c, http.StatusOK, 500, nil, "用户已存在")
		return
	}

	token, err := common.ReleaseToken(&user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token err:%v", err)
		return
	}
	log.Println("%v", user)
	response.Response(c, http.StatusOK, 500, gin.H{"token": token}, "注册成功")
}
func login(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", nil)
}
func loginIn(c *gin.Context) {
	//前端页面填写一个待办事项 点击 提交方式 请求到这里
	//1，从请求中把数据拿出来
	user := SQL.User{}
	user.Name = c.PostForm("username")
	user.Password = c.PostForm("password")
	//decoder := json.NewDecoder(c.Request.Body)
	//var luser User
	//err := decoder.Decode(&luser)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(luser)
	//user.Name = luser.username
	//user.Password = luser.password
	//判断user信息是否符合格式
	isOk, errMsg := SQL.ListIsOK(&user, false)
	if !isOk {
		response.Response(c, http.StatusOK, 500, nil, errMsg)
		return
	}
	u := SQL.Query(user.Name)
	if u.Password != user.Password {
		log.Printf("信息错误！")
	} else {
		token, err := common.ReleaseToken(&user)
		if err != nil {
			response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
			log.Printf("token err:%v", err)
			return
		}
		log.Printf("%v\n", user)
		fmt.Printf("\n%v\n", token)
		//response.Response(c, http.StatusOK, 200, gin.H{"token": token}, "登陆成功")
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

}
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Response(c, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(*SQL.User))}, "登陆成功")
}
