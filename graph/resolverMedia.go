package graph

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/sRRRs-7/loose_style.git/session"
	"github.com/sRRRs-7/loose_style.git/utils"
)

// mutation
func (r *mutationResolver) CreateMediaResolver(ctx context.Context, title string, contents string, img string) (*model.MutationResponse, error) {
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

	args := db.CreateMediaParams{
		Title:     title,
		Contents:  contents,
		Img:       []byte(img),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = r.store.CreateMedia(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to create media: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "CreateMedia OK",
	}

	return res, nil
}

func (r *mutationResolver) UpdateMediaResolver(ctx context.Context, id string, title string, contents string, img string) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	num, _ := strconv.Atoi(id)
	args := db.UpdateMediaParams{
		ID:       int64(num),
		Title:    title,
		Contents: contents,
		Img:      []byte(img),
	}

	err = r.store.UpdateMedia(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to update a media: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "UpdateMedia OK",
	}

	return res, nil
}

func (r *mutationResolver) GetMediaResolver(ctx context.Context, id int) (*model.Media, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	m, err := r.store.GetMedia(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get a media: %v", err)
	}

	media := &model.Media{
		ID:        string(fmt.Sprint(m.ID)),
		Title:     m.Title,
		Contents:  m.Contents,
		Img:       string(m.Img),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return media, nil
}

func (r *mutationResolver) DeleteMediaResolver(ctx context.Context, id int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	err = r.store.DeleteMedia(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to delete a media: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "DeleteMedia OK",
	}

	return res, nil
}

// query
func (r *queryResolver) GetAllMediaResolver(ctx context.Context, first int, skip int) ([]*model.Media, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	args := db.ListMediaParams{
		Limit:  int32(first),
		Offset: int32(skip),
	}

	medias, err := r.store.ListMedia(gc, args)
	if err != nil {
		return nil, fmt.Errorf("media list error : %v", err)
	}

	convertMedias := make([]*model.Media, len(medias))
	for i, m := range medias {
		convertMedias[i] = &model.Media{
			ID:        string(fmt.Sprint(m.ID)),
			Title:     m.Title,
			Contents:  m.Contents,
			Img:       string(m.Img),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
	}

	return convertMedias, nil
}
