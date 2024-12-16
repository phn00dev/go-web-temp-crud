package models

import (
	"fmt"

	"github.com/phn00dev/go-web-temp-crud/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	Title    string
	Slug     string
	Desc     string
	ImageUrl string
	Status   string
	gorm.Model
}

func (post Post) Migrate() {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&post)
}

func (post Post) GetPost(where ...interface{}) Post {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}
	db.First(&post, where...)
	return post
}

func (post Post) GetAllPost(where ...interface{}) []Post {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts []Post
	db.Find(&posts, where...)
	return posts
}

func (post Post) Create() {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&post)
}

func (post Post) Updates(data Post) {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Model(&post).Updates(data)
}

func (post Post) Delete(where ...interface{}) {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Unscoped().Delete(&post, where...)
}
