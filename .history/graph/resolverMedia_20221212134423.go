package graph

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// mutation
func (r *Resolver) CreateMediaResolver(ctx context.Context, title string, contents string, img string) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
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
		Message: "crete a media OK",
	}

	return res, nil
}

func (r *Resolver) UpdateMediaResolver(ctx context.Context, id string, title string, contents string, img string) (*model.MutationResponse, error) {
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
		Message: "update a media OK",
	}

	return res, nil
}

func (r *Resolver) GetMediaResolver(ctx context.Context, id int) (*model.Media, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	m, err := r.store.GetMedia(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get a media: %v", err)
	}

	media := &model.Media{
		ID:        string(rune(m.ID)),
		Title:     m.Title,
		Contents:  m.Contents,
		Img:       string(m.Img),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return media, nil
}

func (r *Resolver) DeleteMediaResolver(ctx context.Context, id int) (*model.MutationResponse, error) {
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
		Message: "delete a media OK",
	}

	return res, nil
}

// query
func (r *Resolver) GetAllMediaResolver(ctx context.Context, first int, skip int) ([]*model.Media, error) {
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

	convertMedias := make([]*model.Media, 0)
	for _, m := range medias {
		id := strconv.Itoa(int(m.ID))
		convertMedias = append(convertMedias, &model.Media{
			ID:        id,
			Title:     m.Title,
			Contents:  m.Contents,
			Img:       string(m.Img),
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}

	fmt.Println(convertMedias)
	return convertMedias, nil
}
