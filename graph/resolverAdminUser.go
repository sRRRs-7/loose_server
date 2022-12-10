package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/sRRRs-7/loose_style.git/cryptography"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

func (r *mutationResolver) CreateAdminUserResolver(ctx context.Context, username string, password string) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "create a admin user OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return res, fmt.Errorf("admin user password encrypt error: %v", err)
	}

	args := db.CreateAdminUserParams{
		Username: username,
		Password: hashPassword,
		CreatedAt: time.Now(),
	}

	admin, err := r.store.CreateAdminUser(gc, args)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create admin user: %v", err)
	}

	fmt.Println(admin)
	return res, nil
}


func (r *mutationResolver) GetAdminResolver(ctx context.Context, username string, password string) (*model.AdminUserResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("password encrypt error: %v", err)
	}

	args := db.GetAdminUserParams{
		Username: username,
		Password: hashPassword,
	}

	user, err := r.store.GetAdminUser(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin user: %v", err)
	}

	b, err := cryptography.VerifyHash(hashPassword, user.Password)
	if err != nil {
		return nil, fmt.Errorf("admin user verify password error: %v", err)
	}

	admin := &model.AdminUserResponse {
		ID: fmt.Sprint(user.ID),
		IsUsername: user.Username == username,
		IsPassword: b,
	}

	return admin, nil
}
