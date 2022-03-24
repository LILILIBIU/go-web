package middle

import (
	"Common/SQL"
	"Common/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		//获取claims中的user_id
		userName := claims.UserName
		var user *SQL.User
		user = SQL.Query(userName)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
