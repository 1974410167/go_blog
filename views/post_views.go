package views

import (
	"go_blog/e"
	"go_blog/models"
	"go_blog/serializers"

	"github.com/gin-gonic/gin"
)

type PostViews struct {
}

func (p *PostViews) Retrieve(c *gin.Context) {
	strId := c.Param("id")
	id, err := serializers.PostInsSer.ConvertType(strId)
	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.INVALID_PARAMS],
		})
	}
	var post *models.Post
	res := models.DB.Model(models.Post{}).Where("id = ?", id).First(post)
	if res.Error == nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.ID_IS_NOT_EXIST],
		})
	} else {
		c.JSON(e.SUCCESS, gin.H{
			"data": post,
		})
	}
}

func (p *PostViews) Create(c *gin.Context) {

	var post models.Post
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.INVALID_PARAMS],
		})
		return
	}
	res := models.DB.Model(models.Post{}).Create(post)
	if res.Error != nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.Create_ERROR],
		})
		return
	}
	c.JSON(e.SUCCESS, gin.H{
		"data": post,
	})
}
