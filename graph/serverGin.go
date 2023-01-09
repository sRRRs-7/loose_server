package graph

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sRRRs-7/loose_style.git/graph/dataloaders"
	"github.com/sRRRs-7/loose_style.git/token"
)

// initialize Gin
func (r *Resolver) GinRouter(tokenMaker token.Maker) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
			"http://localhost:80",
			"http://localhost:3000",
			"http://localhost:8080",
			"http://command-style.com",
			"https://command-style.com",
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
		MaxAge:           60 * time.Minute,
	}))

	// health check router
	healthCheckRouter := router.Group("/api")
	healthCheckRouter.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	})

	// auth router
	playGroundRouter := router.Group("/")
	playGroundRouter.Use(GinContextToContextCookie(tokenMaker))
	playGroundRouter.Use(dataloaders.DataLoaderMiddleware(r.store))
	playGroundRouter.POST("/query", graphqlHandler(r))
	playGroundRouter.GET("/query", graphqlHandler(r))

	// auth router
	adminRouter := router.Group("/admin")
	adminRouter.Use(GinContextToContextCookie(tokenMaker))
	adminRouter.POST("/query", graphqlHandler(r))
	adminRouter.GET("/query", playgroundHandler())

	// create logging files
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// manage endpoint
	fmt.Println("GraphQL playground: ", "http://localhost:8080/admin/query")
	router.Run(r.config.HttpServerAddress)

	//case TLS server
	// r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
