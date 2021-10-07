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
	viper.BindEnv("GIN_MODE")
	viper.AddConfigPath("./config")
	setConfigName()
	viper.ReadInConfig()
	unmarshal()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		unmarshal()
	})

}

// 根据环境变量读取配置文件
func setConfigName() {
	ginMode := viper.Get("GIN_MODE")
	if ginMode == ENV_PROD {
		fmt.Println("gin_mode:release")
		viper.SetConfigName("release.config")
	} else if ginMode == ENV_SLS {
		viper.SetConfigName("serverless.config")
	} else {
		viper.SetConfigName("debug.config")
		fmt.Println("gin_mode:debug")
	}
}

var config SettingConfig = SettingConfig{
	Dynamo: Dynamo{Region: "ap-southeast-1", PostTableName: "YQ_POST_TABLE"},
}

func unmarshal() {
	viper.Unmarshal(&config)
}

func GetConfig() *SettingConfig {
	return &config
}
