package proxy

import (
	"api_gateway/help"
	"fmt"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func Proxy_url() gin.HandlerFunc {

	return func(c *gin.Context) {
		reg, _ := regexp.Compile("^/\\S*?($|/)")

		service_key := reg.FindString(c.Request.URL.String())

		// 如果后面有/，删除后再进行匹配
		reg2, _ := regexp.Compile("/$")
		service_key = reg2.ReplaceAllString(service_key, "")

		service_name, exists := help.Config.Proxy[service_key]
		// 如果配置不存在该服务
		if exists != true {
			fmt.Println("配置中未找到该路径的服务：", service_key)
			c.Next()
			return
		}
		target, err := help.Get_service_url(service_name)
		//		consul未找到该服务
		if err != nil {
			fmt.Println("代理时未找到consul服务：", err)
			c.Next()
			return
		}
		// 路径重写
		c.Request.URL.Path = strings.Replace(c.Request.URL.Path, service_key, "", 1)
		remote, _ := url.Parse("http://" + target)
		proxy := httputil.NewSingleHostReverseProxy(remote)
		// fmt.Printf("准备进行代理----------------%s:  ", remote)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
