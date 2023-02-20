package loader

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"github.com/taxio/gqlgen-ent-todo/ent"
)

type ctxKey string

const loaderKey = ctxKey("dataloaders")

type Loaders struct {
	TodoOwnerLoader *dataloader.Loader
}

func NewLoaders(db *ent.Client) *Loaders {
	todoLoader := &TodoLoader{db: db}
	return &Loaders{
		TodoOwnerLoader: dataloader.NewBatchedLoader(todoLoader.BatchGetTodoOwners),
	}
}

func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	//loaders.UserLoader.ClearAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), loaderKey, loaders)))
	})
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loaderKey).(*Loaders)
}
