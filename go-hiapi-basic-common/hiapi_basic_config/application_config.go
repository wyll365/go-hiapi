package hiapi_basic_config

type ApplicationConfig struct {
	Debug       bool      `mapstructure:"debug"`
	Port        int32     `mapstructure:"port"`
	Environment string    `mapstructure:"environment"`
	Mysql       MysqlConf `mapstructure:"mysql"`
	Redis       RedisConf `mapstructure:"redis"`
}
