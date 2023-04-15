package hiapi_basic_config

type ApplicationConfig struct {
	BasicConf
	Mysql MysqlConf `mapstructure:"mysql"`
	Redis RedisConf `mapstructure:"redis"`
}
