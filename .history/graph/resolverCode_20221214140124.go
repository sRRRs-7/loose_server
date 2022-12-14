package graph

import (
	"context"
	"fmt"
	"strings"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

type SortBy struct {
	Asc  string
	Desc string
}

var EnumSort = SortBy{
	Asc:  "ASC",
	Desc: "DESC",
}

// mutation
func (r *mutationResolver) CreateCodeResolver(ctx context.Context, username string, code string, img string, description string, performance string, star int, tags []string, access int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// tags str convert lower case
	tag := make([]string, len(tags))
	for i, t := range tags {
		tag[i] = strings.ToLower(t)
	}

	args := db.CreateCodeParams{
		Username:    username,
		Code:        code,
		Img:         []byte(img),
		Description: description,
		Performance: performance,
		Star:        int64(star),
		Tags:        tag,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Access:      int64(access),
	}

	err = r.store.CreateCode(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create code: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a code OK",
	}

	return res, nil
}

func (r *mutationResolver) UpdateCodesResolver(ctx context.Context, id int, code string, img string, description string, performance string, tags []string) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "update a code OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.UpdateCodeParams{
		ID:          int64(id),
		Code:        code,
		Img:         []byte(img),
		Description: description,
		Performance: performance,
		Tags:        tags,
		UpdatedAt:   time.Now(),
	}

	err = r.store.UpdateCode(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to get a code: %v", err)
	}

	return res, nil
}

func (r *mutationResolver) UpdateAccessResolver(ctx context.Context, id, access int) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "update a code access count OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	c, err := r.store.GetCode(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get code error in update access: %v", err)
	}

	args := db.UpdateAccessParams{
		ID:     int64(id),
		Access: c.Access + int64(access), // +1 or -1
	}

	err = r.store.UpdateAccess(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to update a code access count : %v", err)
	}

	return res, nil
}

func (r *mutationResolver) DeleteCodeResolver(ctx context.Context, id int) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "delete a code OK",
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	err = r.store.DeleteCode(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to delete a code: %v", err)
	}

	return res, nil
}

func (r *mutationResolver) GetCodeResolver(ctx context.Context, id int) (*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	code, err := r.store.GetCode(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to GetCode: %v", err)
	}

	// get star count
	star, err := r.store.CountStar(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get CountStar: %v", err)
	}

	res := &model.Code{
		ID:          string(fmt.Sprint(code.ID)),
		Username:    code.Username,
		Code:        code.Code,
		Img:         string(code.Img),
		Description: code.Description,
		Performance: code.Performance,
		Star:        int(star),
		Tags:        code.Tags,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Access:      int(code.Access),
	}

	return res, nil
}

// query

func (r *queryResolver) GetAllCodesByTagResolver(ctx context.Context, tags []*string, sortBy model.SortBy, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	t := make([]string, 10)
	for i, tag := range tags {
		t[i] = strings.ToLower(*tag)
	}

	args := db.GetAllCodesByTagParams{
		Column1:  t[0],
		Column2:  t[1],
		Column3:  t[2],
		Column4:  t[3],
		Column5:  t[4],
		Column6:  t[5],
		Column7:  t[6],
		Column8:  t[7],
		Column9:  t[8],
		Column10: t[9],
		Limit:    int32(limit),
		Offset:   int32(skip),
	}

	codes, err := r.store.GetAllCodesByTag(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesByTagsResolver failed : %v", err)
	}

	if sortBy.String() != EnumSort.Asc && sortBy.String() != EnumSort.Desc {
		return nil, fmt.Errorf("sort value 'ASC' 'DESC' only: %v", err)
	}

	if sortBy.String() == EnumSort.Desc {
		for i := 0; i < len(codes)/2; i++ {
			codes[i], codes[len(codes)-1-i] = codes[len(codes)-1-i], codes[i]
		}
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		// get star count
		star, err := r.store.CountStar(gc, int64(c.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to get CountStar: %v", err)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        int(star),
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllCodesByKeywordResolver(ctx context.Context, keyword string, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.GetAllCodesByKeywordParams{
		Username:    "%" + keyword + "%",
		Code:        "%" + keyword + "%",
		Description: "%" + keyword + "%",
		Limit:       int32(limit),
		Offset:      int32(skip),
	}

	codes, err := r.store.GetAllCodesByKeyword(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesByKeywordResolver failed : %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		// get star count
		star, err := r.store.CountStar(gc, int64(c.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to get CountStar: %v", err)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        int(star),
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllCodesSortedStarResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.GetAllCodesSortedStarParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	codes, err := r.store.GetAllCodesSortedStar(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedStarResolver failed : %v", err)
	}
	for i := codes {
		fmt.Println(codes)
	}


	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		// get star count
		star, err := r.store.CountStar(gc, int64(c.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to get CountStar: %v", err)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        int(star),
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllCodesSortedAccessResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.GetAllCodesSortedAccessParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	codes, err := r.store.GetAllCodesSortedAccess(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedAccessResolver failed : %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		// get star count
		star, err := r.store.CountStar(gc, int64(c.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to get CountStar: %v", err)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        int(star),
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
		}
	}

	return list, nil
}

// candidate dataloader function
func (r *queryResolver) GetAllCodesResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.GetAllCodesParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	codes, err := r.store.GetAllCodes(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllCode: %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		// get star count
		star, err := r.store.CountStar(gc, int64(c.ID))
		if err != nil {
			return nil, fmt.Errorf("failed to get CountStar in GetAllCodesResolver: %v", err)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        int(star),
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
		}
	}

	return list, nil
}

// dataloaders function
// func (r *queryResolver) GetAllCodesResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
// 	return r.dataloaders.Retrieve(ctx).GetAllCodesID.Load(int64(skip))
// }
