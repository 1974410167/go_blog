package views

import (
	"go_blog/e"
	"go_blog/models"
	"go_blog/util"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.INVALID_PARAMS],
		})
		return
	}
	var user1 models.User
	models.DB.Model(&models.User{}).Where("username = ?", user.Username).First(&user1)
	if user1.ID == 0 {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if user1.Password != user.Password {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": "密码错误",
		})
		return
	}
	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": "生成token失败",
		})
	}
	c.JSON(e.SUCCESS, gin.H{
		"token":   token,
		"message": "登录成功",
	})
}
