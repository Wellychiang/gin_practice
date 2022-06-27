package entity

type Blog struct {
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
}
