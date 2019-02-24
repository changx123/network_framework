package models

import (
	"blog/extend/mysql"
	"fmt"
	"time"
)

func init() {
	db , err := mysql.GetMysqlDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !db.HasTable(&Users{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Users{})
	}
	if !db.HasTable(&Articles{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Articles{})
	}
	if !db.HasTable(&Category{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Category{})
	}
	if !db.HasTable(&Tag{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Tag{})
	}
	if !db.HasTable(&ArticlesCategory{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ArticlesCategory{})
	}
	if !db.HasTable(&ArticlesTag{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ArticlesTag{})
	}
}

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}