package service

import (
	"api/db"
	"api/entity"
)

type BlogType entity.BlogType

func (BlogType) TableName() string {
	return "blog_type"
}

func (blogtype *BlogType) FindTypeCount() []map[string]interface{} {
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
