package gin

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.JSON(200, gin.H{
			"msg": file,
		})
	})
	r.Run("8080")
}
