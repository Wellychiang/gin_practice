package service

import (
	"api/db"
	"api/entity"

	"gorm.io/gorm"
)

type Comment entity.Comment

func (Comment) TableName() string {
	return "comment"
}

func (comment *Comment) UpdateStatus() *gorm.DB {
	// var t Blog
	return db.Db.Table("comment").Where("id = ?", comment.Id).Update("status", comment.Status)
	// return db.Db.Model(comment).Where("id = ?", comment.Id).Update("status", comment.Status)
	// Model 吃的是 function TableName 裡的 return string, 上面兩個 return 依樣意思
}

func (comment *Comment) Insert() *gorm.DB {
	return db.Db.Create(comment)
}
