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
		Fetch: func(first []int64) ([][]*model.Code, []error) {
			args := db.GetAllCodesParams{
				Limit:  30,
				Offset: int32(first[0]),
			}

			codes, err := store.GetAllCodes(ctx, args)
			if err != nil {
				return nil, []error{fmt.Errorf("could not get code all list : %v", err)}
			}

			convertCode := make(map[int64][]*model.Code, len(first))
			for _, c := range codes {
				convertCode[first[0]] = append(convertCode[first[0]], &model.Code{
					ID:          string(rune(c.ID)),
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

			result := make([][]*model.Code, len(first))
			for i, f := range first {
				result[i] = convertCode[f]
			}

			return result, nil
		},
	})
}
