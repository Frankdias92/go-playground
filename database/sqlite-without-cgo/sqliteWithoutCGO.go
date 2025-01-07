package sqlitewithoutcgo

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func SqliteWithoutCGO() {
	db, err := sql.Open("sqlite", "./bar.foo")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

}
