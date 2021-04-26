package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
