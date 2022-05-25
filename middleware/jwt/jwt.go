package jwt

import (
	"go_blog/e"
	"go_blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		token := ""
		code = e.SUCCESS
		// Token放在请求头中
		if res, ok := c.Request.Header["Token"]; ok {
			token = res[0]
		}
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.Auth_Expire
				default:
					code = e.Auth_Error
				}
			}

			c.Set("username", claims.Username)
			c.Set("password", claims.Password)
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.MsgFlags[e.Auth_Error],
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
