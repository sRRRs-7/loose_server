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

func (r *mutationResolver) CreateStarResolver(ctx context.Context, codeID int) (*model.MutationResponse, error) {
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

	args := db.CreateStarParams{
		UserID: int64(userID),
		CodeID: int64(codeID),
	}

	err = r.store.CreateStar(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create star: %v", err)
	}

	star, err := r.store.CountStar(gc, int64(codeID))
	if err != nil {
		return nil, fmt.Errorf("failed to count star: %v", err)
	}

	args2 := db.UpdateStarParams{
		ID:   int64(codeID),
		Star: star + 1,
	}
	err = r.store.UpdateStar(gc, args2)
	if err != nil {
		return nil, fmt.Errorf("failed to update star in a code : %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "create a star OK",
	}

	return res, nil
}

func (r *mutationResolver) CreateAdminStarResolver(ctx context.Context, userID, codeID int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.CreateStarParams{
		UserID: int64(userID),
		CodeID: int64(codeID),
	}

	err = r.store.CreateStar(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to remove star: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "create a star OK",
	}

	return res, nil
}

func (r *mutationResolver) CountStarResolver(ctx context.Context, codeID int) (int, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("gin context convert error: %v", err)
	}

	starCnt, err := r.store.CountStar(gc, int64(codeID))
	if err != nil {
		return 0, fmt.Errorf("failed to count star: %v", err)
	}

	return int(starCnt), nil
}

func (r *mutationResolver) DeleteStarResolver(ctx context.Context, userID, codeID int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.DeleteStarParams{
		UserID: int64(userID),
		CodeID: int64(codeID),
	}

	err = r.store.DeleteStar(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to remove star: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "delete a star OK",
	}

	return res, nil
}
