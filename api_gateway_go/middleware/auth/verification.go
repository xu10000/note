package auth

import (
	"api_gateway/help"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func is_white_url(url string, method string) bool {
	for _, white_url := range White_list {
		//		如果url地址不匹配
		if bys := white_url.Reg.Find([]byte(url)); len(bys) == 0 {
			fmt.Println("xx ", white_url, "ccc ", url)
			continue
		}
		// 如果方法相同
		for _, inner_method := range white_url.Methods {
			if inner_method == method {
				return true
			}
		}
	}

	return false
}

func Jwt_auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 白名单通过
		if ok := is_white_url(c.Request.URL.String(), c.Request.Method); ok {
			c.Next()
			return
		}

		// token签名验证
		_tokens, exists := c.Request.Header["Authorization"]

		if !exists {
			c.JSON(200, gin.H{
				"success": false,
				"message": "header不包含Authorization",
			})
			c.Abort()
			return
		}

		token_string := _tokens[0]

		token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(Jwt_secret), nil
		})

		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"message": "token验证失败：" + err.Error(),
			})
			c.Abort()
			return
		}

		if token.Valid != true {

			c.JSON(200, gin.H{
				"success": false,
				"message": "token验证失败，valid： false",
			})
			c.Abort()
			return
		}
		// 验证redis中是否存在该token
		claims, _ := token.Claims.(jwt.MapClaims)

		uid := int(claims["uid"].(float64))
		redis := help.Get_redis()

		val, err := redis.Get("gatewayuid" + strconv.Itoa(uid)).Result()
		if err != nil || val != token_string {
			c.JSON(200, gin.H{
				"success": false,
				"message": "redis不存在该token,可能过期了",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
