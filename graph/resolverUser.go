package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/sRRRs-7/loose_style.git/cryptography"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// mutation

func (r *mutationResolver) CreateUserResolver(ctx context.Context, username string, password string, email string, sex string, dateOfBirth string) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("create user password encrypt error: %v", err)
	}

	args := db.CreateUserParams{
		Username:    username,
		Password:    hashPassword,
		Email:       email,
		Sex:         sex,
		DataOfBirth: dateOfBirth,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = r.store.CreateUser(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "CreateUser OK",
	}

	return res, nil
}

func (r *mutationResolver) LoginUserResolver(ctx context.Context, username string, password string) (*model.LoginUserResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("create user password encrypt error: %v", err)
	}

	args := db.LoginUserParams{
		Username: username,
		Password: hashPassword,
	}

	user, err := r.store.LoginUser(gc, args)
	if err != nil {
		return nil, fmt.Errorf("auth user method cannot retrieve user from database : %v", err)
	}

	_, err = cryptography.VerifyHash(hashPassword, user.Password)
	if err != nil {
		return nil, fmt.Errorf("auth user password verification error: %v", err)
	}

	res := &model.LoginUserResponse{
		UserID:   int(user.ID),
		Ok:       true,
		Username: username,
	}

	return res, nil
}

func (r *mutationResolver) UpdateUserResolver(ctx context.Context, username string, updateName string, email string) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.UpdateUserParams{
		Username:   username,
		Username_2: updateName,
		Email:      email,
		UpdatedAt:  time.Now(),
	}

	err = r.store.UpdateUser(gc, args)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser error : %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "UpdateUser OK",
	}

	return res, nil
}

func (r *mutationResolver) DeleteUserResolver(ctx context.Context, username string) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in DeleteUserResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		return nil, fmt.Errorf("GetUser error in deleteUserResolver: %v", err)
	}

	err = r.store.DeleteUser(gc, user.ID)
	if err != nil {
		return nil, fmt.Errorf("delete User error: %v", err)
	}

	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in DeleteUserResolver : %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "DeleteUser OK",
	}

	return res, nil
}
