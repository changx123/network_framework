package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"network_framework/config"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.MYSQL_TABLE_PREFIX + defaultTableName
	}
}

//获取数据库
func GetMysqlDb() (*gorm.DB, error) {
	var err error
	Db, err = gorm.Open("mysql", config.MYSQL_DATA_SOURCE_NAME)
	if err != nil {
		return nil, err
	}
	//最大链接数
	Db.DB().SetMaxOpenConns(config.MYSQL_SET_MAX_OPEN_CONNS)
	//最大闲置链接
	Db.DB().SetMaxIdleConns(config.MYSQL_SET_MAX_IDLE_CONNS)
	return Db, nil
}
