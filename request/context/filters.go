package context

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	"github.com/gin-gonic/gin"
)

type Filter struct {
	ID         string
	Criteria   []string
	Parent     string
	DataFilter *data.Filter
}

type Filters struct {
	Filters map[string]Filter
}

type FiltersType interface {
	GetFilters() map[string]Filter
	GetDataFilters() map[string]data.Filter
}

func NewFilter() Filter {
	filter := Filter{}

	return filter
}

func NewFilters() Filters {
	filters := Filters{}
	filters.Filters = make(map[string]Filter)

	return filters
}

func GetFilters(c *gin.Context) Filters {
	filters := GetFiltersFromContext(c)

	return filters
}

func GetFiltersFromContext(ctx *gin.Context) Filters {
	filters := NewFilters()

	// Get anonymous
	anonymous := GetAnonymousFilterFromContext(ctx)
	filters.Filters[anonymous.ID] = anonymous

	// Get named filters
	named := GetMappedFiltersFromContext(ctx)
	for key, filter := range named.Filters {
		filters.Filters[key] = filter
	}

	return filters
}

func GetAnonymousFilterFromContext(ctx *gin.Context) Filter {
	f, _ := ctx.GetQueryArray("filters[]")

	filter := Filter{}
	filter.ID = getSafeFilterName("0")
	filter.Criteria = append(filter.Criteria, f...)
	filter.DataFilter, _ = FilterToDataFilter(filter)

	return filter
}

func GetMappedFiltersFromContext(ctx *gin.Context) Filters {
	filters := GetNamedFiltersFromQueryString(ctx)

	return filters
}

func GetNamedFiltersFromQueryString(ctx *gin.Context) Filters {
	filters := NewFilters()
	q := ctx.Request.URL.Query()
	filtersWithData := make(map[string]Filter)
	var isKeyDeleted bool

	for key, val := range q {
		filter := Filter{}
		isKeyDeleted = false
		re := regexp.MustCompile(`(filters)(\[[a-zA-Z1-9-]+\])+`)
		id := re.FindString(key)
		id = strings.ReplaceAll(id, "filters", "")
		if id == "" {
			isKeyDeleted = true
			delete(q, key)
		}

		if isKeyDeleted {
			continue
		}

		re = regexp.MustCompile(`\[[a-zA-Z1-9-]+\]`)
		ids := re.FindAllString(key, -1)

		// Work out name for each filter
		if len(ids) > 0 {
			filter.ID = getSafeFilterName(ids[len(ids)-1])
		} else {
			filter.ID = getSafeFilterName("0")
		}

		// Assign Parent filter (if exists)
		if len(ids) > 1 {
			filter.Parent = getSafeFilterName(ids[len(ids)-2])
		}

		// Add Criteria to each Filter
		filter.Criteria = append(filter.Criteria, val...)

		// Turn Filter object into data.Filter
		dataFilter, _ := FilterToDataFilter(filter)
		dataFilter.Filters = make(map[string]*data.Filter)

		// Add to dataFilters
		filter.DataFilter = dataFilter
		filtersWithData[filter.ID] = filter

		// Add filter to list of filters
		filters.Filters[filter.ID] = filter
	}

	return filters
}

func FilterToDataFilter(f Filter) (*data.Filter, error) {
	filter := data.NewFilter()

	// Turn each Directive (Criteria) into data.Crierion
	for i, c := range f.Criteria {
		criterion, err := CriteriaToCriterion(c, i, f.ID)

		if err != nil {
			msg := fmt.Sprintf("Could not parse criteria into criterion: %s", err)
			log.Fatalf(msg)
			return filter, errors.New(msg)
		}

		filter.Criterions = append(filter.Criterions, criterion)
	}

	log.Printf("Criteria To Filter: %v", filter)

	return filter, nil
}

func CriteriaToCriterion(criteria string, index int, name string) (data.Criterion, error) {
	c := data.Criterion{}
	var msg string

	// Check all parts of criteria exist
	parts := strings.Split(criteria, "-")

	// Check all elements of criteria are present
	if len(parts) != 5 {
		msg = fmt.Sprintf(
			"value for filter %s index %d must be string in format of "+
				" {logic}-{field}-{operand}-{type}-{value}"+
				" Example: and-name-eq-value-smith"+
				" Got '%s' instead", name, index, criteria)
		log.Fatalf(msg)
		return c, errors.New(msg)
	}

	if parts[0] != data.CRITERION_LOGIC_INTERSECTION && parts[0] != data.CRITERION_LOGIC_UNION {
		msg = fmt.Sprintf(
			"value for filter %s index %d"+
				", and part 1 (logic) must be one of "+strings.Join(data.GetLogicOptions(), ", ")+
				" Got '%s' instead", name, index, parts[0])

		log.Fatalf(msg)
		return c, errors.New(msg)
	}

	if !data.IsOperand(parts[2]) {
		msg = fmt.Sprintf(
			"value for filter %s index %d"+
				", and part 3 (operand) must be one of "+strings.Join(data.GetOperandOptions(), ", ")+
				" Got '%s' instead", name, index, parts[2])

		log.Fatalf(msg)
		return c, errors.New(msg)
	}

	// Check criteria type is correct
	if parts[3] != data.CRITERION_TYPE_VALUE &&
		parts[3] != data.CRITERION_TYPE_FIELD {
		msg = fmt.Sprintf(
			"value for filter %s index %d"+
				", part 4 (type) must be one of (field or value)"+
				" Got '%s' instead", name, index, parts[3])

		log.Fatalf(msg)
		return c, errors.New(msg)
	}

	// Set Criterion details
	c.Logic = parts[0]
	c.Key = parts[1]
	c.Operand = parts[2]
	c.Type = parts[3]
	c.Value = parts[4]

	log.Printf("Criterion: %v", c)

	return c, nil
}

func getSafeFilterName(key string) string {
	key = strings.ReplaceAll(key, "[", "")
	key = strings.ReplaceAll(key, "]", "")

	return "f_" + key
}

func (f Filters) GetFilters() map[string]Filter {
	return f.Filters
}

func (f Filters) GetDataFilters() map[string]data.Filter {
	filters := make(map[string]data.Filter)
	for _, filter := range f.GetFilters() {
		filters[filter.ID] = *filter.DataFilter
	}

	return filters
}
