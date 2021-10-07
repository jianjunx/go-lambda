package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type App struct {
	Name    string `mapstructure:"name"`
	Env     string `mapstructure:"env"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}
type Dynamo struct {
	Region        string `mapstructure:"region"`
	PostTableName string `mapstructure:"post_table_name"`
}

type SettingConfig struct {
	App    `mapstructure:"app"`
	Dynamo `mapstructure:"dynamo"`
}

func init() {
	// viper.SetConfigFile("config.yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	unmarshal()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		unmarshal()
	})
}

var config SettingConfig = SettingConfig{
	Dynamo: Dynamo{Region: "ap-northeast-1", PostTableName: "YQ_POST_TABLE"},
}

func unmarshal() {
	viper.Unmarshal(&config)
}

func GetConfig() *SettingConfig {
	return &config
}
