package graph

import (
	"fmt"
	"io"
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
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		MaxAge:           5 * time.Second,
	}))

	// bearer auth router
	playGroundRouter := router.Group("/")
	playGroundRouter.Use(GinContextToContextMiddleware(tokenMaker))
	playGroundRouter.Use(dataloaders.DataLoaderMiddleware(r.store))
	playGroundRouter.POST("/query", graphqlHandler(r))
	playGroundRouter.GET("/query", graphqlHandler(r))

	// bearer auth router
	adminRouter := router.Group("/admin")
	adminRouter.Use(GinContextMiddleware())
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
