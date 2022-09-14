package entity

type Blog struct {
	Id         int    `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Title      string `gorm:"type:varchar(30)"json:"title"`
	TypeId     int    `gorm:"column:typeid"json:"typeid"` //關聯 blogType id
	Content    string `gorm:"type:varchar(100)"json:"content"`
	Summary    string `gorm:"type:varchar(30)"json:"summary"`
	ClickHit   int    `gorm:"column:clickhit"json:"clickhit"`
	ReplayHit  int    `gorm:"column:replayhit"json:"replayhit"`
	AddTime    string `gorm:"column:addtime"json:"addtime"`
	UpdateTime string `gorm:"column:updatetime"json:"updatetime"`
	TypeName   string `gorm:"-"json:"typename"` // - 等於可以忽略這字段沒關係
	BloggerId  int    `gorm:"column:bloggerid"json:"bloggerid"`
}
