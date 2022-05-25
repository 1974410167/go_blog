package views

import (
	"fmt"
	"go_blog/e"
	"go_blog/models"

	"github.com/gin-gonic/gin"
)

type TagViews struct {
}

func (t *TagViews) Create(c *gin.Context) {
	var tag models.Tag
	err := c.ShouldBind(&tag)
	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.INVALID_PARAMS],
		})
		return
	}
	res := models.DB.Create(&tag)
	if res.Error != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.Create_ERROR],
		})
		return
	}
	c.JSON(e.SUCCESS, gin.H{
		"data": tag,
	})
}

func (t *TagViews) List(c *gin.Context) {
	var tags []models.Tag
	fmt.Println(c.Get("username"))
	fmt.Println(c.Get("password"))

	res := models.DB.Model(models.Tag{}).Find(&tags)
	if res.Error != nil {
		c.JSON(e.SERVERERROR, gin.H{
			"message": e.MsgFlags[e.SERVERERROR],
		})
	}
	c.JSON(e.SUCCESS, gin.H{
		"data": tags,
	})
}
