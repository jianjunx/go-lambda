package config

import (
	"fmt"
	"os"

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
	BookTableName string `mapstructure:"book_table_name"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	Maxsize    int    `mapstructure:"maxsize"`
	Maxage     int    `mapstructure:"maxage"`
	Maxbackups int    `mapstructure:"maxbackups"`
}

type SettingConfig struct {
	App    `mapstructure:"app"`
	Dynamo `mapstructure:"dynamo"`
	Log    `mapstructure:"log"`
}

var config SettingConfig = SettingConfig{}

func init() {
	ok := setConfigName()
	if !ok {
		return
	}
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	unmarshal()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		unmarshal()
	})

}

// 根据环境变量读取配置文件
func setConfigName() bool {
	mode := os.Getenv(ENV_NAME)
	if mode == ENV_SLS {
		// lambda读取配置文件有些问题 这里直接赋值
		config.App.Port = 80
		config.App.Env = mode
		config.Dynamo.Region = os.Getenv("DB_REGION")
		config.Dynamo.PostTableName = os.Getenv("POST_TABLE_NAME")
		config.Dynamo.BookTableName = os.Getenv("BOOK_TABLE_NAME")
		return false
	} else if mode == ENV_PROD {
		viper.SetConfigName("prod.config")
	} else {
		viper.SetConfigName("dev.config")
	}
	return true
}

func unmarshal() {
	viper.Unmarshal(&config)
}

func GetConfig() *SettingConfig {
	return &config
}
