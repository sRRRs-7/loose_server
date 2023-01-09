package graph

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sRRRs-7/loose_style.git/session"
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

// gin context Bearer authentication
func GinContextToContextMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader(authorizationHeaderKey) != "Bearer undefined" {
			// verify header
			authorizationHeader := c.GetHeader(authorizationHeaderKey)
			if strings.Contains(authorizationHeader, "undefined") {
				err := errors.New("authorization header is undefined")
				log.Println("authorization header is undefined")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provide")
				log.Println("authorization header is not provide")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// verify paseto token length 3 (version, purpose, payload)
			fields := strings.Split(authorizationHeader, " ")
			if len(fields) < 2 {
				err := errors.New("invalid authorization header for less than 2 length")
				log.Println("invalid authorization header for less than 2 length")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// verify authenticate type
			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := errors.New("invalid authorization type. need bearer authorization header")
				log.Println("invalid authorization type. need bearer authorization header")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// expired token
			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				err := errors.New("invalid token expired ")
				log.Println("invalid token expired ")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			c.Set(authorizationHeaderKey, payload)

			// check session
			redisValue := session.GetSession(c, accessToken)
			if redisValue == nil {
				err := errors.New("redis value not exist")
				log.Println("redis value not exist")
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

// gin context Cookie authentication
func GinContextToContextCookie(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get cookie
		cookie, err := c.Cookie("228e81fb33c862aa")
		if err != nil {
			// convert gin context
			ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			if cookie == "" {
				err := errors.New("no Cookie")
				log.Println("No Cookie")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// verify paseto token length 3 (version, purpose, payload)
			fields := strings.Split(cookie, ".")
			if len(fields) < 4 {
				err := errors.New("invalid cookie token for less than 4 length")
				log.Println("invalid cookie token for less than 4 length")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			// expired token
			payload, err := tokenMaker.VerifyToken(cookie)
			if err != nil {
				err := errors.New("invalid token expired ")
				log.Println("invalid token expired ")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}
			c.Set(authorizationHeaderKey, payload)

			// check session
			redisValue := session.GetSession(c, cookie)
			if redisValue == nil {
				err := errors.New("redis value not exist")
				log.Println("redis value not exist")
				c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}

			// convert gin context
			ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}

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
