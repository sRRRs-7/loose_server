package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/sRRRs-7/loose_style.git/cryptography"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/sRRRs-7/loose_style.git/session.go"
)

// mutation

func (r *mutationResolver) CreateAdminCollectionResolver(ctx context.Context, userID, codeID int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.CreateCollectionParams{
		UserID: int64(userID),
		CodeID: int64(codeID),
	}

	err = r.store.CreateCollection(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create cart: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a admin collection OK",
	}

	return res, nil
}

func (r *mutationResolver) CreateCollectionResolver(ctx context.Context, codeID int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// user id get from redis
	authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	fields := strings.Split(authorizationHeader, " ")
	accessToken := fields[1]

	key, err := cryptography.HashPassword(accessToken)
	if err != nil {
		return nil, fmt.Errorf("CreateCollectionResolver error: %v", err)
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
		return nil, fmt.Errorf("GetUser in all collection error : %v", err)
	}

	args := db.CreateCollectionParams{
		UserID: int64(userID),
		CodeID: int64(codeID),
	}

	err = r.store.CreateCollection(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create collection: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a collection OK",
	}

	return res, nil
}

func (r *mutationResolver) GetCollectionResolver(ctx context.Context, id int) (*model.CodeWithCollectionID, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	code, err := r.store.GetCollection(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("GetCollectionResolver error : %v", err)
	}

	star, err := r.store.CountStar(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf(" get CountStar error in GetCollection resolver : %v", err)
	}

	res := &model.CodeWithCollectionID{
		ID:          string(fmt.Sprint(code.ID)),
		Username:    code.Username,
		Code:        code.Code,
		Img:         string(code.Img),
		Description: code.Description,
		Performance: code.Performance,
		Star:        int(star),
		Tags:        code.Tags,
		CreatedAt:   code.CreatedAt,
		UpdatedAt:   code.UpdatedAt,
		Access:      int(code.Access),
	}

	return res, nil
}

func (r *mutationResolver) DeleteCollectionResolver(ctx context.Context, id int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	err = r.store.DeleteCollection(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to delete cart: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "delete a collection OK",
	}

	return res, nil
}

// query

func (r *queryResolver) GetAllCollectionResolver(ctx context.Context, limit, skip int) ([]*model.CodeWithCollectionID, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	fields := strings.Split(authorizationHeader, " ")
	accessToken := fields[1]

	key, err := cryptography.HashPassword(accessToken)
	if err != nil {
		return nil, fmt.Errorf("GetAllCollectionResolver error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, key)
	if redisValue == nil {
		return nil, fmt.Errorf("GetAllCartCollection error in get redis value : %v", err)
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
		return nil, fmt.Errorf("GetUser error in GetAllCollection: %v", err)
	}

	args := db.GetAllCollectionsParams{
		UserID: id,
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	// get all collection
	collections, err := r.store.GetAllCollections(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionResolver error : %v", err)
	}

	convertCol := make([]*model.CodeWithCollectionID, len(collections))
	for i, col := range collections {
		convertCol[i] = &model.CodeWithCollectionID{
			ID:           string(fmt.Sprint(col.ID)),
			Username:     col.Username,
			Code:         col.Code,
			Img:          string(col.Img),
			Description:  col.Description,
			Performance:  col.Performance,
			Tags:         col.Tags,
			CreatedAt:    col.CreatedAt,
			UpdatedAt:    col.UpdatedAt,
			Access:       int(col.Access),
			CollectionID: int(col.ID_2),
		}
	}

	return convertCol, nil
}

func (r *queryResolver) GetAllCollectionBySearchResolver(ctx context.Context, keyword string, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	fields := strings.Split(authorizationHeader, " ")
	accessToken := fields[1]

	key, err := cryptography.HashPassword(accessToken)
	if err != nil {
		return nil, fmt.Errorf("GetAllCollectionResolver error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, key)
	if redisValue == nil {
		return nil, fmt.Errorf("GetAllCartCollection error in get redis value : %v", err)
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
		return nil, fmt.Errorf("GetUser error in GetAllCollection: %v", err)
	}

	args := db.GetAllCollectionsParams{
		UserID: id,
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	// get all collection
	collections, err := r.store.GetAllCollections(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionBySearchResolver error : %v", err)
	}

	convertCol := make([]*model.CodeWithCollectionID, len(collections))
	for i, col := range collections {
		convertCol[i] = &model.CodeWithCollectionID{
			ID:           string(fmt.Sprint(col.ID)),
			Username:     col.Username,
			Code:         col.Code,
			Img:          string(col.Img),
			Description:  col.Description,
			Performance:  col.Performance,
			Tags:         col.Tags,
			CreatedAt:    col.CreatedAt,
			UpdatedAt:    col.UpdatedAt,
			Access:       int(col.Access),
			CollectionID: int(col.ID_2),
		}
	}

	return convertCol, nil
}
