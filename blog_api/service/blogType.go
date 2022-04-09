package service

import (
	"api/db"
	"api/entity"
	"api/utils"

	"gorm.io/gorm"
)

type BlogType entity.BlogType

func (BlogType) TableName() string {
	return "blogtype"
}

func (blogType *BlogType) FindTypeCount() []map[string]interface{} {
	maps := make([]map[string]interface{}, 0)
	sql := "SELECT b.typeid AS typeId,COUNT(b.id) AS b_count,bt.name AS b_name " +
		"FROM blog b LEFT JOIN blogType bt " +
		"On b.typeid = bt.id " +
		"GROUP BY b.typeid, bt.name"
	rows, _ := db.Db.Raw(sql).Rows()
	defer rows.Close()
	for rows.Next() {
		var typeId int
		var bCount int
		var bName string

		rows.Scan(&typeId, &bCount, &bName)
		Map := make(map[string]interface{})
		Map["type_id"] = typeId
		Map["b_count"] = bCount
		Map["b_name"] = bName
		maps = append(maps, Map)
	}
	return maps
}

func (blogType *BlogType) FindList(page *utils.Page) []*BlogType {
	blogTypes := make([]*BlogType, 0)
	db.Db.Model(blogType).Limit(page.Size).Offset(page.GetsStart()).Order("sort desc").Find(&blogTypes)
	return blogTypes
}

func (blogType *BlogType) Count() (count int64) {
	db.Db.Model(blogType).Count(&count)
	return
}

// func (blogType *BlogType) Count() int64 {
// 	var count int64
// 	db.Db.Model(blogType).Count(&count)
// 	return count
// }

func (blogType *BlogType) Insert() *gorm.DB {
	return db.Db.Model(blogType).Create(blogType)
}

func (blogType *BlogType) Delete() *gorm.DB {
	return db.Db.Model(blogType).Delete(blogType)
}

func (blogType *BlogType) Update() *gorm.DB {
	return db.Db.Model(blogType).Updates(blogType)
}
