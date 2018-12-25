package auth

import (
	"api_gateway/help"
	"api_gateway/syslog"
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	resty "gopkg.in/resty.v1"
)

func token_string(result *help.Result) (string, error) {
	authorities := result.Data.(map[string]interface{})["roles"]
	f_uid := result.Data.(map[string]interface{})["id"].(float64)
	uid := int(f_uid)
	name := result.Data.(map[string]interface{})["account"]

	token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":        name,
		"uid":         uid,
		"authorities": authorities,
		"ctime":       time.Now(),
	})

	token, err := token_obj.SignedString([]byte(Jwt_secret))

	if err != nil {
		return "", err
	}

	return token, nil

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印所有的请求
		url := c.Request.URL
		method := c.Request.Method
		syslog.Debug.Printf("请求路径: %s  请求方法：%s \n", url, method)
		// 如果是不是登录接口，则结束
		if match, _ := regexp.MatchString("^/api/v1/login($|/$|\\?\\S*)", url.Path); !match {
			c.Next()
			return

		}

		if method != "POST" {
			c.Next()
			return
		}
		// 获取account服务
		account_url, err := help.Get_service_url("api_account")
		// 如果错误返回
		if err != nil {
			syslog.Debug.Printf("获取服务失败:%s ", err)
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		var _param interface{}
		c.BindJSON(&_param)
		//param, _ := json.Marshal(&_param)
		// 代理登录
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(_param).
			SetResult(&help.Result{false, "", ""}).
			Post("http://" + account_url + "/api/v1/login")

		// 登录失败返回相应的错误信息
		if err != nil {

			syslog.Debug.Printf("登录请求发生错误： %s", err)
			c.JSON(500, gin.H{
				"success": false,
				"message": "登录失败:" + err.Error(),
			})
			c.Abort()
			return
		}

		result := resp.Result().(*help.Result)
		if result.Success != true {
			err_str := result.Message.(string)
			syslog.Debug.Printf("登录请求发生错误： %s", err_str)
			c.JSON(401, gin.H{
				"success": false,
				"message": "录失败:" + err_str,
			})
			c.Abort()
			return
		}
		//替换token
		token, err := token_string(result)

		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": "替换token出错:" + err.Error(),
			})
			c.Abort()
			return
		}
		//登录成功后设置redis缓存
		redis := help.Get_redis()

		f_uid := result.Data.(map[string]interface{})["id"].(float64)
		uid := int(f_uid)
		expr := 24 * time.Hour
		err2 := redis.Set("gatewayuid"+strconv.Itoa(uid), token, expr).Err()
		if err2 != nil {
			panic(err2)
		}

		result.Data.(map[string]interface{})["token"] = token
		//登录成功返回正确信息
		c.JSON(200, gin.H{
			"success": true,
			"data":    result.Data,
		})

		c.Abort()
		return
	}
}
