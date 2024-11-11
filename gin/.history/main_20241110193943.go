package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
		c.JSON(200, gin.H{
			"msg": file,
		})
	})
	r.Run(":8080")
}
