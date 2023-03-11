package probes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewReadinessProbe(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		var status int
		var content gin.H

		select {
		case <-ctx.Done():
			status = http.StatusGone
			content = gin.H{"status": "bad"}
		default:
			status = http.StatusOK
			content = gin.H{"status": "ok"}

		}

		c.IndentedJSON(status, content)
	}
}
