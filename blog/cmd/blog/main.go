package main

import (
	"context"
	"log"

	"github.com/ccsunnyfd/practice/blog/internal/di"
	"github.com/ccsunnyfd/practice/blog/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, closeFunc, err := di.InitApp()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		tag, err := service.CreateTag(context.Background(), client)
		if err != nil {
			log.Printf("failed creating tag: %v", err)
		}
		c.JSON(200, gin.H{
			"message": tag,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
