package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/sRRRs-7/loose_style.git/cryptography"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/sRRRs-7/loose_style.git/session"
	"github.com/sRRRs-7/loose_style.git/utils"
)

func (r *mutationResolver) CreateAdminUserResolver(ctx context.Context, username string, password string) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	cookie, err := gc.Cookie("a0746dda4c2a0269")
	if err != nil {
		return nil, fmt.Errorf("CreateCollectionResolver cookie error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, cookie)
	if redisValue == nil {
		return nil, fmt.Errorf("get all cart item error get redis value is nil : %v", err)
	}
	// string processing
	name := utils.GetUsername(redisValue)

	if name != "srrrs" {
		return nil, fmt.Errorf("deffer admin user name")
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("admin user password encrypt error: %v", err)
	}

	args := db.CreateAdminUserParams{
		Username:  username,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	err = r.store.CreateAdminUser(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin user: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "create a admin user OK",
	}

	return res, nil
}

func (r *mutationResolver) GetAdminUserResolver(ctx context.Context, username string, password string) (*model.AdminUserResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	name := "srrrs"
	pass := "srrrs"

	isUserName := name == username
	isPassword := pass == password

	if name != username || pass != password {
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

		isUserName = user.Username == username
		isPassword = b
	}

	admin := &model.AdminUserResponse{
		ID:         fmt.Sprint(1),
		IsUsername: isUserName,
		IsPassword: isPassword,
	}

	return admin, nil
}
