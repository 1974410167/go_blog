package views

import (
	"go_blog/e"
	"go_blog/models"
	"go_blog/serializers"
	"log"

	"github.com/gin-gonic/gin"
)

type PostViews struct {
}

func (p *PostViews) Retrieve(c *gin.Context) {
	strId := c.Param("id")
	id, err := serializers.PostInsSer.ConvertType(strId)
	if err != nil {
		log.Fatalf("%v", err)
	}
	postIns := models.GetPostByID(id)
	if postIns == nil {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"message": e.MsgFlags[e.ID_IS_NOT_EXIST],
		})
	} else {
		c.JSON(e.SUCCESS, gin.H{
			"data": postIns,
		})
	}
}
