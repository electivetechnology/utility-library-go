package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceNotFound struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
