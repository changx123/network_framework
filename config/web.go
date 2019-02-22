package config

//WEB配置
const (
	//调试模式
	WEB_DEBUG = true
	//开启http
	HTTP_OPEN = true
	//监听地址
	HTTP_LISTEN_ADDRESS = "0.0.0.0:1234"
	//启用http 热更新
	HTTP_HOT_UPDATE = true
	//热更新处理剩余连接超时时间
	HTTP_HOT_UPDATE_TIMEOUT = 20
	//开启https
	HTTPS_OPEN = false
	//监听地址
	HTTPS_LISTEN_ADDRESS = "0.0.0.0:443"
	//证书cer 文件路径
	HTTPS_CERTFILE_PATH = ""
	//证书key 文件路径
	HTTPS_KEYFILE_PATH = ""
)