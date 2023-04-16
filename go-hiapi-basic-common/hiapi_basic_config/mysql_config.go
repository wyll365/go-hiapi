package hiapi_basic_config

type MysqlConf struct {
	Url      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
}
