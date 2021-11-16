package context

import (
	"errors"
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

func GetOffset(c *gin.Context) (Offset, error) {
	offset, err := GetOffsetFromContext(c)
	if err != nil {
		log.Fatalf(err.Error())
		return offset, err
	}

	return offset, nil
}

func NewOffset() Offset {
	offset := Offset{Offset: 0}

	return offset
}

func GetOffsetFromContext(ctx *gin.Context) (Offset, error) {
	offset := NewOffset()

	// Get Offset from query (GET method)
	off := ctx.Query("offset")

	if off == "" {
		// Override with 0
		off = "0"
	}

	o, err := strconv.Atoi(off)
	if err != nil {
		msg := fmt.Sprintf("Expected offset to be numeric, got '%v' instead", ctx.Query("offset"))
		log.Fatalf(msg)
		return offset, errors.New(msg)
	}

	offset.Offset = o

	return offset, nil
}

func (o Offset) GetOffset() int {
	return o.Offset
}
