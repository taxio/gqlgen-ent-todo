package resolver

import "github.com/taxio/gqlgen-ent-todo/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *ent.Client
}

func NewResolver(db *ent.Client) *Resolver {
	return &Resolver{db}
}
