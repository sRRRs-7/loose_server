package graph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sRRRs-7/loose_style.git/session.go"
	"github.com/sRRRs-7/loose_style.git/token"
)

// Gin gonic server logic
type key string

const GinContextKey key = "GinContextKey"

// convert gin context key
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}
	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// gin context middleware playground
func GinContextToContextMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader(authorizationHeaderKey) != "Bearer undefined" {
			// verify header
			authorizationHeader := c.GetHeader(authorizationHeaderKey)
			if strings.Contains(authorizationHeader, "undefined") {
				err := errors.New("authorization header is undefined")
				fmt.Println("authorization header is undefined")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provide")
				fmt.Println("authorization header is not provide")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// verify paseto token length 3 (version, purpose, payload)
			fields := strings.Split(authorizationHeader, " ")
			if len(fields) < 2 {
				err := errors.New("invalid authorization header for less than 2 length")
				fmt.Println("invalid authorization header for less than 2 length")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// verify authenticate type
			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := errors.New("invalid authorization type. need bearer authorization header")
				fmt.Println("invalid authorization type. need bearer authorization header")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// expired token
			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				err := errors.New("invalid token expired ")
				fmt.Println("invalid token expired ")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			c.Set(authorizationHeaderKey, payload)

			// check session
			redisValue := session.GetSession(c, accessToken)
			if redisValue == nil {
				err := errors.New("expired redis value token  session")
				fmt.Println("expired redis value token session")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
		}

		// convert gin context
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
