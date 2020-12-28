package gin_cors_mw

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(route *gin.Engine,
	accessControlAllowOrigin string,
	accessControlMaxAge string,
	accessControlAllowMethods string,
	accessControlAllowHeaders string,
	accessControlExposeHeaders string,
	accessControlAllowCredentials string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
		c.Writer.Header().Set("Access-Control-Max-Age", accessControlMaxAge)
		c.Writer.Header().Set("Access-Control-Allow-Methods", accessControlAllowMethods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", accessControlAllowHeaders)
		c.Writer.Header().Set("Access-Control-Expose-Headers", accessControlExposeHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", accessControlAllowCredentials)


		switch c.Request.Method {
		case "GET":
			c.Next()
		case "POST":
			c.Next()
		case "PUT":
			c.Next()
		case "DELETE":
			c.Next()
		case "PATCH":
			c.Next()
		default:
			c.AbortWithStatusJSON(406,"unsupported cors options.")

		}
	}
}
