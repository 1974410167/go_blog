package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Excerpt    string   `json:"excerpt"`
	PageViews  string   `json:"page_views"`
	Tags       []*Tag   `gorm:"many2many:post_tag"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}

type PostTag struct {
	gorm.Model
	PostID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
}

func GetPostByID(id uint) *Post {
	var post *Post
	DB.Model(&Post{}).Where("id = ?", id).First(post)
	return post
}

func CreatePost(post *Post) (*Post, error) {
	if res := DB.Create(post); res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}
