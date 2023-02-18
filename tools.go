//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)

//go:generate go build -v -o=./bin/gqlgen github.com/99designs/gqlgen
