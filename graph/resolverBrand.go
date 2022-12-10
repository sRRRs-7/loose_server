package graph

import (
	"context"
	"fmt"

	"github.com/sRRRs-7/loose_style.git/graph/model"
)

func (r *mutationResolver) CreateBrandResolver(ctx context.Context, brand string) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "create a brand OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	b, err := r.store.CreateBrand(gc, brand)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create brand: %v", err)
	}

	fmt.Println(b)
	return res, nil
}