package graph

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// mutation

var posts []*model.Post = make([]*model.Post, 0)

func (r *Resolver) CreatePostResolver(ctx context.Context, title string, url string) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: false,
		Message: "crete a post OK",
	}

	post := &model.Post{
		ID:        fmt.Sprintf("%d", len(posts)+1),
		Title:     title,
		Votes:     0,
		URL:       url,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	posts = append(posts, post)

	fmt.Println(post)
	return res, nil
}

func (r *Resolver) UpdatePostResolver(ctx context.Context, id string, votes int) (*model.MutationResponse, error) {
	res := &model.MutationResponse{
		IsError: true,
		Message: "update a post OK",
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		res.IsError = true
		res.Message = "error"
		return res, fmt.Errorf("id is invalid: %v", id)
	}
	posts[i-1].Votes = votes

	fmt.Println(posts[i-1])
	return res, nil
}

// query

func (r *Resolver) GetAllPostsResolver(ctx context.Context, orderBy *model.OrderBy, first int, skip int) ([]*model.Post, error) {
	if skip > len(posts) {
		skip = len(posts)
	}
	if (first + skip) > len(posts) {
		first = len(posts) - skip
	}
	sortPosts := make([]*model.Post, len(posts))
	copy(sortPosts, posts)
	if orderBy != nil && *orderBy == "createdAt_DESC" {
		sort.SliceStable(sortPosts, func(i, j int) bool {
			return sortPosts[i].CreatedAt > sortPosts[j].CreatedAt
		})
	}
	slicePosts := sortPosts[skip : skip+first]
	return slicePosts, nil
}

func (r *Resolver) AllPostsMetaResolver(ctx context.Context) (*model.PostsMeta, error) {
	res := &model.PostsMeta{
		Count: len(posts),
	}
	return res, nil
}