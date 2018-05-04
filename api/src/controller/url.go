package controller

import (
	"github.com/IkezoeMakoto/short-url/api/src/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const url_key_prefix = "url-"

const url_expire_hour = 3 * 24

type Url struct {
	RedisClient *lib.RedisClient
}

func (u *Url) Get(c *gin.Context) {
	hash := c.Param("hash")
	url, err := u.RedisClient.Get(url_key_prefix + hash).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

func (u *Url) Add(c *gin.Context) {
	url := c.PostForm("url")
	hash := lib.GetRand()
	_, err := u.RedisClient.SetNX(url_key_prefix+hash, url, url_expire_hour*time.Hour).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not found",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"hash": hash,
	})
}
