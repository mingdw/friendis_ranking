package sys_common

import (
	"friends_ranking/utils/redis_factory"
	"math/rand"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetVerfyCode(key string) string {
	var code = ""
	for i := 0; i <= 5; i++ {

		code = code + string(letters[rand.Intn(64)])
	}
	redCli := redis_factory.GetOneRedisClient()
	redisKey := "verfyCode: " + key
	redCli.Execute("set", redisKey, code, "EX", 60)
	return code
}
