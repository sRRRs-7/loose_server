package graph

import (
	"context"
	"fmt"

	"github.com/sRRRs-7/loose_style.git/graph/model"
)

func (r *mutationResolver) CreateCategoryResolver(ctx context.Context, category string) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "create a category OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	cate, err := r.store.CreateCategory(gc, category)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create cart: %v", err)
	}

	fmt.Println(cate)
	return res, nil
}