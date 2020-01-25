package dbTransaction

import (
	"log"
	"time"
	
	"AppServer/src/common"
	"github.com/garyburd/redigo/redis"
)

//使用redis储存校验码，检测用户登录状态以及是否过期

var (
	redisCon            redis.Conn
	allowRedisPanic     bool = true
)

func TurnRedisPanic() {
	allowRedisPanic = !allowRedisPanic
	common.IF(allowSqlPanic, "panic of redis operation was allowed",
		"panic of redis operation was forbidden")
}

func panicOrLog(err error, passInfo ...string) {
	if err != nil {
		if allowRedisPanic {
			panic(err)
		} else {
			log.Println(err)
		}
	} else {
		for _, info := range passInfo {
			log.Println(info)
		}
	}
}

func InitDatabaseRedis() {
	var err error
	redisCon, err = redis.Dial("tcp", "127.0.0.1:6379")
	panicOrLog(err, "Connect to redis db succeed.")
}

func DisconnectDatabaseRedis() {
	if redisCon != nil {
		err := redisCon.Close()
		panicOrLog(err)
	}
}

func setUserToken(userId string) {
	_, err := redisCon.Do("SET", userId, "TestToken")
	panicOrLog(err)
}

func getUserToken(userId string) string {
	token, err := redis.String(redisCon.Do("GET", userId))
	panicOrLog(err)
	return token
}

func CheckUserToken(userId string) int {
	now := time.Now()
	former, err := time.Parse(
		common.AesDecrypt(userId, getUserToken(userId)),
		now.String(),
	)
	panicOrLog(err)
	check := now.Sub(former)
	if check > 600.0 {
		return -1
	} else {
		setUserToken(
			common.AesDecrypt(userId, getUserToken(userId)),
		)
		return 1
	}
}