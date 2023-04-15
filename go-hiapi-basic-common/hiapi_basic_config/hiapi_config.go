package hiapi_basic_config

type BasicConf struct {
	Debug       bool   `mapstructure:"debug"`
	Port        int32  `mapstructure:"port"`
	Environment string `mapstructure:"environment"`
}
