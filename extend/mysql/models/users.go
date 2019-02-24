package models

type Users struct {
	//用户id
	Id uint `gorm:"primary_key;AUTO_INCREMENT"`
	//昵称
	Nickname string `gorm:"type:CHAR(30);not null;default:''"`
	//用户名
	Username string `gorm:"type:CHAR(30);unique_index;not null;default:''"`
	//密码 sha 256
	Passwd string `gorm:"type:VARCHAR(256);not null;default:''"`
	Model
}

