package mysql

import (
	"network_framework/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mysqlDb *gorm.DB

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.MYSQL_TABLE_PREFIX + defaultTableName
	}
}

//获取数据库
func GetMysqlDb() (*gorm.DB, error) {
	if mysqlDb == nil{
		var err error
		db, err := connect()
		if err != nil {
			return db, err
		}
		mysqlDb = db
	}
	return mysqlDb, nil
}

func connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.MYSQL_DATA_SOURCE_NAME)
	if err != nil {
		return nil, err
	}
	//最大链接数
	db.DB().SetMaxOpenConns(config.MYSQL_SET_MAX_OPEN_CONNS)
	//最大闲置链接
	db.DB().SetMaxIdleConns(config.MYSQL_SET_MAX_IDLE_CONNS)
	db.DB().Ping()
	//表名后缀
	db.SingularTable(config.MYSQL_TABLE_SINGULAR)
	return db, nil
}
