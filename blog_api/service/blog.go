package service

import (
	"api/db"
	"api/entity"
	"api/utils"
)

type Blog entity.Blog

func (Blog) TableName() string {
	return "blog"
}

// 查詢 blogger 列表
func (blog *Blog) FindList(page *utils.Page) ([]*Blog, error) {
	bs := make([]*Blog, 0)
	curDb := db.Db.Table("blog b").Select("b.*, bt.name as type_name").
		Joins("left join blog_type bt on b.id = bt.id")

	if blog.TypeId > 0 {
		curDb = curDb.Where("b.typeId = ? ", blog.TypeId)
	}

	// Limit: 指定要查詢的最大紀錄數, Offset: 指定開始返回記錄前要跳過的紀錄數
	result := curDb.Limit(page.Size).Offset(page.GetsStart()).Order("`addtime asc`").Find(&bs)
	return bs, result.Error
}

func (blog *Blog) Count() (count int64) {
	db.Db.Model(blog).Count(&count)
	return
}
