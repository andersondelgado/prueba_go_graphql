package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andersondelgado/prueba_go_graphql/pkg/config"
	"github.com/andersondelgado/prueba_go_graphql/pkg/datasources/mysql"
	"github.com/andersondelgado/prueba_go_graphql/pkg/enum"
	"github.com/andersondelgado/prueba_go_graphql/pkg/graphql"
	"github.com/andersondelgado/prueba_go_graphql/pkg/graphql/resolver"
	"github.com/andersondelgado/prueba_go_graphql/pkg/guard"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func graphqlHandler() gin.HandlerFunc {
	Resolver := resolver.NewResolver()
	graphqlConfig := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: Resolver}))

	return func(c *gin.Context) {
		graphqlConfig.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", string(enum.GraphqlURL))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	config.InitEnvironment()
	mysql.InitDefaultDB()
	router := gin.New()

	// Serve frontend static files
	var size int64
	size = 200 * 1024 * 1024
	router.Use(limits.RequestSizeLimiter(size))
	router.Use(static.Serve(string(enum.Root), static.LocalFile(string(enum.DirViews), true)))
	router.Static(string(enum.PathFile), string(enum.DirFile))

	router.Use(guard.CORS())

	router.POST(string(enum.GraphqlURL), graphqlHandler())
	router.POST(string(enum.GraphqlAuthURL), guard.AuthMiddleware(), graphqlHandler())
	router.GET(string(enum.Root), playgroundHandler())

	api := router.Group(string(enum.Api))
	{
		api.GET(string(enum.Root), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	//auth.Routes(api)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = config.Environment.Port
	}
	router.Run(":" + port)
}
