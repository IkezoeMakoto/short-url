package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/IkezoeMakoto/short-url/api/src/middleware"
	"github.com/IkezoeMakoto/short-url/api/src/lib"
	"github.com/IkezoeMakoto/short-url/api/src/controller"
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

	rp, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		fmt.Println(err)
		panic("invalid redis port")
	}
	rc := lib.RedisConnector{
		os.Getenv("REDIS_HOST"),
		rp,
	}
	redisClient := rc.Connect()
	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	// setting routing
	u := controller.Url{redisClient}
	r.GET("/move/:hash", u.Get)

	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
