package main

import (
	"log"

	"api_gateway/help"
	"api_gateway/middleware/auth"
	"api_gateway/middleware/proxy"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"github.com/spf13/viper"
)

type System struct{}

func (s *System) Hello(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "hello, api_gateway",
	})
}

func (s *System) Loginout(c *gin.Context) {

	redis := help.Get_redis()
	// 进入到该函数，已经确定有Authorization
	claims_map := help.Get_claims(c.Request)
	uid := int(claims_map["uid"].(float64))
	err := redis.Del("gatewayuid" + strconv.Itoa(uid)).Err()

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "退出登录失败，redis删除失败: " + err.Error(),
		})
		return
	}
	//Println("----------------- ", uid)
	c.JSON(200, gin.H{
		"success": true,
		"message": "退出登录成功",
	})

}

func main() {
	// consul注册参数
	my_registry := help.Get_registry()

	// 生成服务
	app_name := viper.GetString("gateway.name")
	app_host := viper.GetString("gateway.host")
	app_port := viper.GetString("gateway.port")
	app_url := app_host + ":" + app_port

	service := web.NewService(
		web.Name(app_name),
		web.Address(app_url),
		web.Registry(my_registry),
	)

	router := gin.Default()
	/*
	** 使用中间件
	 */
	// 代理登录
	router.Use(auth.Login())
	// 验证token
	router.Use(auth.Jwt_auth())
	// 使用代理
	router.Use(proxy.Proxy_url())
	// 注册路由，并有一个接口，供consul检查是否健康,另一个登出
	system := new(System)
	router.GET("/", system.Hello)
	router.POST("/api/v1/loginout", system.Loginout)
	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
