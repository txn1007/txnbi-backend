package conf

import "github.com/spf13/viper"

func init() {
	// Viper 初始化
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 初始化项目启动配置
	GetTxnBIConfig()
	// 初始化 MySQL
	GetMySQLConfig()
	// 初始化 JWT
	GetJWTConfig()
	// 初始化 Redis
	GetRedisConfig()
	// 初始化 Limiter 使用的 Redis
	GetLimiterRedisConfig()
}

func GetMySQLConfig() {
	MySQLCfg = new(MySQLConfig)
	MySQLCfg.User = viper.GetString("mysql.user")
	MySQLCfg.Password = viper.GetString("mysql.password")
	MySQLCfg.Host = viper.GetString("mysql.host")
	MySQLCfg.Port = viper.GetInt("mysql.port")
	MySQLCfg.DBName = viper.GetString("mysql.dbname")
	MySQLCfg.Charset = viper.GetString("mysql.charset")
	MySQLCfg.Location = viper.GetString("mysql.location")
	MySQLCfg.ParseTime = viper.GetString("mysql.parse_time")
}

func GetTxnBIConfig() {
	TxnBICfg = new(TxnBIConfig)
	TxnBICfg.Host = viper.GetString("txnbi.host")
	TxnBICfg.Port = viper.GetInt("txnbi.port")
}

func GetJWTConfig() {
	JWTCfg = new(JWTConfig)
	JWTCfg.SignKey = viper.GetString("jwt.sign_key")
}

func GetRedisConfig() {
	RedisCfg = new(RedisConfig)
	RedisCfg.Host = viper.GetString("redis.host")
	RedisCfg.Port = viper.GetInt("redis.port")
	RedisCfg.Password = viper.GetString("redis.password")
	RedisCfg.DB = viper.GetInt("redis.DB")
}

func GetLimiterRedisConfig() {
	LimiterRedisCfg = new(LimiterRedisConfig)
	LimiterRedisCfg.Host = viper.GetString("limiterRedis.host")
	LimiterRedisCfg.Port = viper.GetInt("limiterRedis.port")
	LimiterRedisCfg.Password = viper.GetString("limiterRedis.password")
	LimiterRedisCfg.DB = viper.GetInt("limiterRedis.DB")
}
