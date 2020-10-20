package Router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("User")
	}
}

func LocationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Location")
	}
}
