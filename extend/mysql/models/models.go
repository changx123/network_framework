package models

import (
	"time"
)

func init() {
	//db , err := mysql.GetMysqlDb()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if !db.HasTable(&Users{}) {
	//	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Users{})
	//}
}

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}