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
	viper.BindEnv(ENV_NAME)
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
	ginMode := viper.Get(ENV_NAME)
	if ginMode == ENV_PROD {
		viper.SetConfigName("prod.config")
	} else if ginMode == ENV_SLS {
		viper.SetConfigName("sls.config")
	} else {
		viper.SetConfigName("dev.config")
	}
}

var config SettingConfig = SettingConfig{
	Dynamo: Dynamo{Region: "us-west-1", PostTableName: "YQ_POST_TABLE"},
}

func unmarshal() {
	viper.Unmarshal(&config)
}

func GetConfig() *SettingConfig {
	return &config
}
