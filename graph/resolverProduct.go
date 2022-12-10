package graph

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

type SortBy struct {
	Asc  string
	Desc string
}

var EnumSort = SortBy{
	Asc:  "ASC",
	Desc: "DESC",
}

// mutation

func (r *Resolver) CreateProductsResolver(
	ctx context.Context,
	productName string,
	description string,
	img string,
	unitPrice int,
	discount float64,
	stock int,
	brand int,
	category int,
) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a product OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.CreateProductParams{
		ProductName: productName,
		Description: sql.NullString{String: description, Valid: true},
		Img:         sql.NullString{String: description, Valid: true},
		UnitPrice:   int32(unitPrice),
		Discount:    discount,
		Stock:       int32(stock),
		BrandID:     int64(brand),
		Category:    int64(category),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = r.store.CreateProduct(gc, args)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create product: %v", err)
	}

	fmt.Println("createProduct OK")
	return res, nil
}

func (r *mutationResolver) GetAllProductsByKeywordResolver(ctx context.Context, keyword string, sortBy model.SortBy, first int, skip int) ([]*model.Product, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// change necessary
	args := db.ListProductByKeywordParams{
		Description: sql.NullString{String: keyword, Valid: true},
		ProductName: keyword,
		Limit:       int32(first),
		Offset:      int32(skip),
	}

	products, err := r.store.ListProductByKeyword(gc, args)
	if err != nil {
		return nil, fmt.Errorf("getAllProductsByKeywordResolver failed : %v", err)
	}

	if sortBy.String() == EnumSort.Asc || sortBy.String() == EnumSort.Desc {
	} else {
		return nil, fmt.Errorf("sort value 'ASC' 'DESC' only: %v", err)
	}

	if sortBy.String() == EnumSort.Desc {
		for i := 0; i < len(products)/2; i++ {
			products[i], products[len(products)-1-i] = products[len(products)-1-i], products[i]
		}
	}

	convertProc := make([]*model.Product, 0)
	for _, p := range products {
		id := strconv.Itoa(int(p.ID))
		category := strconv.Itoa(int(p.Category))
		convertProc = append(convertProc, &model.Product{
			ID:          id,
			ProductName: p.ProductName,
			Description: p.Description.String,
			Img:         p.Img.String,
			UnitPrice:   int(p.UnitPrice),
			Discount:    p.Discount,
			Stock:       int(p.Stock),
			Brand:       int(p.BrandID),
			Category:    category,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return convertProc, nil
}

// query

// dataloaders function
func (r *Resolver) GetAllProductsResolver(ctx context.Context, first int, skip int) ([]*model.Product, error) {
	return r.dataloaders.Retrieve(ctx).GetAllProductsID.Load(int64(skip))
}

// dataloaders function
func (r *Resolver) GetProductResolver(ctx context.Context, productID int) (*model.Product, error) {
	return r.dataloaders.Retrieve(ctx).GetProductID.Load(int64(productID))
}

func (r *Resolver) GetAllProductsByCategoryResolver(ctx context.Context, category string, first int, skip int) ([]*model.Product, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// change necessary
	args := db.ListProductByCategoryParams{
		Category: category,
		Limit:    int32(first),
		Offset:   int32(skip),
	}

	products, err := r.store.ListProductByCategory(gc, args)
	if err != nil {
		return nil, fmt.Errorf("getAllProductsByCategory failed : %v", err)
	}

	convertProc := make([]*model.Product, 0)
	for _, p := range products {
		id := strconv.Itoa(int(p.ID))
		category := strconv.Itoa(int(p.Category))
		convertProc = append(convertProc, &model.Product{
			ID:          id,
			ProductName: p.ProductName,
			Description: p.Description.String,
			Img:         p.Img.String,
			UnitPrice:   int(p.UnitPrice),
			Discount:    p.Discount,
			Stock:       int(p.Stock),
			Brand:       int(p.BrandID),
			Category:    category,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return convertProc, nil
}
