package mid

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("方法前")
		c.Next()
		fmt.Println("方法后")
	}
}
