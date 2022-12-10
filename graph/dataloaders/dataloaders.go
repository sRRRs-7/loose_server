package dataloaders

import (
	"context"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
)

type ctxKey string
const Key = ctxKey("dataloaders")

// individual loader define
type Loaders struct {
	GetProductID *GetProductLoader
	GetAllProductsID *GetAllProductsLoader
}

func NewLoaders(ctx context.Context, store db.Store) *Loaders {
	return &Loaders {
		GetProductID: newGetProductID(ctx, store),
		GetAllProductsID: newGetAllProductsID(ctx, store),
	}
}

type Retriever struct {
	key ctxKey
}

func (r *Retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

func NewRetriever() *Retriever {
	return &Retriever{key: Key}
}