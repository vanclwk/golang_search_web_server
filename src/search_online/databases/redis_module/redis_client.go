package redis_module

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go_search_online/src/search_online/logics/config_logics"
)

var GlobalRedis *redis.Client

func init() {
	fmt.Printf("Loading Redis module...\n")
	GlobalRedis = initRedis()
}

func initRedis() *redis.Client {
	globalConfig := config_logics.GlobalConfig
	redisOption := redis.Options{
		Addr:globalConfig.GetString("database.redis.host") + ":" + globalConfig.GetString("database.redis.port"),
		Password:globalConfig.GetString("database.redis.password"),
		DB:globalConfig.GetInt("database.redis.database"),
	}
	client := redis.NewClient(&redisOption)
	return client
}
