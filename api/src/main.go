package main

import (
	"os"
	"net/http"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/IkezoeMakoto/short-url/api/src/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	if os.Getenv("APP_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// setting middleware
	r.Use(middleware.NewCors())
	r.Use(cors.Default())
	r.Use(middleware.NewRecovery(log.New(
		gin.DefaultErrorWriter,
		"\n\n\x1b[31m",
		log.LstdFlags)))

	// setting routing
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1.GET("/:path", Index)

	r.Run()
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello %s", c.Param("path"))
}
