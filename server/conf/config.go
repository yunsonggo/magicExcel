package conf

type AppConfig struct {
	AppName      string   `mapstructure:"name"`
	AppMode      string   `mapstructure:"mode"`
	AppListen    string   `mapstructure:"listen"`
	Cors         []string `mapstructure:"cors"`
	AppLimit int `mapstructure:"limit"`
	AppLimitTime int `mapstructure:"limit_time"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*JwtConfig   `mapstructure:"jwt"`
}

type LogConfig struct {
	LogLevel      string `mapstructure:"level"`
	LogFileName   string `mapstructure:"filename"`
	LogMaxSize    int    `mapstructure:"max_size"`
	LogMaxAge     int    `mapstructure:"max_age"`
	LogMaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	MysqlAddr   string `mapstructure:"addr"`
	MysqlUser   string `mapstructure:"user"`
	MysqlPass   string `mapstructure:"pass"`
	MysqlDbName string `mapstructure:"dbname"`
}

type RedisConfig struct {
	RedisAddr string `mapstructure:"addr"`
	RedisPass string `mapstructure:"pass"`
	RedisDb   int    `mapstructure:"db"`
}
type JwtConfig struct {
	JwtKey string `mapstructure:"jwt_key"`
	Issuer string `mapstructure:"issuer"`
	Expire int    `mapstructure:"expire"`
}
