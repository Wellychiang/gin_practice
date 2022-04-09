package service

import (
	"api/db"
	"api/entity"
	"api/utils"
	"fmt"

	"gorm.io/gorm"
)

type Blog entity.Blog

func (Blog) TableName() string {
	return "blog"
}

func (blog *Blog) FindBlogContent() (b *Blog) {
	b = new(Blog)

	fmt.Println(blog.Id)
	db.Db.Table("blog b").Select("b.*, bt.id as type_name").
		Joins("left join blogType bt on b.typeid = bt.id").
		Where("b.id = ?", blog.Id).Order("bt.sort asc").Find(b)

	fmt.Println(Blog{})
	fmt.Println(*blog)
	fmt.Println(b)
	fmt.Println(b.Id)
	fmt.Println(&b)
	fmt.Println(*b)

	return
}

// func (blog *Blog) FindBlogContent() Blog {
// 	var b Blog
// 	db.Db.Table("blog b").Select("b.*, bt.id as type_name").
// 		Joins("left join blogType bt on b.typeid = bt.id").
// 		Where("b.id = ?", blog.Id).Order("bt.sort asc").Find(&b)

// 	fmt.Println(b)
// 	fmt.Println(&b)
// 	return b
// }

func (blog *Blog) FindCommentByBlog() []Comment {
	comments := make([]Comment, 0)
	result := db.Db.Table("comment").
		Where("blogid = ? and status = 1", blog.Id).
		Order("addtime asc").
		Find(&comments)

	if result.Error != nil {
		return nil
	}

	return comments
}

func (blog *Blog) FindNextOne() (b *Blog) {
	b = new(Blog)
	result := db.Db.Where("id > ?", blog.Id).First(b)
	if result.Error != nil {
		return nil
	}
	return
}

func (blog *Blog) FindPreviousOne() (b *Blog) {
	b = new(Blog)
	result := db.Db.Where("id < ?", blog.Id).Order("id desc").First(b)
	// result := db.Db.Where("id < ?", blog.Id).First(b) = 1
	// ex. blod.Id = 4, where id < 4 ---> 1,2,3 這三個的第一個就會選到1, Order by id desc 降序
	// 變成 3, 2, 1 第一個也就是 3(也就是4的上一個)
	if result.Error != nil {
		return nil
	}
	return
}

// 查詢 blogger 列表
func (blog *Blog) FindList(page *utils.Page) ([]*Blog, error) {
	bs := make([]*Blog, 0)
	curDb := db.Db.Table("blog b").Select("b.*, bt.name as type_name").
		Joins("left join blogtype bt on b.id = bt.id")

	fmt.Print(curDb)
	if blog.TypeId > 0 {
		curDb = curDb.Where("b.typeid = ? ", blog.TypeId)
	}

	// Limit: 指定要查詢的最大紀錄數, Offset: 指定開始返回記錄前要跳過的紀錄數
	result := curDb.Limit(page.Size).Offset(page.GetsStart()).Order("addtime asc").Find(&bs)
	return bs, result.Error
}

func (blog *Blog) Count() (count int64) {
	db.Db.Model(blog).Count(&count)
	return
}

func (blog *Blog) UpdateClick() *gorm.DB {
	return db.Db.Model(blog).Where("id = ?", blog.Id).Update("clickhit", gorm.Expr("clickhit + ?", 1))

}

func (blog *Blog) UpdateReplay() *gorm.DB {
	return db.Db.Model(blog).Where("id = ?", blog.Id).Update("replayhit", gorm.Expr("replayhit + ?", 1))

}

func (blog *Blog) FindById() int64 {
	var count int64
	db.Db.Model(blog).Where("typeid = ?", blog.Id).Count(&count)
	return count
}
