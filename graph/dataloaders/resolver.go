package dataloaders

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// dataloader function

func newGetProductID(ctx context.Context, store db.Store) *GetProductLoader {
	return NewGetProductLoader(GetProductConfig{
		Wait: 10 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(productID []int64) ([]*model.Product, []error) {
			p, err := store.GetProduct(ctx, int64(productID[0]))
			if err != nil {
				return nil, []error{ fmt.Errorf("get product dataloader error: %v", err) }
			}

			id := strconv.Itoa(int(p.ID))
			category := strconv.Itoa(int(p.Category))
			convertProc := &model.Product{
					ID: id,
					ProductName: p.ProductName,
					Description: p.Description.String,
					Img: p.Img.String,
					UnitPrice: int(p.UnitPrice),
					Discount: p.Discount,
					Stock: int(p.Stock),
					Brand: int(p.BrandID),
					Category: category,
					CreatedAt: p.CreatedAt,
					UpdatedAt: p.UpdatedAt,
			}
			return []*model.Product{ convertProc }, nil
		},
	})
}


func newGetAllProductsID(ctx context.Context, store db.Store) *GetAllProductsLoader {
	return NewGetAllProductsLoader(GetAllProductsConfig{
		Wait: 100 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(first []int64) ([][]*model.Product, []error) {
			args := db.ListProductParams {
				Limit: 30,
				Offset: int32(first[0]),
			}

			products, err := store.ListProduct(ctx, args)
			if err != nil {
				return nil, []error{ fmt.Errorf("could not get product all list : %v", err) }
			}

			convertProc := make(map[int64][]*model.Product, len(first))
			for _, p := range products {
				id := strconv.Itoa(int(p.ID))
				category := strconv.Itoa(int(p.Category))
				convertProc[first[0]] = append(convertProc[first[0]], &model.Product{
					ID: id,
					ProductName: p.ProductName,
					Description: p.Description.String,
					Img: p.Img.String,
					UnitPrice: int(p.UnitPrice),
					Discount: p.Discount,
					Stock: int(p.Stock),
					Brand: int(p.BrandID),
					Category: category,
					CreatedAt: p.CreatedAt,
					UpdatedAt: p.UpdatedAt,
				})
			}

			result := make([][]*model.Product, len(first))
			for i, f := range first {
				result[i] = convertProc[f]
			}

			return result, nil
		},
	})
}