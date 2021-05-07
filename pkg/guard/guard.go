package guard

import (
	"context"
	"fmt"
	"github.com/andersondelgado/prueba_go_graphql/pkg/enum"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/model"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/repository"
	"github.com/andersondelgado/prueba_go_graphql/pkg/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("vary", "Origin")
		//c.Writer.Header().Set("vary", "Access-Control-Request-Method")
		//c.Writer.Header().Set("vary", "Access-Control-Request-Headers")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "custId, appId,X-Requested-With, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Authorization, authenticated")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("x-Frame-Options", "SAMEORIGIN")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "strict-origin")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; frame-src *; img-src * data:; media-src *; object-src *; script-src * 'unsafe-inline' 'unsafe-eval'; style-src * 'unsafe-inline';")
		c.Writer.Header().Set("Permissions-Policy", "fullscreen=()")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if values, _ := c.Request.Header[string(enum.RequestAuthorizationDefault)]; len(values) <= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Empty token authorization"})
			c.Next()
			return
		}

		strToken0 := c.Request.Header.Get(string(enum.RequestAuthorizationDefault))

		//print("strtoken: ",strToken0)
		if !strings.HasPrefix(strToken0, "") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			c.Next()
			return
		}

		strSplit := strings.Split(strToken0, " ")
		if len(strSplit) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			c.Next()
			return
		}
		strToken := strSplit[1]
		//print("\n",strToken)
		prefixToken := strings.HasPrefix(strToken, "")

		if !prefixToken {
			panic("Empty token")
		}

		ID, err := util.DecodeToken(strToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Next()
			return
		}

		rep := repository.NewUserRepository()
		pk, _ := strconv.Atoi(ID)
		entity, er := rep.GetByParam(map[string]interface{}{"id": pk})

		if er != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
			c.Next()
			return
		}
		user := model.User{ID: entity.ID, Username: entity.Username}

		ctx := context.WithValue(c.Request.Context(), string(enum.GinContextKeyAuthDefault), user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
