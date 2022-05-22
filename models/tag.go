package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string  `json:"name"`
	Posts []*Post `gorm:"many2many:post_tag;"`
}
