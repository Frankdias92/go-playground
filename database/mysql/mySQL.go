package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseMySQL() {
	// Configuração do DSN
	dsn := "root:root@tcp(0.0.0.0:3306)/mysql_db"
	// dsn := "root:root@/mysql_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	// Configurações de conexão
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Testa a conexão com o banco de dados
	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		return
	}

	// Criação da tabela
	query := "CREATE TABLE IF NOT EXISTS foo (id BIGINT AUTO_INCREMENT PRIMARY KEY, bar VARCHAR(255));"
	if _, err := db.Exec(query); err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		return
	}

	// Inserção de dados
	query = "INSERT INTO foo (bar) VALUES (?)"
	if _, err := db.Exec(query, "testMySql"); err != nil {
		fmt.Printf("Error inserting data: %v\n", err)
		return
	}

	// Consulta ao banco de dados
	query = "SELECT id, bar FROM foo LIMIT 1;"
	type Foobar struct {
		ID  int64
		Bar string
	}
	var res Foobar
	if err := db.QueryRow(query).Scan(&res.ID, &res.Bar); err != nil {
		fmt.Printf("Error querying data: %v\n", err)
		return
	}

	// Exibe os resultados
	fmt.Printf("Result: ID=%d, Bar=%s\n", res.ID, res.Bar)
}
