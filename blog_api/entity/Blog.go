package entity

type Blog struct {
	Id         int    `gorm:"column:id"json:"id"`
	Title      string `gorm:"column:title"json:"title"`
	TypeId     int    `gorm:"column:typeId"json:"typeId"` //關聯 blogType id
	Content    string `gorm:"column:content"json:"content"`
	Summary    string `gorm:"column:summary"json:"summary"`
	ClickHit   int    `gorm:"column:click_hit"json:"click_hit"`
	ReplayHit  int    `gorm:"column:replay_hit"json:"replay_hit"`
	AddTime    string `gorm:"column:add_time"json:"add_time"`
	UpdateTime string `gorm:"column:update_time"json:"update_time"`
	TypeName   string `gorm:"-"json:"type_name"` // - 等於可以忽略這字段沒關係
}
