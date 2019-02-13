package config

//MYSQL配置
const (
	//mysql 连接cdn地址
	MYSQL_DATA_SOURCE_NAME = "root:nihaoma123?@tcp(server.cocoxb.cn:3306)/blog?charset=utf8"
	//连接池最大链接数
	MYSQL_SET_MAX_OPEN_CONNS = 30
	//最大闲置链接数
	MYSQL_SET_MAX_IDLE_CONNS = 15
	//数据表前缀
	MYSQL_TABLE_PREFIX = "b_"
)