package entity

type BlogType struct {
	Id   int    `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Name string `gorm:"column:name"json:"name"`
	Sort int    `gorm:"column:sort"json:"sort"`
}
