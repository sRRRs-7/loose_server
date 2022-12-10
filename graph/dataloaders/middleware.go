package dataloaders

import (
	"context"

	"github.com/gin-gonic/gin"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
)


func DataLoaderMiddleware(store db.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), Key, c)
		loaders := NewLoaders(c, store)
		argumentCtx := context.WithValue(ctx, Key, loaders)
		c.Request = c.Request.WithContext(argumentCtx)
		c.Next()
	}
}