package graph

import (
	"context"
	"fmt"
	"strings"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
	"github.com/sRRRs-7/loose_style.git/session"
	"github.com/sRRRs-7/loose_style.git/utils"
)

type SortBy struct {
	Asc  string
	Desc string
}

var EnumSort = SortBy{
	Asc:  "ASC",
	Desc: "DESC",
}

func (r *mutationResolver) AdminCreateCodeResolver(ctx context.Context, username string, code string, img string, description string, performance string, star []int, tags []string, access int) (*model.MutationResponse, error) {
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

	// tags str convert lower case
	tag := make([]string, len(tags))
	for i, t := range tags {
		tag[i] = strings.ToLower(t)
	}

	stars := make([]int64, len(star))
	for i := range star {
		num := star[i]
		stars[i] = int64(num)
	}

	args := db.CreateCodeParams{
		Username:    username,
		Code:        code,
		Img:         []byte(img),
		Description: description,
		Performance: performance,
		Star:        stars,
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
		Message: "AdminCreateCode OK",
	}

	return res, nil
}

// mutation
func (r *mutationResolver) CreateCodeResolver(ctx context.Context, code string, img string, description string, performance string, star []int, tags []string, access int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")
	// accessToken := fields[1]

	// key, err := cryptography.HashPassword(accessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("CreateCollectionResolver error: %v", err)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("CreateCollectionResolver cookie error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, cookie)
	if redisValue == nil {
		return nil, fmt.Errorf("get all cart item error get redis value is nil : %v", err)
	}
	// string processing
	username := utils.GetUsername(redisValue)

	// tags str convert lower case
	tag := make([]string, 0)
	for i := range tags {
		if utils.CheckRegex("\\S", tags[i]) {
			tag = append(tag, tags[i])
		}
	}
	for i, t := range tag {
		tag[i] = strings.ToLower(t)
	}

	stars := make([]int64, len(star))
	for i := range star {
		num := star[i]
		stars[i] = int64(num)
	}

	args := db.CreateCodeParams{
		Username:    username,
		Code:        code,
		Img:         []byte(img),
		Description: description,
		Performance: performance,
		Star:        stars,
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
		Message: "CreateCode OK",
	}

	return res, nil
}

func (r *mutationResolver) UpdateCodesResolver(ctx context.Context, id int, code string, img string, description string, performance string, tags []string) (*model.MutationResponse, error) {
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
		return nil, fmt.Errorf("failed to UpdateCodes: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "UpdateCodes OK",
	}

	return res, nil
}

func (r *mutationResolver) UpdateAccessResolver(ctx context.Context, id, access int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	c, err := r.store.GetCode(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to GetCode error in UpdateAccessResolver: %v", err)
	}

	args := db.UpdateAccessParams{
		ID:     int64(id),
		Access: c.Access + int64(access), // +1 or -1
	}

	err = r.store.UpdateAccess(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to UpdateAccess : %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "UpdateAccess OK",
	}

	return res, nil
}

func (r *mutationResolver) UpdateStarResolver(ctx context.Context, codeID int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}
	// // bearer auth
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")
	// accessToken := fields[1]
	// // token hash
	// key, err := cryptography.HashPassword(accessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("UpdateStarResolver error: %v", err)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("UpdateStarResolver cookie error: %v", err)
	}

	// redis value get
	redisValue := session.GetRedis(gc, cookie)
	if redisValue == nil {
		return nil, fmt.Errorf("UpdateStarResolver error in get redis value : %v", err)
	}
	// string processing
	username := utils.GetUsername(redisValue)

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in UpdateStarResolver: %v", err)
	}
	defer tx.Rollback(gc)

	// get user id
	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		return nil, fmt.Errorf("GetUser error in UpdateStarResolver: %v", err)
	}

	code, err := r.store.GetCode(gc, int64(codeID))
	if err != nil {
		return nil, fmt.Errorf("failed to get code error in UpdateStarResolver: %v", err)
	}

	stars := utils.StarContains(code.Star, user.ID)

	args := db.UpdateStarParams{
		ID:   int64(codeID),
		Star: stars,
	}

	err = r.store.UpdateStar(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to update a star : %v", err)
	}
	//commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in UpdateStarResolver : %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "UpdateStar OK",
	}

	return res, nil
}

func (r *mutationResolver) DeleteCodeResolver(ctx context.Context, id int) (*model.MutationResponse, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	err = r.store.DeleteCode(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to delete a code: %v", err)
	}

	res := &model.MutationResponse{
		IsError: false,
		Message: "DeleteCode OK",
	}
	return res, nil
}

func (r *queryResolver) GetCodeResolver(ctx context.Context, id int) (*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]
	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	var username string
	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err == nil {
		redisValue := session.GetRedis(gc, cookie)
		// string processing
		username = utils.GetUsername(redisValue)
	}

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetCodeResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
	}

	code, err := r.store.GetCode(gc, int64(id))
	if err != nil {
		return nil, fmt.Errorf("failed to GetCode: %v", err)
	}
	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetCodeResolver : %v", err)
	}

	stars := make([]int, len(code.Star))
	for i := range code.Star {
		num := code.Star[i]
		stars[i] = int(num)
	}

	res := &model.Code{
		ID:          string(fmt.Sprint(code.ID)),
		Username:    code.Username,
		Code:        code.Code,
		Img:         string(code.Img),
		Description: code.Description,
		Performance: code.Performance,
		Star:        stars,
		Tags:        code.Tags,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Access:      int(code.Access),
		UserID:      int(user.ID),
	}

	return res, nil
}

// query

