package views

import (
	"go_blog/e"
	"go_blog/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.INVALID_PARAMS],
		})
		return
	}
	res := models.DB.Model(&models.User{}).Create(&user)
	if res.Error != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.Create_ERROR],
		})
		return
	}
	c.JSON(e.SUCCESS, gin.H{
		"message": "注册成功",
	})
}
