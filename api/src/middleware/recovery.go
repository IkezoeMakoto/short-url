package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRecovery(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Printf("[Recovery] panic recovered:\n%s\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "INTERNAL_SERVER_ERROR",
				})
				return
			}
		}()

		c.Next()
	}
}
