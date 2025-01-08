package postgres_db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func PostgresMain() {
	urlExample := "postgres://postgres:postgres@localhost:5432/postgres_db"

	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	queries := New(db)
	ctx := context.Background()

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(authors)

	author, err := queries.CreateAuthor(ctx, CreateAuthorParams{
		Name: "Franklin",
		Bio:  pgtype.Text{String: "Dev", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(author)

}
