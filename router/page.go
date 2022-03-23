package router

import (
	"Common/SQL"
	"github.com/gin-gonic/gin"
	"net/http"
)

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func radio(c *gin.Context) {
	c.HTML(http.StatusOK, "Radio.html", nil)
}

func layout(c *gin.Context) {
	c.HTML(http.StatusOK, "Layout.html", nil)
}

func loginIn(c *gin.Context) {
	//前端页面填写一个待办事项 点击 提交方式 请求到这里
	//1，从请求中把数据拿出来
	user := SQL.User{}
	user.Name = c.PostForm("username")
	user.Password = c.PostForm("password")
	//判断user信息是否符合格式
	isOk, errMsg := SQL.TodoIsOK(&user, false)
	if !isOk {
		c.JSON(http.StatusOK, gin.H{"code": 500, "errMsg": errMsg})
		return
	}

}
