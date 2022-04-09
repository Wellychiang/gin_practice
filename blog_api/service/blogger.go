package service

import (
	"api/db"
	"api/entity"

	"gorm.io/gorm"
)

type Blogger entity.Blogger

func (Blogger) TableName() string {
	return "blogger"
}

func (blogger *Blogger) SearchUser() (b *Blogger) {
	b = new(Blogger)
	db.Db.Where("username = ? ", blogger.Username).First(b)
	return
}

func (blogger *Blogger) Register() *gorm.DB {
	return db.Db.Table("blogger").Create(blogger).Omit("id")
}

// 查找blogger
func (blogger *Blogger) Find() (b *Blogger) {
	b = new(Blogger)
	db.Db.Where("id = 1").First(b)
	return
}

func (blogger *Blogger) Insert() *gorm.DB {
	return db.Db.Create(blogger)
}

func (blogger *Blogger) UpdateInfo() *gorm.DB {
	// if blogger.Password != "" {
	// 	// db.Db.Table("blogger").Where("id = ?", blogger.Id).Update("id", blogger)
	// 	return db.Db.Save(blogger).Omit("password")
	// }
	// return db.Db.Model(blogger).Where("username = ?", blogger.Username).Omit("password").
	// 	Updates(blogger)
	return db.Db.Model(blogger).Where("username = ?", blogger.Username).Omit("username").Updates(blogger)

	// Save 會把所有欄位沒有值的清空, Updates 不會
	// return db.Db.Save(blogger)
}

func (blogger *Blogger) UpdatePassword() *gorm.DB {
	return db.Db.Model(blogger).Where("username = ?", blogger.Username).Update("password", blogger.Password)
}
