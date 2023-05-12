package config

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Profile string `mapstructure:"profile"`
	Version string `mapstructure:"version"`

	*LogConfig `mapstructure:"log"`
	*DbConfig  `mapstructure:"db"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type DbConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
