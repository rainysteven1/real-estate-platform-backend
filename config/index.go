package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	// Config is the global config
	Config *viper.Viper
)

var Conf = new(AppConfig)

func init() {
	Config = viper.New()
	Config.SetConfigName("application")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./resources/")
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		log.Printf("Error reading config, err:%v\n", err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		log.Printf("viper.Unmarshal failed, err:%v\n", err)
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s\n", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			log.Printf("viper.Unmarshal failed, err:%v\n", err)
			panic(err)
		}
	})
}
