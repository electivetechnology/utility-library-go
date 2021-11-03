package context

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Offset struct {
	Offset int
}

type OffsetType interface {
	GetOffset() int
}

func GetOffset(c *gin.Context) Offset {
	offset := GetOffsetFromContext(c)

	return offset
}

func NewOffset() Offset {
	offset := Offset{Offset: 0}

	return offset
}

func GetOffsetFromContext(ctx *gin.Context) Offset {
	offset := NewOffset()

	// Get Offset from query (GET method)
	o, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		msg := fmt.Sprintf("Expected offset to be numeric, got '%v' instead", ctx.Query("offset"))
		log.Fatalf(msg)
		return offset
	}

	offset.Offset = o

	return offset
}

func (o Offset) GetOffset() int {
	return o.Offset
}
