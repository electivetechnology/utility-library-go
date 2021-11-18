package context

import (
	"errors"
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

func GetLimit(c *gin.Context) (Limit, error) {
	limit, err := GetLimitFromContext(c)
	if err != nil {
		log.Fatalf(err.Error())
		return limit, err
	}

	return limit, nil
}

func NewLimit() Limit {
	limit := Limit{Limit: 0}

	return limit
}

func GetLimitFromContext(ctx *gin.Context) (Limit, error) {
	limit := NewLimit()

	// Get Limit from query (GET method)
	lmt := ctx.Query("limit")

	if lmt == "" {
		// Override with 0
		lmt = "0"
	}

	l, err := strconv.Atoi(lmt)
	if err != nil {
		msg := fmt.Sprintf("Expected limit to be numeric, got '%v' instead", ctx.Query("limit"))
		log.Fatalf(msg)
		return limit, errors.New(msg)
	}

	limit.Limit = l

	return limit, nil
}

func (l Limit) GetLimit() int {
	return l.Limit
}
