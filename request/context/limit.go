package context

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Limit struct {
	Limit int
}

type LimitType interface {
	GetLimit() int
}

func GetLimit(c *gin.Context) Limit {
	limit := GetLimitFromContext(c)

	return limit
}

func NewLimit() Limit {
	limit := Limit{Limit: 0}

	return limit
}

func GetLimitFromContext(ctx *gin.Context) Limit {
	limit := NewLimit()

	// Get Limit from query (GET method)
	l, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		msg := fmt.Sprintf("Expected limit to be numeric, got '%v' instead", ctx.Query("limit"))
		log.Fatalf(msg)
		return limit
	}

	limit.Limit = l

	return limit
}

func (l Limit) GetLimit() int {
	return l.Limit
}
