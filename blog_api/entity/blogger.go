package entity

type Blogger struct {
	Id       int    `gorm:"column:id"json:"id"`
	Username string `gorm:"column:username"json:"username"` //用戶名
	Password string `gorm:"column:password"json:"password"` //密碼
	Nickname string `gorm:"column:nickname"json:"nickname"` //暱稱
	Sign     string `gorm:"column:sign"json:"sign"`         //個性簽名
	Profile  string `gorm:"column:profile"json:"profile"`   //個人簡介
	Img      string `gorm:"column:img"json:"img"`           //個人頭像
}
