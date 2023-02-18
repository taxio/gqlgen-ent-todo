//go:build tools
// +build tools

package tools

import (
	_ "entgo.io/ent"
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)

//go:generate go build -v -o=./bin/gqlgen github.com/99designs/gqlgen
//go:generate go build -v -o=./bin/ent entgo.io/ent/cmd/ent
