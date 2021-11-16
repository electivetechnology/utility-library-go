package context

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/data"
	SortType "github.com/electivetechnology/utility-library-go/data/types/pseudo/sorts"
	"github.com/electivetechnology/utility-library-go/validation"
	"github.com/gin-gonic/gin"
)

type Sort struct {
	ID        string
	Directive string
	DataSort  *data.Sort
}

type Sorts struct {
	Sorts map[string]Sort
}

type SortsType interface {
	GetSorts() map[string]Sort
	GetSort(key string) (Sort, error)
	GetDataSorts() map[string]data.Sort
}

func NewSorts() Sorts {
	sorts := Sorts{}
	sorts.Sorts = make(map[string]Sort)

	return sorts
}

func GetSorts(c *gin.Context, requirements validation.ValidatorRequirements) (Sorts, error) {
	sorts, err := GetSortsFromContext(c)
	if err != nil {
		log.Fatalf(err.Error())
		return sorts, err
	}

	for idx, sort := range sorts.GetDataSorts() {
		_, err := ValidateSort(sort, idx, requirements)
		if err != nil {
			log.Fatalf(err.Error())
			return sorts, err
		}
	}

	return sorts, nil
}

func ValidateSort(sort data.Sort, namespace string, requirements validation.ValidatorRequirements) (data.Sort, error) {
	// Check if sort is valid
	_, err := SortType.Check(sort, requirements)
	if err != nil {
		msg := fmt.Sprintf("Sort %s failed validation with message: %s", namespace, err.Error())
		log.Fatalf(msg)
		return sort, errors.New(msg)
	}

	return sort, nil
}

// GetSortsFromContext returns Sorts passed via request sorts[] parameter
// Order of the sorts matters. Anonymous sorts are always processed first,
// folowed by named sorts in alpha-numerical order.
// A query string sorts[z]=name-asc&sorts[a]=email-asc&sorts[]=id-asc will result in following sorts:
// s_00:id-asc, s_a:email-asc, s_z:name-asc
func GetSortsFromContext(ctx *gin.Context) (Sorts, error) {
	sorts := NewSorts()

	// Get anonymous
	anonymous, err := GetAnonymousSortsFromContext(ctx)
	if err != nil {
		log.Fatalf(err.Error())
		return sorts, err
	}

	for key, sort := range anonymous.Sorts {
		sorts.Sorts[key] = sort
	}

	// Get named sorts
	named, err := GetMappedSortsFromContext(ctx)
	if err != nil {
		log.Fatalf(err.Error())
		return sorts, err
	}

	for key, sort := range named.Sorts {
		sorts.Sorts[key] = sort
	}

	return sorts, nil
}

func GetAnonymousSortsFromContext(ctx *gin.Context) (Sorts, error) {
	s, _ := ctx.GetQueryArray("sorts[]")
	sorts := NewSorts()

	for idx, directive := range s {
		sort := Sort{}
		sort.ID = getSafeSortName("0" + strconv.Itoa(idx))
		sort.Directive = directive
		dataSort, err := DirectiveToDataSort(directive, idx, sort.ID)
		// Check for errors
		if err != nil {
			log.Fatalf(err.Error())
			return sorts, err
		}

		// Assign values
		sort.DataSort = dataSort
		sorts.Sorts[sort.ID] = sort
	}

	return sorts, nil
}

func GetMappedSortsFromContext(ctx *gin.Context) (Sorts, error) {
	s, _ := ctx.GetQueryMap("sorts")
	sorts := NewSorts()
	i := 0
	for idx, directive := range s {
		sort := Sort{}
		sort.ID = getSafeSortName(idx)
		sort.Directive = directive
		dataSort, err := DirectiveToDataSort(directive, i, sort.ID)

		// Check for errors
		if err != nil {
			log.Fatalf(err.Error())
			return sorts, err
		}

		// Assign values
		sort.DataSort = dataSort
		i++
		sorts.Sorts[sort.ID] = sort
	}

	return sorts, nil
}

func getSafeSortName(key string) string {
	key = strings.ReplaceAll(key, "[", "")
	key = strings.ReplaceAll(key, "]", "")

	return "s_" + key
}

func DirectiveToDataSort(directive string, index int, name string) (*data.Sort, error) {
	sort := data.NewSort()
	var msg string

	// Check all parts of directive exist
	parts := strings.Split(directive, "-")

	// Check all elements of directive are present
	if len(parts) != 2 {
		msg = fmt.Sprintf(
			"value for sort %s index %d must be string in format of "+
				" {field}-{direction}"+
				" Example: id-asc, name-desc", name, index)
		log.Fatalf(msg)
		return sort, errors.New(msg)
	}

	if parts[1] != data.SORT_DIRECTION_ASC && parts[1] != data.SORT_DIRECTION_DESC {
		msg = fmt.Sprintf(
			"sort direction for sort %s index %d (%s) must be one of "+
				data.SORT_DIRECTION_ASC+
				" or "+data.SORT_DIRECTION_DESC+
				". '%s' given instead", name, index, parts[0], parts[1])
		log.Fatalf(msg)
		return sort, errors.New(msg)
	}

	// Set Sort details
	sort.Field = parts[0]
	sort.Direction = parts[1]

	return sort, nil
}

func (s Sorts) GetSorts() map[string]Sort {
	return s.Sorts
}

func (s Sorts) GetSort(key string) (Sort, error) {
	sort, found := s.Sorts[key]

	if found {
		return sort, nil
	}

	msg := fmt.Sprintf("sort with name %s does not exists", key)
	log.Fatalf(msg)

	return Sort{}, errors.New(msg)
}

func (s Sorts) GetDataSorts() map[string]data.Sort {
	sorts := make(map[string]data.Sort)
	for _, sort := range s.GetSorts() {
		sorts[sort.ID] = *sort.DataSort
	}

	return sorts
}
