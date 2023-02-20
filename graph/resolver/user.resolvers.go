package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/taxio/gqlgen-ent-todo/ent"
	qtodo "github.com/taxio/gqlgen-ent-todo/ent/todo"
	quser "github.com/taxio/gqlgen-ent-todo/ent/user"
	"github.com/taxio/gqlgen-ent-todo/graph"
	"github.com/taxio/gqlgen-ent-todo/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	newUser, err := r.db.User.Create().SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, err
	}

	nodeId, err := encodeNodeId("User", newUser.ID)
	if err != nil {
		return nil, err
	}

	return &model.CreateUserPayload{
		User: &model.User{
			ID:   nodeId,
			Name: newUser.Name,
		},
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	_, userID, err := parseNodeID(id)
	if err != nil {
		return nil, err
	}

	user, err := r.db.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	nodeId, err := encodeNodeId("User", user.ID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   nodeId,
		Name: user.Name,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, first *int, after *string) (*model.UserConnection, error) {
	query := r.db.User.Query()
	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}

	if first != nil {
		query = query.Limit(*first)
	}
	if after != nil {
		_, userID, err := parseNodeID(*after)
		if err != nil {
			return nil, err
		}
		query = query.Where(quser.IDGT(userID))
	}
	users, err := query.Order(ent.Asc(quser.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}

	edges := make([]*model.UserEdge, 0, len(users))
	for _, user := range users {
		nodeId, err := encodeNodeId("User", user.ID)
		if err != nil {
			return nil, err
		}
		edges = append(edges, &model.UserEdge{
			Node: &model.User{
				ID:   nodeId,
				Name: user.Name,
			},
			Cursor: nodeId,
		})
	}
	return &model.UserConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage: func(lenUsers int, first *int) bool {
				if first == nil {
					return false
				}
				return lenUsers == *first
			}(len(users), first),
			EndCursor: func(edges []*model.UserEdge) *string {
				if len(edges) == 0 {
					return nil
				}
				return &edges[len(edges)-1].Cursor
			}(edges),
		},
		TotalCount: totalCount,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User, first *int, after *string) (*model.TodoConnection, error) {
	if obj == nil {
		return nil, fmt.Errorf("parent user is nil")
	}
	_, userID, err := parseNodeID(obj.ID)
	if err != nil {
		return nil, err
	}
	query := r.db.Todo.Query().Where(qtodo.HasOwnerWith(quser.IDEQ(userID)))
	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	if first != nil {
		query = query.Limit(*first)
	}
	if after != nil {
		_, todoID, err := parseNodeID(*after)
		if err != nil {
			return nil, err
		}
		query = query.Where(qtodo.IDGT(todoID))
	}
	todos, err := query.Order(ent.Asc(qtodo.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}

	edges := make([]*model.TodoEdge, 0, len(todos))
	for _, todo := range todos {
		nodeId, err := encodeNodeId("Todo", todo.ID)
		if err != nil {
			return nil, err
		}
		edges = append(edges, &model.TodoEdge{
			Node: &model.Todo{
				ID:      nodeId,
				Content: todo.Content,
				Done:    todo.Done,
			},
			Cursor: nodeId,
		})
	}
	return &model.TodoConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage: func(lenTodos int, first *int) bool {
				if first == nil {
					return false
				}
				return lenTodos == *first
			}(len(todos), first),
			EndCursor: func(edges []*model.TodoEdge) *string {
				if len(edges) == 0 {
					return nil
				}
				return &edges[len(edges)-1].Cursor
			}(edges),
		},
		TotalCount: totalCount,
	}, nil
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
