package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func middle1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("方法1前")
		c.Next()
		fmt.Println("方法1后")
	}
}

func middle2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("方法2前")
		c.Next()
		fmt.Println("方法2后")
	}
}
