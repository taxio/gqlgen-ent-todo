package main

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/dialect/sql"
	"github.com/taxio/gqlgen-ent-todo/ent"
	"github.com/taxio/gqlgen-ent-todo/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	entClient := ent.NewClient(ent.Driver(db))
	defer entClient.Close()

	if err := entClient.Debug().Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		return err
	}

	return nil
}