func (r *queryResolver) GetAllCodesByTagResolver(ctx context.Context, tags []*string, sortBy model.SortBy, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]
	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesByTagResolver cookie error: %v", err)
	}
	redisValue := session.GetRedis(gc, cookie)
	// string processing
	username := utils.GetUsername(redisValue)

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetAllCodesByTagResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
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
		return nil, fmt.Errorf("GetAllCodesByTagResolver failed : %v", err)
	}
	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetAllCodesByTagResolver : %v", err)
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
		stars := make([]int, len(c.Star))
		for i := range c.Star {
			num := c.Star[i]
			stars[i] = int(num)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        stars,
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
			UserID:      int(user.ID),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllCodesByKeywordResolver(ctx context.Context, keyword string, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]

	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesByKeywordResolver cookie error: %v", err)
	}
	redisValue := session.GetRedis(gc, cookie)
	// string processing
	username := utils.GetUsername(redisValue)

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetAllCodesByKeywordResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
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

	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetAllCodesByKeywordResolver : %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		stars := make([]int, len(c.Star))
		for i := range c.Star {
			num := c.Star[i]
			stars[i] = int(num)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        stars,
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
			UserID:      int(user.ID),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllCodesSortedStarResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]
	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedStarResolver cookie error: %v", err)
	}
	redisValue := session.GetRedis(gc, cookie)
	// string processing
	username := utils.GetUsername(redisValue)

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetAllCodesSortedStarResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
	}

	args := db.GetAllCodesSortedStarParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	codes, err := r.store.GetAllCodesSortedStar(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllCodesSortedStarResolver: %v", err)
	}

	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetAllCodesSortedStarResolver : %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		stars := make([]int, len(c.Star))
		for i := range c.Star {
			num := c.Star[i]
			stars[i] = int(num)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        stars,
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
			UserID:      int(user.ID),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllCodesSortedAccessResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]
	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedAccessResolver cookie error: %v", err)
	}
	redisValue := session.GetRedis(gc, cookie)
	// string processing
	username := utils.GetUsername(redisValue)

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetAllCodesSortedAccessResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
	}

	args := db.GetAllCodesSortedAccessParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	codes, err := r.store.GetAllCodesSortedAccess(gc, args)
	if err != nil {
		return nil, fmt.Errorf("GetAllCodesSortedAccessResolver failed : %v", err)
	}

	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetAllCodesSortedAccessResolver : %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		stars := make([]int, len(c.Star))
		for i := range c.Star {
			num := c.Star[i]
			stars[i] = int(num)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        stars,
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
			UserID:      int(user.ID),
		}
	}

	return list, nil
}

func (r *queryResolver) GetAllOwnCodesResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("gin context convert error: %v", err)
	}

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]
	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err != nil {
		return nil, fmt.Errorf("GetAllOwnCodesResolver cookie error: %v", err)
	}
	redisValue := session.GetRedis(gc, cookie)
	// string processing
	username := utils.GetUsername(redisValue)

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetAllOwnCodesResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
	}

	arg := db.GetAllOwnCodesParams{
		Username: user.Username,
		Limit:    int32(limit),
		Offset:   int32(skip),
	}

	codes, err := r.store.GetAllOwnCodes(gc, arg)
	if err != nil {
		return nil, fmt.Errorf("GetAllOwnCodesResolver failed : %v", err)
	}

	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetAllOwnCodesResolver : %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		stars := make([]int, len(c.Star))
		for i := range c.Star {
			num := c.Star[i]
			stars[i] = int(num)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        stars,
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
			UserID:      int(user.ID),
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

	// // user id get from redis
	// authorizationHeader := gc.GetHeader(authorizationHeaderKey)
	// fields := strings.Split(authorizationHeader, " ")

	// var username string
	// if fields[1] != "undefined" {
	// 	accessToken := fields[1]

	// 	key, _ := cryptography.HashPassword(accessToken)
	// 	// redis value get
	// 	redisValue := session.GetRedis(gc, key)
	// 	// string processing
	// 	username = utils.GetUsername(redisValue)
	// }

	var username string
	cookie, err := gc.Cookie("228e81fb33c862aa")
	if err == nil {
		redisValue := session.GetRedis(gc, cookie)
		// string processing
		username = utils.GetUsername(redisValue)
	}

	// transaction
	tx, err := r.tx.Begin(gc)
	if err != nil {
		return nil, fmt.Errorf("transaction begin error in GetAllCodesResolver: %v", err)
	}
	defer tx.Rollback(gc)

	user, err := r.store.GetUserByUsername(gc, username)
	if err != nil {
		user.ID = 0
	}

	args := db.GetAllCodesParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	}

	codes, err := r.store.GetAllCodes(gc, args)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllCode in GetAllCodesResolver: %v", err)
	}

	// commit
	if err = tx.Commit(gc); err != nil {
		return nil, fmt.Errorf("transaction commit error in GetAllCodesResolver: %v", err)
	}

	list := make([]*model.Code, len(codes))
	for i, c := range codes {
		stars := make([]int, len(c.Star))
		for i := range c.Star {
			num := c.Star[i]
			stars[i] = int(num)
		}
		list[i] = &model.Code{
			ID:          string(fmt.Sprint(c.ID)),
			Username:    c.Username,
			Code:        c.Code,
			Img:         string(c.Img),
			Description: c.Description,
			Performance: c.Performance,
			Star:        stars,
			Tags:        c.Tags,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
			Access:      int(c.Access),
			UserID:      int(user.ID),
		}
	}

	return list, nil
}

// dataloaders function
// func (r *queryResolver) GetAllCodesResolver(ctx context.Context, limit int, skip int) ([]*model.Code, error) {
// 	return r.dataloaders.Retrieve(ctx).GetAllCodesID.Load(int64(skip))
// }
