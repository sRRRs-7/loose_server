package graph

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// mutation

func (r *Resolver) CreateOrderResolver(ctx context.Context, userID string, productID int, quantity int, postage int, price int, status bool) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "create a order OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}


	args := db.CreateOrderParams {
		UserID: userID,
		ProductID: int64(productID),
		Quantity: int32(quantity),
		Postage: int32(postage),
		Price: int32(price),
		Status: false,
		CreatedAt: time.Now(),
	}

	order, err := r.store.CreateOrder(gc, args)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create order: %v", err)
	}

	fmt.Println(order)
	return res, nil
}

// query

func (r *Resolver) GetAllOrdersResolver(ctx context.Context, userID string) ([]*model.Order, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	orders, err := r.store.ListOrders(gc, userID)
	if err != nil {
		return nil, fmt.Errorf("GetAllOrders error: %v", err)
	}
	convertOrder := make([]*model.Order, 0)
	for _, o := range orders {
		id := strconv.Itoa(int(o.ID))
		convertOrder = append(convertOrder, &model.Order{
			ID: id,
			UserID: userID,
			ProductID: int(o.ProductID),
			Quantity: int(o.Quantity),
			Postage: int(o.Postage),
			Price: int(o.Price),
			Status: o.Status,
			CreatedAt: time.Now(),
		})
	}

	return convertOrder, nil
}

func (r *Resolver) GetOrderResolver(ctx context.Context, userID string, productID int) (*model.Order, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.GetOrderParams {
		UserID: userID,
		ProductID: int64(productID),
	}

	order, err := r.store.GetOrder(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetOrder error: %v", err)
	}

	id := strconv.Itoa(int(order.ID))
	convertOrder := &model.Order{
		ID: id,
		UserID: userID,
		ProductID: int(order.ProductID),
		Quantity: int(order.Quantity),
		Postage: int(order.Postage),
		Price: int(order.Price),
		Status: order.Status,
		CreatedAt: order.CreatedAt,
	}

	return convertOrder, nil
}