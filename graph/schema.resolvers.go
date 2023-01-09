package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sRRRs-7/loose_style.git/graph/generated"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

func (r *mutationResolver) CreateAdminUser(ctx context.Context, username string, password string) (*model.MutationResponse, error) {
	res, err := r.CreateAdminUserResolver(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("CreateAdminUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetAdminUser(ctx context.Context, username string, password string) (*model.AdminUserResponse, error) {
	res, err := r.GetAdminUserResolver(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("GetAdminUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) AdminCreateCode(ctx context.Context, username string, code string, img string, description string, performance string, star []int, tags []string, access int) (*model.MutationResponse, error) {
	res, err := r.AdminCreateCodeResolver(ctx, username, code, img, description, performance, star, tags, access)
	if err != nil {
		return nil, fmt.Errorf("CreateCode error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateCode(ctx context.Context, code string, img string, description string, performance string, star []int, tags []string, access int) (*model.MutationResponse, error) {
	res, err := r.CreateCodeResolver(ctx, code, img, description, performance, star, tags, access)
	if err != nil {
		return nil, fmt.Errorf("CreateCode error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) UpdateCodes(ctx context.Context, id int, code string, img string, description string, performance string, tags []string) (*model.MutationResponse, error) {
	res, err := r.UpdateCodesResolver(ctx, id, code, img, description, performance, tags)
	if err != nil {
		return nil, fmt.Errorf("UpdateCodes error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) UpdateStar(ctx context.Context, codeID int) (*model.MutationResponse, error) {
	res, err := r.UpdateStarResolver(ctx, codeID)
	if err != nil {
		return nil, fmt.Errorf("UpdateAccess error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) UpdateAccess(ctx context.Context, id int, access int) (*model.MutationResponse, error) {
	res, err := r.UpdateAccessResolver(ctx, id, access)
	if err != nil {
		return nil, fmt.Errorf("UpdateAccess error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) DeleteCode(ctx context.Context, id int) (*model.MutationResponse, error) {
	res, err := r.DeleteCodeResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("DeleteCode error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateCollection(ctx context.Context, codeID int) (*model.MutationResponse, error) {
	res, err := r.CreateCollectionResolver(ctx, codeID)
	if err != nil {
		return nil, fmt.Errorf("CreateCollection error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateAdminCollection(ctx context.Context, userID int, codeID int) (*model.MutationResponse, error) {
	res, err := r.CreateAdminCollectionResolver(ctx, userID, codeID)
	if err != nil {
		return nil, fmt.Errorf("CreateAdminCollection error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetCollection(ctx context.Context, id int) (*model.CodeWithCollectionID, error) {
	res, err := r.GetCollectionResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetCollection error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) DeleteCollection(ctx context.Context, id int) (*model.MutationResponse, error) {
	res, err := r.DeleteCollectionResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("DeleteCollection error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) UpdateMedia(ctx context.Context, id string, title string, contents string, img string) (*model.MutationResponse, error) {
	res, err := r.UpdateMediaResolver(ctx, id, title, contents, img)
	if err != nil {
		return nil, fmt.Errorf("UpdateMedia error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateMedia(ctx context.Context, title string, contents string, img string) (*model.MutationResponse, error) {
	res, err := r.CreateMediaResolver(ctx, title, contents, img)
	if err != nil {
		return nil, fmt.Errorf("CreateMedia error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetMedia(ctx context.Context, id int) (*model.Media, error) {
	res, err := r.GetMediaResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetMedia error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) DeleteMedia(ctx context.Context, id int) (*model.MutationResponse, error) {
	res, err := r.DeleteMediaResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("DeleteMedia error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, username string, password string, email string, sex string, dateOfBirth string) (*model.MutationResponse, error) {
	res, err := r.CreateUserResolver(
		ctx,
		username,
		password,
		email,
		sex,
		dateOfBirth,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, username string, updateName string, email string) (*model.MutationResponse, error) {
	res, err := r.UpdateUserResolver(ctx, username, updateName, email)
	if err != nil {
		return nil, fmt.Errorf("UpdateUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, username string, password string) (*model.LoginUserResponse, error) {
	res, err := r.LoginUserResolver(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("GetOrder error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, username string) (*model.MutationResponse, error) {
	res, err := r.DeleteUserResolver(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("DeleteUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateToken(ctx context.Context, username string) (string, error) {
	res, err := r.CreateTokenResolver(ctx, username)
	if err != nil {
		return "", fmt.Errorf("CreateToken error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateAdminToken(ctx context.Context, username string, password string) (string, error) {
	res, err := r.CreateAdminTokenResolver(ctx, username)
	if err != nil {
		return "", fmt.Errorf("CreateAdminToken error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCodes(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	res, err := r.GetAllCodesResolver(ctx, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodes error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCodesByKeyword(ctx context.Context, keyword string, limit int, skip int) ([]*model.Code, error) {
	res, err := r.GetAllCodesByKeywordResolver(ctx, keyword, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesByKeyword error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCodesSortedStar(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	res, err := r.GetAllCodesSortedStarResolver(ctx, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedStar error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCodesSortedAccess(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	res, err := r.GetAllCodesSortedAccessResolver(ctx, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedAccess error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCodesByTag(ctx context.Context, tags []*string, sortBy model.SortBy, limit int, skip int) ([]*model.Code, error) {
	res, err := r.GetAllCodesByTagResolver(ctx, tags, sortBy, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesByTag error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllOwnCodes(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	res, err := r.GetAllOwnCodesResolver(ctx, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetOwnCodes error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetCode(ctx context.Context, id int) (*model.Code, error) {
	res, err := r.GetCodeResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetCode error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCollection(ctx context.Context, limit int, skip int) ([]*model.CodeWithCollectionID, error) {
	res, err := r.GetAllCollectionResolver(ctx, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCollection error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCollectionBySearch(ctx context.Context, keyword string, limit int, skip int) ([]*model.CodeWithCollectionID, error) {
	res, err := r.GetAllCollectionBySearchResolver(ctx, keyword, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllCollectionBySearch error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllMedia(ctx context.Context, limit int, skip int) ([]*model.Media, error) {
	res, err := r.GetAllMediaResolver(ctx, limit, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllMedia error: %v", err)
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
