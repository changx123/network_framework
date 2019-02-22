package config

//MYSQL配置
const (
	//mysql 连接cdn地址
	MYSQL_DATA_SOURCE_NAME = "root:nihaoma123?@tcp(192.168.3.235:3306)/blog?charset=utf8&parseTime=true"
	//连接池最大链接数
	MYSQL_SET_MAX_OPEN_CONNS = 400
	//最大闲置链接数
	MYSQL_SET_MAX_IDLE_CONNS = 200
	//数据表前缀
	MYSQL_TABLE_PREFIX = "b_"
	//关闭表名复数例如"type table struct"表名为 "tables" 关闭 为 "table"
	MYSQL_TABLE_SINGULAR = true
)