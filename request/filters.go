package request

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

type Filters struct {
	Filters map[string]*Filter
}

type Filter struct {
	ID         string
	Parent     string
	Criteria   []string
	DataFilter *data.Filter
}

func GetFilters(c *gin.Context) []*data.Filter {
	// Get Query Map
	q := c.Request.URL.Query()

	filters := mapFilters(q)

	return filters
}

func mapFilters(m url.Values) []*data.Filter {
	var dataFilters []*data.Filter
	filters := make(map[string]Filter)
	filtersWithData := make(map[string]Filter)

	for k, v := range m {
		filter := Filter{}
		re := regexp.MustCompile(`\[[a-zA-Z1-9-]+\]`)
		ids := re.FindAllString(k, -1)

		if len(ids) > 0 {
			filter.ID = getSafeFilterName(ids[len(ids)-1])
		} else {
			filter.ID = getSafeFilterName("0")
		}

		if len(ids) > 1 {
			filter.Parent = getSafeFilterName(ids[len(ids)-2])
		}

		for _, c := range v {
			filter.Criteria = append(filter.Criteria, c)
		}

		// Turn Filter object into data.Filter
		dataFilter := FilterToDataFilter(filter)
		dataFilter.Filters = make(map[string]*data.Filter)

		// Add to dataFilters
		filter.DataFilter = dataFilter
		filtersWithData[filter.ID] = filter

		// Add Subquery if needed
		if strings.Contains(filter.ID, "sub") {
			filter.DataFilter.Subquery.IsEnabled = true
			filter.DataFilter.Subquery.Key = "engagement_id"
			filter.DataFilter.Subquery.Set = "connect-f7e5b.staging_reporting.transcripts"
		}

		// Add filter to list of filters
		filters[filter.ID] = filter
	}

	// Iterate over each filter and assign parent
	for _, f := range filters {
		if f.Parent != "" {
			var a = data.Filter{}
			a = *filters[f.ID].DataFilter
			filtersWithData[f.Parent].DataFilter.Filters[f.ID] = &a
		}
	}

	for _, df := range filtersWithData {
		if len(df.Parent) == 0 {
			dataFilters = append(dataFilters, df.DataFilter)
		}
	}

	return dataFilters
}

func FilterToDataFilter(f Filter) *data.Filter {
	filter := data.NewFilter()

	// Turn each Criteria into data.Crierion
	for i, c := range f.Criteria {
		criterion, err := CriteriaToCriterion(c, i)

		if err != nil {
			log.Errorf("Could not parse criteria into criterion: %s", err)
		}

		filter.Criterions = append(filter.Criterions, criterion)
	}

	return filter
}

func CriteriaToCriterion(criteria string, index int) (data.Criterion, error) {
	c := data.Criterion{}

	// Check all parts of criteria exist
	parts := strings.Split(criteria, "-")

	// Check all elements of criteria are present
	if len(parts) != 5 ||
		(parts[0] != data.CRITERION_LOGIC_INTERSCTION && parts[0] != data.CRITERION_LOGIC_UNION) {
		return c, errors.New(
			"value for index " + strconv.Itoa(index) +
				", and part 1 (logic) must be one of " +
				"(or, and)" +
				", and part 3 (operand) must be one of " +
				"bool, eq, ne, nei, eqi")
	}

	// Check criteria type is correct
	if parts[3] != data.CRITERION_TYPE_VALUE &&
		parts[3] != data.CRITERION_TYPE_FIELD {
		return c, errors.New(
			"value for index " + strconv.Itoa(index) +
				", part 4 (type) must be one of (field or value)")
	}

	// Set Criterion details
	c.Logic = parts[0]
	c.Key = parts[1]
	c.Operand = parts[2]
	c.Type = parts[3]
	c.Value = parts[4]

	return c, nil
}

func getSafeFilterName(key string) string {
	key = strings.ReplaceAll(key, "[", "")
	key = strings.ReplaceAll(key, "]", "")

	return "f" + key
}
