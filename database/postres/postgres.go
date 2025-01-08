package postres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabasePostgres() {
	urlExample := "postgres://postgres:postgres@localhost:5432/postgres_db"

	// Inicializa o pool de conexões
	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Testa a conexão
	if err := db.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error pinging database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database successfully!")

	// Criação da tabela
	query := `CREATE TABLE IF NOT EXISTS foo (
		id BIGSERIAL PRIMARY KEY,
		bar VARCHAR(255)
	)`
	if _, err := db.Exec(context.Background(), query); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating table: %v\n", err)
		return
	}
	fmt.Println("Table created successfully!")

	// Inserção de dados
	query = "INSERT INTO foo (bar) VALUES ($1)"
	if _, err := db.Exec(context.Background(), query, "testPostgres"); err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting data: %v\n", err)
		return
	}
	fmt.Println("Data inserted successfully!")

	// Consulta de dados
	query = "SELECT id, bar FROM foo LIMIT 1"
	type foobar struct {
		ID  int64
		Bar string
	}
	var res foobar
	if err := db.QueryRow(context.Background(), query).Scan(&res.ID, &res.Bar); err != nil {
		fmt.Fprintf(os.Stderr, "Error querying data: %v\n", err)
		return
	}

	// Exibe os resultados
	fmt.Printf("Result: ID=%d, Bar=%s\n", res.ID, res.Bar)
}
