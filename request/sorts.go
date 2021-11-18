package request

// DEPRECATED - Please use new request/context package insted
import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
)

type Sorts struct {
	Sorts map[string]*Sort
}

type Sort struct {
	ID         string
	Directives []string
	DataSort   *data.Sort
}

func GetSorts(c *gin.Context) []*data.Sort {
	// Get Query Map
	q := c.Request.URL.RawQuery

	sorts := GetSortsFromQueryString(q)

	return sorts
}

func GetSortsFromQueryString(query string) []*data.Sort {
	// Add query string if required
	if !strings.HasPrefix(query, "?") {
		query = "?" + query
	}

	// Parse URL
	u, err := url.Parse(query)
	if err != nil {
		log.Debugf("Could not parse query: %v\n", err)
	}

	// Extract values
	q := u.Query()

	// Map sorts
	sorts := mapSorts(q)

	return sorts
}

func mapSorts(m url.Values) []*data.Sort {
	var dataSorts []*data.Sort

	for k, directives := range m {
		sort := Sort{}

		re := regexp.MustCompile(`\[[a-zA-Z1-9-]+\]`)
		ids := re.FindAllString(k, -1)

		if len(ids) > 0 {
			sort.ID = getSafeFilterName(ids[len(ids)-1])
		} else {
			sort.ID = getSafeFilterName("0")
		}

		for i, directive := range directives {
			// Turn each Sort Directive object into data.Sort
			dataSort, err := DirectiveToDataSort(directive, i)

			if err != nil {
				log.Errorf("Could not parse sort directive into sort: %s", err)
			}

			// Add sort to list of sorts
			dataSorts = append(dataSorts, dataSort)
		}
	}

	return dataSorts
}

func getSafeSortName(key string) string {
	key = strings.ReplaceAll(key, "[", "")
	key = strings.ReplaceAll(key, "]", "")

	return "s" + key
}

func DirectiveToDataSort(directive string, index int) (*data.Sort, error) {
	sort := data.NewSort()

	// Check all parts of directive exist
	parts := strings.Split(directive, "-")

	// Check all elements of directive are present
	if len(parts) != 2 ||
		(parts[1] != data.SORT_DIRECTION_ASC && parts[1] != data.SORT_DIRECTION_DESC) {
		return sort, errors.New(
			"value for index " + strconv.Itoa(index) +
				" must be string in format of " +
				" {field}-{direction}" +
				" Example: id-asc, name-desc")
	}

	// Set Sort details
	sort.Field = parts[0]
	sort.Direction = parts[1]

	return sort, nil
}
