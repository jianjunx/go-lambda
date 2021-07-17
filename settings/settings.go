package settings

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

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Passwd   string `mapstructure:"passwd"`
	Poolsize int    `mapstructure:"poolsize"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	Maxsize    int    `mapstructure:"maxsize"`
	Maxage     int    `mapstructure:"maxage"`
	Maxbackups int    `mapstructure:"maxbackups"`
}

type SettingConfig struct {
	App   `mapstructure:"app"`
	Mysql `mapstructure:"mysql"`
	Redis `mapstructure:"redis"`
	Log   `mapstructure:"log"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	unmarshal()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		unmarshal()
	})
	return
}

var Config SettingConfig

func unmarshal() {
	viper.Unmarshal(&Config)
}
