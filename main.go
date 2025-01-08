package main

import (
	"myFirstProject/database"
	postgres_db "myFirstProject/sql"
)

func main() {
	database.DatabaseFoo()
	postgres_db.PostgresMain()
}
