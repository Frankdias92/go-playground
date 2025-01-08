package database

import (
	"myFirstProject/database/redis"
)

func DatabaseFoo() {

	// mysql.DatabaseMySQL()
	// postres.DatabasePostgres()
	redis.RedisDatabase()
}
