package entity

type Comment struct {
	Id        int    `gorm:"column:id"json:"id"`
	Ip        string `gorm:"column:ip"json:"ip"`
	Content   string `gorm:"column:content"json:"content"`
	BlogId    int    `gorm:"column:blogid"json:"blogid"`
	Status    int    `gorm:"column:status"json:"status"`
	AddTime   string `gorm:"column:addtime"json:"addtime"`
	BlogTitle string `gorm:"column:blogtitle"json:"blogtitle"`
}
