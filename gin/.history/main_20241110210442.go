package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("v1").Use(middle())
	v1.POST("/login", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		name := c.PostForm("name")
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
		c.Writer.Header().Add("Access-Control-Allow-Origin")
		c.JSON(200, gin.H{
			"msg":  file,
			"name": name,
		})
	})
	r.Run(":8080")
}
