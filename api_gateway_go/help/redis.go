package help

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func Get_redis() *redis.Client {

	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	Addr := host + ":" + port
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return client
}
