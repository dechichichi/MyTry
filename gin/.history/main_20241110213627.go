package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	test2()
}

func test1() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		name := c.PostForm("name")
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
		c.JSON(200, gin.H{
			"msg":  file,
			"name": name,
		})
	})
	r.Run(":8080")
}

func test2() {
	r := gin.Default()
	v1 := r.Group("v1").Use(middle1(), middle2())
	v1.GET("/test", func(c *gin.Context) {
		fmt.Println("我在分组内")
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	v1.GET("/gm", func(c *gin.Context) {
		fmt.Println("我在分组内")
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	r.Run(":8080")
}
