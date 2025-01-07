package postres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DatabasePostgres() {
	urlExample := "postgres://postgres:postgres@localhost:5432/postgres_db"
	db, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	if err := db.Ping(context.Background()); err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		return
	}

	// Inserção de dados
	query := "CREATE TABLE foo (id bigserial primary key, bar varchar(255))"
	if _, err := db.Exec(context.Background(), query); err != nil {
		fmt.Printf("Error creating data: %v\n", err)
		return
	}

	query = "INSERT INTO FOO (bar) VALUES ($1);"
	if _, err := db.Exec(context.Background(), query, "testPostgres"); err != nil {
		fmt.Printf("Error insert data: %v\n", err)
		return
	}

	query = "SELECT * FROM foo LIMIT 1"
	type foobar struct {
		ID  int64
		Bar string
	}
	var res foobar
	if err := db.QueryRow(context.Background(), query).Scan(&res.ID, &res.Bar); err != nil {
		fmt.Printf("Error querying data: %v\n", err)
		return
	}

	// Exibe os resultados
	fmt.Printf("Result: ID=%d, Bar=%s\n", res.ID, res.Bar)
}
