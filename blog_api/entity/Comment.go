package entity

type Comment struct {
	Id        uint64 `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Ip        string `gorm:"type:varchar(20)"json:"ip"`
	Content   string `gorm:"type:varchar(100)"json:"content"`
	BlogId    int    `gorm:"column:blogid"json:"blogid"`
	Status    int    `gorm:"column:status"json:"status"`
	AddTime   string `gorm:"column:addtime"json:"addtime"`
	BlogTitle string `gorm:"type:varchar(20);column:blogtitle"json:"blogtitle"`
	BloggerId int    `gorm:"column:bloggerid"json:"bloggerid"`
}
