package loader

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/opentracing/opentracing-go/log"
	"github.com/taxio/gqlgen-ent-todo/ent"
	qtodo "github.com/taxio/gqlgen-ent-todo/ent/todo"
)

type TodoLoader struct {
	db *ent.Client
}

func (t *TodoLoader) BatchGetTodoOwners(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	todoIDs := make([]int, 0, len(keys))
	for _, key := range keys {
		todoID, err := strconv.Atoi(key.String())
		if err != nil {
			log.Error(err)
			return nil
		}
		todoIDs = append(todoIDs, todoID)
	}

	todos, err := t.db.Todo.Query().WithOwner().Where(qtodo.IDIn(todoIDs...)).All(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}
	todoByID := make(map[string]*ent.Todo, len(todos))
	for _, todo := range todos {
		todoByID[strconv.Itoa(todo.ID)] = todo
	}

	output := make([]*dataloader.Result, 0, len(keys))
	for _, key := range keys {
		todo, ok := todoByID[key.String()]
		if ok {
			output = append(output, &dataloader.Result{Data: todo.Edges.Owner})
		} else {
			output = append(output, &dataloader.Result{
				Data:  nil,
				Error: fmt.Errorf("todo not found %s", key.String()),
			})
		}
	}
	return output
}

func GetOwner(ctx context.Context, todoID int) (*ent.User, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.TodoOwnerLoader.Load(ctx, dataloader.StringKey(strconv.Itoa(todoID)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*ent.User), nil
}
