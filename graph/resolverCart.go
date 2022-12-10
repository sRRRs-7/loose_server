package graph

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sRRRs-7/loose_style.git/cryptography"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/sRRRs-7/loose_style.git/session.go"
)

// mutation

func (r *Resolver) CreateAdminCartResolver(ctx context.Context, userID, productID int) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a cart OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.CreateCartItemParams {
		UserID: int64(userID),
		ProductID: int64(productID),
	}

	user, err := r.store.CreateCartItem(gc, args)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create cart: %v", err)
	}

	fmt.Println(user)
	return res, nil
}


func (r *Resolver) CreateCartResolver(ctx context.Context, productID int) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a cart OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	// user id get from redis
	authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	fields := strings.Split(authorizationHeader, " ")
	accessToken := fields[1]

	key, err := cryptography.HashPassword(accessToken)
	if err != nil {
		return nil, fmt.Errorf("GetAllCartItemResolver error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, key)
	if redisValue == nil {
		return nil, fmt.Errorf("get all cart item error get redis value is nil : %v", err)
	}
	// string processing
	s := strings.Split(redisValue.String(), ",")
	s = strings.Split(s[1], ":")
	userId := s[1]
	userId = userId[1:]
	userId = userId[:len(userId)-1]

	// get user id
	userID, err := r.store.GetUser(gc, userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser in all cart error : %v", err)
	}

	args := db.CreateCartItemParams {
		UserID: int64(userID),
		ProductID: int64(productID),
	}

	user, err := r.store.CreateCartItem(gc, args)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create cart: %v", err)
	}

	fmt.Println(user)
	return res, nil
}

func (r *mutationResolver) GetCartItemResolver(ctx context.Context, id int) (*model.Product, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	product, err := r.store.GetProduct(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("GetCartItemResolver error : %v", err)
	}

	cart_id := strconv.Itoa(int(product.ID))
	cate := strconv.Itoa(int(product.Category))
	convertProc := &model.Product {
		ID: cart_id,
		ProductName: product.ProductName,
		Description: product.Description.String,
		Img: product.Img.String,
		UnitPrice: int(product.UnitPrice),
		Discount: product.Discount,
		Stock: int(product.Stock),
		Brand: int(product.BrandID),
		Category: cate,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	return convertProc, nil
}

func (r *Resolver) DeleteCartResolver(ctx context.Context, id int) (*model.MutationResponse, error)  {
	res := &model.MutationResponse{
		IsError: false,
		Message: "delete a cart OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	err = r.store.DeleteCartItem(gc, int64(id))
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to delete cart: %v", err)
	}

	return res, nil
}

// query

func (r *queryResolver) GetAllCartItemsResolver(ctx context.Context, first, skip int) ([]*model.ProductCartID, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	fields := strings.Split(authorizationHeader, " ")
	accessToken := fields[1]

	key, err := cryptography.HashPassword(accessToken)
	if err != nil {
		return nil, fmt.Errorf("GetAllCartItemResolver error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, key)
	if redisValue == nil {
		return nil, fmt.Errorf("get all cart item error get redis value is nil : %v", err)
	}
	// string processing
	s := strings.Split(redisValue.String(), ",")
	s = strings.Split(s[1], ":")
	userId := s[1]
	userId = userId[1:]
	userId = userId[:len(userId)-1]

	// get user id
	id, err := r.store.GetUser(gc, userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser in all cart error : %v", err)
	}

	args := db.GetAllCartItemParams {
		UserID: id,
		Limit: int32(first),
		Offset: int32(skip),
	}

	// all cart item
	products, err := r.store.GetAllCartItem(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetCartResolver error : %v", err)
	}

	convertProc := make([]*model.ProductCartID, 0)
	for _, p := range products {
		id := strconv.Itoa(int(p.ID))
		cate := strconv.Itoa(int(p.Category))
		convertProc = append(convertProc, &model.ProductCartID {
			ID: id,
			ProductName: p.ProductName,
			Description: p.Description.String,
			Img: p.Img.String,
			UnitPrice: int(p.UnitPrice),
			Discount: p.Discount,
			Stock: int(p.Stock),
			Brand: int(p.BrandID),
			Category: cate,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			CartID: int(p.ID_2),
		})
	}

	return convertProc, nil
}

