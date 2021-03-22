package request

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Filters struct {
}

type Filter struct {
	Criterions []Criterion
	Logic      []string
}

type Criterion struct {
	Logic   string
	Key     string
	Operand string
	Type    string
	Value   string
}

func GetFilters(c *gin.Context) *Filters {
	// Get Query Map
	q := c.Request.URL.Query()
	log.Printf("%v\n", q)
	return &Filters{}
}
