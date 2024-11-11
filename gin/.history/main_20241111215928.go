package main

import (
	"fmt"
	"io"
	dbdata "main/gorm"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	dbdata.Go()
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
	v1 := r.Group("v1").Use(middle1(), middle2()) //洋葱
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

func test3() {
	db, err := gorm.Open("mysql", "root:Ly05985481282@/ginclass?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.POST("/student")
}
