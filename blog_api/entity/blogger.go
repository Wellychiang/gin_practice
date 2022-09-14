package entity

type Blogger struct {
	Id       int    `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Username string `gorm:"type:varchar(30)"json:"username"` //用戶名
	Password string `gorm:"type:varchar(33)"json:"password"` //密碼
	Nickname string `gorm:"type:varchar(15)"json:"nickname"` //暱稱
	Sign     string `gorm:"type:varchar(30)"json:"sign"`     //個性簽名
	Profile  string `gorm:"type:varchar(50)"json:"profile"`  //個人簡介
	Img      string `gorm:"type:bytea"json:"img"`            //個人頭像
}
