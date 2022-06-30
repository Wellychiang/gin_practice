package service

import (
	"api/db"
	"api/entity"
	"api/utils"
	"fmt"

	"gorm.io/gorm"
)

type Blog entity.Blog

type BlogContent struct {
	// 要乘載撈出來的資料就需要一個像這個完整的 struct, 連 gorm tag 都要有
	Id         int    `gorm:"column:id"json:"id"`
	Title      string `gorm:"column:title"json:"title"`
	TypeId     int    `gorm:"column:typeid"json:"typeid"` //關聯 blogType id
	Content    string `gorm:"column:content"json:"content"`
	Summary    string `gorm:"column:summary"json:"summary"`
	ClickHit   int    `gorm:"column:clickhit"json:"clickhit"`
	ReplayHit  int    `gorm:"column:replayhit"json:"replayhit"`
	AddTime    string `gorm:"column:addtime"json:"addtime"`
	UpdateTime string `gorm:"column:updatetime"json:"updatetime"`
	TypeName   string `gorm:"-"json:"typename"` // - 等於可以忽略這字段沒關係
	BloggerId  string `gorm:"column:bloggerid"json:"bloggerid"`
	Blogger    string `gorm:"column:blogger"json:"blogger"`
}

type BlogComment struct {
	// 要乘載撈出來的資料就需要一個像這個完整的 struct, 連 gorm tag 都要有
	Id        int    `gorm:"column:id"json:"id"`
	Ip        string `gorm:"column:ip"json:"ip"`
	Content   string `gorm:"column:content"json:"content"`
	BlogId    int    `gorm:"column:blogid"json:"blogid"`
	Status    int    `gorm:"column:status"json:"status"`
	AddTime   string `gorm:"column:addtime"json:"addtime"`
	BlogTitle string `gorm:"column:blogtitle"json:"blogtitle"`
	BloggerId int    `gorm:"column:bloggerid"json:"bloggerid"`
	NickName  string `gorm:"column:nickname"json:"nickname"`
}

func (Blog) TableName() string {
	return "blog"
}

func (blog *Blog) FindBlogContent() (b *BlogContent) {

	b = new(BlogContent)

	// fmt.Println(blog.Id)
	db.Db.Table("blog b").Select("b.*, bt.id as type_name, ber.username as blogger").
		Joins("left join blogType bt on b.typeid = bt.id").
		Joins("left join blogger ber on b.bloggerid = ber.id").
		Where("b.id = ?", blog.Id).Order("bt.sort asc").Find(b)

	fmt.Printf("%+v", b)
	// fmt.Println(Blog{})
	// fmt.Println(*blog)
	// fmt.Println(b)
	// fmt.Println(b.Id)
	// fmt.Println(&b)
	// fmt.Println(*b)

	return
}

func (blog *Blog) FindCommentByBlog() []BlogComment {

	c := make([]BlogComment, 0)

	db.Db.Table("comment c").Select("c.*, ber.nickname as nickname").
		Joins("left join blogger ber on c.bloggerid = ber.id").
		Where("c.blogid = ? and c.status = 1", blog.Id).
		Order("addtime asc").Find(&c)

	fmt.Println(c)

	return c
}

// func (blog *Blog) FindCommentByBlog() []Comment {
// 	comments := make([]Comment, 0)
// 	result := db.Db.Table("comment").
// 		Where("blogid = ? and status = 1", blog.Id).
// 		Order("addtime asc").
// 		Find(&comments)

// 	if result.Error != nil {
// 		return nil
// 	}

// 	return comments
// }

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

func (blog *Blog) Insert() *gorm.DB {
	return db.Db.Create(blog).Omit("id")
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
