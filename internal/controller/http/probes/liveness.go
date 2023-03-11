package probes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLivenessProbe() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
