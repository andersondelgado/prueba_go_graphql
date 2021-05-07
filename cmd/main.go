package main

import (
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/andersondelgado/prueba_go_graphql/pkg/config"
	"github.com/andersondelgado/prueba_go_graphql/pkg/guard"
	"github.com/andersondelgado/prueba_go_graphql/pkg/enum"
	"github.com/andersondelgado/prueba_go_graphql/pkg/datasources/mysql"
	"net/http"
	"os"
)

func main()  {
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
