package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Todo struct {
	ent.Schema
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "todos"},
	}
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Bool("done").Default(false),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("todos").Required().Unique(),
	}
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimestampsMixin{},
	}
}
