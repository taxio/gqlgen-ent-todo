package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	_ "github.com/go-sql-driver/mysql"
	"github.com/taxio/gqlgen-ent-todo/graph/loader"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/taxio/gqlgen-ent-todo/ent"
	"github.com/taxio/gqlgen-ent-todo/graph"
	"github.com/taxio/gqlgen-ent-todo/graph/resolver"
)

const defaultPort = "8080"

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	entClient := ent.NewClient(ent.Driver(db))
	defer entClient.Close()

	ldrs := loader.NewLoaders(entClient)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver.NewResolver(entClient)}))
	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", loader.Middleware(ldrs, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}
