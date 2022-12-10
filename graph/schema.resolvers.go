package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sRRRs-7/loose_style.git/graph/generated"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, title string, url string) (*model.MutationResponse, error) {
	res, err := r.CreatePostResolver(ctx, title, url)
	if err != nil {
		return nil, fmt.Errorf("CreatePost error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, votes int) (*model.MutationResponse, error) {
	res, err := r.UpdatePostResolver(ctx, id, votes)
	if err != nil {
		return nil, fmt.Errorf("UpdatePost error: %v", err)
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

func (r *mutationResolver) UpdateMedia(ctx context.Context, id string, title string, contents string, img string) (*model.MutationResponse, error) {
	res, err := r.UpdateMediaResolver(ctx, id, title, contents, img)
	if err != nil {
		return nil, fmt.Errorf("UpdateMedia error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, userID string, password string, username string, email string, sex string, dateOfBirth string, destinationFamilyName string, destinationFirstName string, postcode int, prefectureCode string, city string, street string, building string, phone string) (*model.MutationResponse, error) {
	res, err := r.CreateUserResolver(
		ctx,
		userID,
		password,
		username,
		email,
		sex,
		dateOfBirth,
		destinationFamilyName,
		destinationFirstName,
		postcode,
		prefectureCode,
		city,
		street,
		building,
		phone,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, userID string, password string) (bool, error) {
	res, err := r.LoginUserResolver(ctx, userID, password)
	if err != nil {
		return false, fmt.Errorf("GetOrder error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetAdmin(ctx context.Context, username string, password string) (*model.AdminUserResponse, error) {
	res, err := r.GetAdminResolver(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("GetOrder error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateProducts(ctx context.Context, productName string, description string, img string, unitPrice int, discount float64, stock int, brand int, category int) (*model.MutationResponse, error) {
	res, err := r.CreateProductsResolver(ctx, productName, description, img, unitPrice, discount, stock, brand, category)
	if err != nil {
		return nil, fmt.Errorf("CreateProducts error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetProduct(ctx context.Context, productID int) (*model.Product, error) {
	res, err := r.GetProductResolver(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("GetProduct error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetAllProductsByCategory(ctx context.Context, category string, first int, skip int) ([]*model.Product, error) {
	res, err := r.GetAllProductsByCategoryResolver(ctx, category, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllProductsByCategory error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetAllProductsByKeyword(ctx context.Context, keyword string, sortBy model.SortBy, first int, skip int) ([]*model.Product, error) {
	res, err := r.GetAllProductsByKeywordResolver(ctx, keyword, sortBy, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllProductsByKeyword error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) GetCartItem(ctx context.Context, id int) (*model.Product, error) {
	res, err := r.GetCartItemResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetCartItem error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateCart(ctx context.Context, productID int) (*model.MutationResponse, error) {
	res, err := r.CreateCartResolver(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("CreateCart error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateAdminCart(ctx context.Context, userID int, productID int) (*model.MutationResponse, error) {
	res, err := r.CreateAdminCartResolver(ctx, userID, productID)
	if err != nil {
		return nil, fmt.Errorf("CreateAdminCart error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) DeleteCart(ctx context.Context, id int) (*model.MutationResponse, error) {
	res, err := r.DeleteCartResolver(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("DeleteCart error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, userID string, productID int, quantity int, postage int, price int, status bool) (*model.MutationResponse, error) {
	res, err := r.CreateOrderResolver(ctx, userID, productID, quantity, postage, price, status)
	if err != nil {
		return nil, fmt.Errorf("CreateOrder error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, category string) (*model.MutationResponse, error) {
	res, err := r.CreateCategoryResolver(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("CreateCategory error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateBrand(ctx context.Context, brand string) (*model.MutationResponse, error) {
	res, err := r.CreateBrandResolver(ctx, brand)
	if err != nil {
		return nil, fmt.Errorf("CreateBrand error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateAdminUser(ctx context.Context, username string, password string) (*model.MutationResponse, error) {
	res, err := r.CreateAdminUserResolver(ctx, username, password)
	if err != nil {
		return nil, fmt.Errorf("CreateAdminUser error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateToken(ctx context.Context, userID string) (string, error) {
	res, err := r.CreateTokenResolver(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("CreateToken error: %v", err)
	}
	return res, nil
}

func (r *mutationResolver) CreateAdminToken(ctx context.Context, userID string) (string, error) {
	res, err := r.CreateAdminTokenResolver(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("CreateAdminToken error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllPosts(ctx context.Context, orderBy *model.OrderBy, first int, skip int) ([]*model.Post, error) {
	res, err := r.GetAllPostsResolver(ctx, orderBy, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllPosts error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) AllPostsMeta(ctx context.Context) (*model.PostsMeta, error) {
	res, err := r.AllPostsMetaResolver(ctx)
	if err != nil {
		return nil, fmt.Errorf("AllPostsMeta error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllMedia(ctx context.Context, first int, skip int) ([]*model.Media, error) {
	res, err := r.GetAllMediaResolver(ctx, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllMedia error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetUserList(ctx context.Context, first int, skip int) ([]*model.User, error) {
	res, err := r.GetUserListResolver(ctx, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetUserList error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllProducts(ctx context.Context, first int, skip int) ([]*model.Product, error) {
	res, err := r.GetAllProductsResolver(ctx, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetAllProducts error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllCartItems(ctx context.Context, first int, skip int) ([]*model.ProductCartID, error) {
	res, err := r.GetAllCartItemsResolver(ctx, first, skip)
	if err != nil {
		return nil, fmt.Errorf("GetCartItem error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetAllOrders(ctx context.Context, userID string) ([]*model.Order, error) {
	res, err := r.GetAllOrdersResolver(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("GetAllOrders error: %v", err)
	}
	return res, nil
}

func (r *queryResolver) GetOrder(ctx context.Context, userID string, productID int) (*model.Order, error) {
	res, err := r.GetOrderResolver(ctx, userID, productID)
	if err != nil {
		return nil, fmt.Errorf("GetOrder error: %v", err)
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
