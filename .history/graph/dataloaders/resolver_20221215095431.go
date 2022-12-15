package dataloaders

import (
	"context"
	"fmt"
	"time"

	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/model"
)

// dataloader function
func newGetAllCodesID(ctx context.Context, store db.Store) *GetAllCodesLoader {
	return NewGetAllCodesLoader(GetAllCodesConfig{
		Wait:     100 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(limit []int64) ([][]*model.Code, []error) {
			args := db.GetAllCodesParams{
				Limit:  30,
				Offset: int32(limit[0]),
			}

			codes, err := store.GetAllCodes(ctx, args)
			if err != nil {
				return nil, []error{fmt.Errorf("could not get all code : %v", err)}
			}

			convertCode := make(map[int64][]*model.Code, len(limit))
			for _, c := range codes {
				convertCode[limit[0]] = append(convertCode[limit[0]], &model.Code{
					ID:          string(fmt.Sprint(c.ID)),
					Username:    c.Username,
					Code:        c.Code,
					Img:         string(c.Img),
					Description: c.Description,
					Performance: c.Performance,
					Tags:        c.Tags,
					CreatedAt:   c.CreatedAt,
					UpdatedAt:   c.UpdatedAt,
					Access:      int(c.Access),
				})
			}

			result := make([][]*model.Code, len(limit))
			for i, f := range limit {
				result[i] = convertCode[f]
			}

			return result, nil
		},
	})
}
