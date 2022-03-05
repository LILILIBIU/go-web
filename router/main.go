package router

import (
	"Common/SQL"
	"Common/common"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter(DB *sql.DB) {
	//初始化路由
	r := gin.Default()
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/", root)
	r.GET("/Radio", radio)
	r.GET("/Layout", layout)
	r.GET("/Login", login)
	r.GET("/Register", register)

	//用于执行用户操作的apifewr

	AuthGroup := r.Group("/auth")
	{
		//Register 用户注册
		AuthGroup.POST("/register", func(c *gin.Context) {
			//前端页面填写一个待办事项 点击 提交方式 请求到这里
			//1，从请求中把数据拿出来
			user := SQL.User{}
			err := c.BindJSON(&user)
			if err != nil {
				log.Printf("c.BindJSON faild!")
				return
			}
			//判断user信息是否符合格式
			isOk, errMsg := SQL.TodoIsOK(&user)
			if !isOk {
				c.JSON(http.StatusOK, gin.H{"code": 500, "errMsg": errMsg})
				return
			}
			//判断用户是否建立成功 建立失败返回150 用户存在返回100 成功建立回200
			OkValue := SQL.CreatAccount(DB, &user)
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
		})
		//查看所有待办事项
		AuthGroup.POST("/login", func(c *gin.Context) {
			user := SQL.User{}
			err := c.BindJSON(&user)
			if err != nil {
				log.Printf("c.BindJSON faild!")
				return
			}

		})
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

	err := r.Run(":8080")
	if err != nil {
		log.Printf("r.run faild!")
		return
	}
}
