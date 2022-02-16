package service

import (
	"api/db"
	"api/entity"
)

type Blogger entity.Blogger

func (Blogger) TableName() string {
	return "blogger"
}

func (blogger *Blogger) Login() (b *Blogger) {
	b = new(Blogger)
	db.Db.Where("username = ? ", blogger.Username).First(b)
	return
}

// 查找blogger
func (blogger *Blogger) Find() (b *Blogger) {
	b = new(Blogger)
	db.Db.Where("id = 1").First(b)
	return
}
