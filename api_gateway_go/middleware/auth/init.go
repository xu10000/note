package auth

import (
	"api_gateway/help"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type _White_url struct {
	Reg     *regexp.Regexp
	Methods []string
}

var (
	Jwt_secret string
	White_list []_White_url
)

func capital_methods(methods []string) []string {
	my_methods := []string{}

	for _, method := range methods {
		my_methods = append(my_methods, strings.ToUpper(method))
	}

	return my_methods
}

func init() {

	// 将config的中白名单转换成正则存储在变量
	for _, url := range help.Config.White_list {

		reg := help.Get_url_regexp(url.Path)
		// 将方法都转为大写字母
		methods := capital_methods(url.Methods)
		white_url := _White_url{reg, methods}
		White_list = append(White_list, white_url)
	}

	Jwt_secret = viper.GetString("jwt.serect")
}
