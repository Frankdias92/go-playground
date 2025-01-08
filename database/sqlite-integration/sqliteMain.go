package sqliteintegration

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func SqliteMain() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// SQL para criar a tabela
	createTableSql := `
	CREATE TABLE IF NOT EXISTS foo (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT
	);`

	_, err = db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table created.")

	// SQL para inserir um registro
	insertSql := `
	INSERT INTO foo (id, name) VALUES (?, ?);`

	_, err = db.Exec(insertSql, 1, "test Name")
	if err != nil {
		panic(err)
	}
	fmt.Println("Register created.")

	// Estrutura para mapear os dados retornados
	type User struct {
		ID   int64
		Name string
	}

	querySql := `
	SELECT id, name FROM foo WHERE id = ?;`

	var u User
	err = db.QueryRow(querySql, 1).Scan(&u.ID, &u.Name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User found: ID=%d, Name=%s\n", u.ID, u.Name)
}
