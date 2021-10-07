package boot

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
type Mysql struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Passwd       string `mapstructure:"passwd"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type SettingConfig struct {
	App   `mapstructure:"app"`
	Mysql `mapstructure:"mysql"`
}

func init() {
	// viper.SetConfigFile("config.yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	viper.ReadInConfig()
	unmarshal()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		unmarshal()
	})
}

var Config SettingConfig

func unmarshal() {
	viper.Unmarshal(&Config)
}
