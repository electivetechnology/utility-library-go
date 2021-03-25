package request

import (
	"log"
	"net/url"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/gin-gonic/gin"
)

type Filters struct {
	Filters map[string]*data.Filter
}

func GetFilters(c *gin.Context) *Filters {
	// Get Query Map
	q := c.Request.URL.Query()
	qm := c.Request.URL.Query().Get("filters")

	filters := mapFilters(q)

	log.Printf("%v\n", qm)
	//log.Printf("%v\n", filters)
	return &Filters{filters}
}

func mapFilters(m url.Values) map[string]*data.Filter {
	filters := make(map[string]*data.Filter)

	for k, v := range m {
		filter := data.NewFilter()
		log.Printf("Key is: %v\n", k)
		for _, c := range v {
			log.Printf("Value is: %v\n", c)
		}

		filters[k] = filter
	}

	return filters
}
