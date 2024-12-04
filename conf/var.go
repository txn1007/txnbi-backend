package conf

type MySQLConfig struct {
	User      string
	Password  string
	Host      string
	Port      int
	DBName    string
	Charset   string
	ParseTime string
	Location  string
}

type TxnBIConfig struct {
	Host string
	Port int
}

type JWTConfig struct {
	SignKey string
}

var (
	// MySQL 配置全局变量
	MySQLCfg *MySQLConfig
	// 项目启动配置全局变量
	TxnBICfg *TxnBIConfig
	// JWT 配置全局变量
	JWTCfg *JWTConfig
)
