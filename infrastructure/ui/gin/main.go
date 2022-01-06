package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/handler/customer"
)

func main() {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler, _ := customer.New()
	handler.Make(r)

	r.Run()
}
