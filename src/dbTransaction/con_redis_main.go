package dbTransaction

import (
	"AppServer/src/common"
	"github.com/garyburd/redigo/redis"
)

var (
	redisCon            redis.Conn
	allowRedisPanic     bool = true
)

func TurnRedisPanic() {
	allowRedisPanic = !allowRedisPanic
	common.IF(allowSqlPanic, "panic of redis operation was allowed",
		"panic of redis operation was forbidden")
}

func InitRedis() {

}