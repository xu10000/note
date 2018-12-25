package help

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type _Config struct {
	White_list []struct {
		Path    string   `mapstructure:"path"`
		Methods []string `mapstructure:"methods"`
	} `mapstructure:"white_list"`

	Proxy map[string]string `mapstructure:"proxy"`
}

var (
	Config _Config
)

func init() {

	var my_env string
	// 加载配置文件
	flag.StringVar(&my_env, "id", "local", "the code of env")
	flag.Parse()

	config_name := "config-" + my_env // without extension
	viper.SetConfigName(config_name)
	viper.AddConfigPath("./config/") // optionally look for config in the working directory
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err := viper.Unmarshal(&Config)

	if err != nil {
		panic("viper解析config出错 " + err.Error())
	}

}
