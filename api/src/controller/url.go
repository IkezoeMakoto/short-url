package controller

import (
	"github.com/IkezoeMakoto/short-url/api/src/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

const url_key_prefix = "url-"

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
