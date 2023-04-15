package hiapi_basic_config

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
	Database int    `mapstructure:"database"`
	Port     int32  `mapstructure:"port"`
}
