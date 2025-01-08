package database

import (
	"github.com/frankdias92/myFirstProject/database/mysql"
	"github.com/frankdias92/myFirstProject/database/postres"
	"github.com/frankdias92/myFirstProject/database/redis"
)

func DatabaseFoo() {

	mysql.DatabaseMySQL()
	postres.DatabasePostgres()
	redis.RedisDatabase()
}
