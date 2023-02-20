package main

import (
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"entgo.io/ent/dialect/sql"
	"github.com/jaswdr/faker"
	"github.com/taxio/gqlgen-ent-todo/ent"
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

	fake := faker.New()

	for i := 0; i < 10; i++ {
		user, err := entClient.User.Create().SetName(fake.Person().Name()).Save(ctx)
		if err != nil {
			return err
		}

		for j := 0; j < 10; j++ {
			_, err := entClient.Todo.Create().SetOwner(user).SetContent(fake.Lorem().Sentence(10)).SetDone(fake.Bool()).Save(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
