package graph

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sRRRs-7/loose_style.git/cryptography"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// mutation

func (r *Resolver) CreateUserResolver(
	ctx context.Context,
	userID string,
	password string,
	username string,
	email string,
	sex string,
	dateOfBirth string,
	destinationFamilyName string,
	destinationFirstName string,
	postcode int,
	prefectureCode string,
	city string,
	street string,
	building string,
	phone string,
) (
	*model.MutationResponse, error,
	) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a user OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("gin context convert error: %v", err)
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return res, fmt.Errorf("create user password encrypt error: %v", err)
	}

	args := db.CreateUserParams {
		UserID: userID,
		Password: hashPassword,
		Username: username,
		Email: email,
		Sex: sex,
		DataOfBirth: dateOfBirth,
		DestinationFamilyName: destinationFamilyName,
		DestinationFirstName: destinationFirstName,
		Postcode: int32(postcode),
		PrefectureCode: prefectureCode,
		City: city,
		Street: street,
		Building: building,
		Phone: phone,
		CreatedAt: time.Now(),
		UpdatedAt:time.Now(),
	}

	user, err := r.store.CreateUser(gc, args)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("failed to create user: %v", err)
	}

	fmt.Println(user)
	return res, nil
}

// query

func (r *Resolver) LoginUserResolver(ctx context.Context, userID string, password string) (bool, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return false, fmt.Errorf("gin context convert error: %v", err)
	}

	hashPassword, err := cryptography.HashPassword(password)
	if err != nil {
		return false, fmt.Errorf("create user password encrypt error: %v", err)
	}

	args := db.LoginUserParams{
		UserID: userID,
		Password: hashPassword,
	}

	user, err := r.store.LoginUser(gc, args)
	if err != nil {
		return false, fmt.Errorf("auth user method cannot retrieve user from database : %v", err)
	}

	_, err = cryptography.VerifyHash(hashPassword, user.Password)
	if err != nil {
		return false, fmt.Errorf("auth user password verification error: %v", err)
	}

	return true, nil
}

func (r *queryResolver) GetUserListResolver(ctx context.Context, first int, skip int) ([]*model.User, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.ListUsersParams {
		Limit: int32(first),
		Offset: int32(skip),
	}

	users, err := r.store.ListUsers(gc, args)
	if err != nil {
		return nil, fmt.Errorf("auth user method cannot retrieve user info from database : %v", err)
	}

	convertUsers := make([]*model.User, 0)
	for _, u := range users {
		id := strconv.Itoa(int(u.ID))
		convertUsers = append(convertUsers, &model.User {
			ID: id,
			UserID: u.UserID,
			Password: u.Password,
			Username: u.Username,
			Email: u.Email,
			Sex: u.Sex,
			DateOfBirth: u.DataOfBirth,
			DestinationFamilyName: u.DestinationFamilyName,
			DestinationFirstName: u.DestinationFirstName,
			Postcode: int(u.Postcode),
			PrefectureCode: u.PrefectureCode,
			City: u.City,
			Street:u.Street,
			Building: u.Building,
			Phone: u.Phone,
			CreatedAt:u.CreatedAt,
			UpdatedAt:u.UpdatedAt,
		})
	}

	return convertUsers, nil
}