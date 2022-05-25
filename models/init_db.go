package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:wsghy.5637@tcp(gz-cynosdbmysql-grp-0ku3m9v9.sql.tencentcdb.com:29978)/go_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("启动...............................")
	db.AutoMigrate(&Post{}, &Category{}, &Tag{}, &User{})
	DB = db
}
