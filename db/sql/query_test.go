package sql

import (
	"testing"

	"github.com/electivetechnology/utility-library-go/data"
)

func TestPrepareNew(t *testing.T) {
	q := NewQuery("candidates")
	expected := "SELECT * FROM candidates LIMIT 1000 OFFSET 0"
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewBigQuery(t *testing.T) {
	q := NewQuery("`connect-f7e5b.staging_reporting.candidates`")
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` LIMIT 1000 OFFSET 0"
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimple(t *testing.T) {
	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates LIMIT 1000 OFFSET 0"
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimpleBigQuery(t *testing.T) {
	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` LIMIT 1000 OFFSET 0"
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimpleWithSort(t *testing.T) {
	// Prepare sorts
	sorts := make(map[string]data.Sort)

	// First Sort
	s1 := data.NewSort()
	s1.Field = "id"
	sorts["s_00"] = *s1

	// Second Sort
	s2 := data.NewSort()
	s2.Field = "email"
	s2.Direction = data.SORT_DIRECTION_DESC
	sorts["s_z"] = *s2

	// Third Sort
	s3 := data.NewSort()
	s3.Field = "first_name"
	sorts["s_a"] = *s3

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates ORDER BY `id` ASC, `first_name` ASC, `email` DESC LIMIT 1000 OFFSET 0"

	// Add sorts
	q.Sorts = sorts

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimpleBigQueryWithSort(t *testing.T) {
	// Prepare sorts
	sorts := make(map[string]data.Sort)

	// First Sort
	s1 := data.NewSort()
	s1.Field = "id"
	sorts["s_00"] = *s1

	// Second Sort
	s2 := data.NewSort()
	s2.Field = "email"
	s2.Direction = data.SORT_DIRECTION_DESC
	sorts["s_z"] = *s2

	// Third Sort
	s3 := data.NewSort()
	s3.Field = "first_name"
	sorts["s_a"] = *s3

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` ORDER BY `id` ASC, `first_name` ASC, `email` DESC LIMIT 1000 OFFSET 0"

	// Add sorts
	q.Sorts = sorts

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}
}

func TestPrepareNewSimpleWithFilter(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "first_name",
		Operand: "eq",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "organisation",
		Operand: "eq",
		Type:    "value",
		Value:   "Ds7q0eBi2Iyy",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "gt",
		Type:    "value",
		Value:   "2",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates WHERE (`first_name` =  CAST(:f_0_w_filter_0 AS CHAR) COLLATE utf8mb4_bin AND `organisation` =  CAST(:f_0_w_filter_1 AS CHAR) COLLATE utf8mb4_bin AND CAST(`id` AS NUMERIC) > CAST(:f_0_w_filter_2 AS NUMERIC)) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM candidates WHERE (`first_name` =  CAST(\"Kris\" AS CHAR) COLLATE utf8mb4_bin AND `organisation` =  CAST(\"Ds7q0eBi2Iyy\" AS CHAR) COLLATE utf8mb4_bin AND CAST(`id` AS NUMERIC) > CAST(\"2\" AS NUMERIC)) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleBigQueryWithFilter(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "firstName",
		Operand: "eq",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "organisation",
		Operand: "eq",
		Type:    "value",
		Value:   "Ds7q0eBi2Iyy",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "gt",
		Type:    "value",
		Value:   "2",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`firstName` AS STRING) =  CAST(:f_0_w_filter_0 AS STRING) AND CAST(`organisation` AS STRING) =  CAST(:f_0_w_filter_1 AS STRING) AND CAST(`id` AS NUMERIC) > CAST(:f_0_w_filter_2 AS NUMERIC)) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`firstName` AS STRING) =  CAST(\"Kris\" AS STRING) AND CAST(`organisation` AS STRING) =  CAST(\"Ds7q0eBi2Iyy\" AS STRING) AND CAST(`id` AS NUMERIC) > CAST(\"2\" AS NUMERIC)) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestGetSelectSqlSimple(t *testing.T) {
	q := NewQuery("candidates")

	selectClause := GetSelectSql(&q)

	expected := "SELECT *"

	if selectClause.Statement != expected {
		t.Errorf("GetSelectSql(*Query) failed, expected %v, got %v", expected, selectClause.Statement)
	}
}

func TestGetSelectSqlWithDisplays(t *testing.T) {
	q := NewQuery("candidates")

	fieldMap := make(map[string]string)
	fieldMap["*"] = "candidates"
	fieldMap["id"] = "candidates"
	fieldMap["candidates.lastName"] = "candidates"
	fieldMap["candidate.firstName"] = "candidates"

	d1 := data.NewDisplay()
	d1.Field = "id"

	d2 := data.NewDisplay()
	d2.Field = "candidates.lastName"

	d3 := data.NewDisplay()
	d3.Field = "candidate.firstName"
	d3.Alias = "fName"

	ds3 := make(map[string]data.Display)
	ds3["d1"] = *d1
	ds3["d2"] = *d2
	ds3["d3"] = *d3

	expected := "SELECT `candidates`.`firstName` AS fName, `candidates`.`id`, `candidates`.`lastName` FROM candidates LIMIT 1000 OFFSET 0"

	q.Displays = ds3
	q.FieldMap = fieldMap

	// Prepare query
	q.Prepare()

	if q.GetSql() != expected {
		t.Errorf("GetSelectSql(*Query) failed, expected %v, got %v", expected, q.GetSql())
	}
}

func TestPrepareNewSimpleWithFilterRelative(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "gt",
		Type:    "value",
		Value:   "2",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "ge",
		Type:    "value",
		Value:   "2",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "lt",
		Type:    "value",
		Value:   "2",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "id",
		Operand: "le",
		Type:    "value",
		Value:   "2",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates WHERE (`id` > :f_0_w_filter_0 AND `id` >= :f_0_w_filter_1 AND `id` < :f_0_w_filter_2 AND `id` <= :f_0_w_filter_3) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM candidates WHERE (`id` > \"2\" AND `id` >= \"2\" AND `id` < \"2\" AND `id` <= \"2\") LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleBigQueryWithFilterRelative(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "mid",
		Operand: "gt",
		Type:    "value",
		Value:   "3449302388198566",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "mid",
		Operand: "ge",
		Type:    "value",
		Value:   "3449302388198566",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "mid",
		Operand: "lt",
		Type:    "value",
		Value:   "3449302388198566",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "mid",
		Operand: "le",
		Type:    "value",
		Value:   "3449302388198566",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`mid` AS NUMERIC) > CAST(:f_0_w_filter_0 AS NUMERIC) AND CAST(`mid` AS NUMERIC) >= CAST(:f_0_w_filter_1 AS NUMERIC) AND CAST(`mid` AS NUMERIC) < CAST(:f_0_w_filter_2 AS NUMERIC) AND CAST(`mid` AS NUMERIC) <= CAST(:f_0_w_filter_3 AS NUMERIC)) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`mid` AS NUMERIC) > CAST(\"3449302388198566\" AS NUMERIC) AND CAST(`mid` AS NUMERIC) >= CAST(\"3449302388198566\" AS NUMERIC) AND CAST(`mid` AS NUMERIC) < CAST(\"3449302388198566\" AS NUMERIC) AND CAST(`mid` AS NUMERIC) <= CAST(\"3449302388198566\" AS NUMERIC)) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleWithFilterContains(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "inc",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "inci",
		Type:    "value",
		Value:   "Kris",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "ninc",
		Type:    "value",
		Value:   "Kris",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "ninci",
		Type:    "value",
		Value:   "Kris",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates WHERE (`email` LIKE  CONCAT(\"%\", CAST(:f_0_w_filter_0 AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` LIKE  CONCAT(\"%\", CAST(:f_0_w_filter_1 AS CHAR), \"%\") COLLATE utf8mb4_general_ci AND `email` NOT LIKE  CONCAT(\"%\", CAST(:f_0_w_filter_2 AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` NOT LIKE  CONCAT(\"%\", CAST(:f_0_w_filter_3 AS CHAR), \"%\") COLLATE utf8mb4_general_ci) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM candidates WHERE (`email` LIKE  CONCAT(\"%\", CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` LIKE  CONCAT(\"%\", CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_general_ci AND `email` NOT LIKE  CONCAT(\"%\", CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` NOT LIKE  CONCAT(\"%\", CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_general_ci) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleBigQueryWithFilterContains(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "inc",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "inci",
		Type:    "value",
		Value:   "Kris",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "ninc",
		Type:    "value",
		Value:   "Kris",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "ninci",
		Type:    "value",
		Value:   "Kris",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`email` AS STRING) LIKE  CONCAT(\"%\", CAST(:f_0_w_filter_0 AS STRING), \"%\") AND LOWER(CAST(`email` AS STRING)) LIKE  CONCAT(\"%\", LOWER(CAST(:f_0_w_filter_1 AS STRING)), \"%\") AND CAST(`email` AS STRING) NOT LIKE  CONCAT(\"%\", CAST(:f_0_w_filter_2 AS STRING), \"%\") AND LOWER(CAST(`email` AS STRING)) NOT LIKE  CONCAT(\"%\", LOWER(CAST(:f_0_w_filter_3 AS STRING)), \"%\")) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`email` AS STRING) LIKE  CONCAT(\"%\", CAST(\"Kris\" AS STRING), \"%\") AND LOWER(CAST(`email` AS STRING)) LIKE  CONCAT(\"%\", LOWER(CAST(\"Kris\" AS STRING)), \"%\") AND CAST(`email` AS STRING) NOT LIKE  CONCAT(\"%\", CAST(\"Kris\" AS STRING), \"%\") AND LOWER(CAST(`email` AS STRING)) NOT LIKE  CONCAT(\"%\", LOWER(CAST(\"Kris\" AS STRING)), \"%\")) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleWithFilterBegins(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "begins",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "nbegins",
		Type:    "value",
		Value:   "Kris",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "beginsi",
		Type:    "value",
		Value:   "Kris",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "nbeginsi",
		Type:    "value",
		Value:   "Kris",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates WHERE (`email` LIKE  CONCAT(CAST(:f_0_w_filter_0 AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` NOT LIKE  CONCAT(CAST(:f_0_w_filter_1 AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` LIKE  CONCAT(CAST(:f_0_w_filter_2 AS CHAR), \"%\") COLLATE utf8mb4_general_ci AND `email` NOT LIKE  CONCAT(CAST(:f_0_w_filter_3 AS CHAR), \"%\") COLLATE utf8mb4_general_ci) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM candidates WHERE (`email` LIKE  CONCAT(CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` NOT LIKE  CONCAT(CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_bin AND `email` LIKE  CONCAT(CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_general_ci AND `email` NOT LIKE  CONCAT(CAST(\"Kris\" AS CHAR), \"%\") COLLATE utf8mb4_general_ci) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleBigQueryWithFilterBegins(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "begins",
		Type:    "value",
		Value:   "Kris",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "nbegins",
		Type:    "value",
		Value:   "Kris",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "beginsi",
		Type:    "value",
		Value:   "Kris",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "email",
		Operand: "nbeginsi",
		Type:    "value",
		Value:   "Kris",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`email` AS STRING) LIKE  CONCAT(CAST(:f_0_w_filter_0 AS STRING), \"%\") AND CAST(`email` AS STRING) NOT LIKE  CONCAT(CAST(:f_0_w_filter_1 AS STRING), \"%\") AND LOWER(CAST(`email` AS STRING)) LIKE  CONCAT(LOWER(CAST(:f_0_w_filter_2 AS STRING)), \"%\") AND LOWER(CAST(`email` AS STRING)) NOT LIKE  CONCAT(LOWER(CAST(:f_0_w_filter_3 AS STRING)), \"%\")) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (CAST(`email` AS STRING) LIKE  CONCAT(CAST(\"Kris\" AS STRING), \"%\") AND CAST(`email` AS STRING) NOT LIKE  CONCAT(CAST(\"Kris\" AS STRING), \"%\") AND LOWER(CAST(`email` AS STRING)) LIKE  CONCAT(LOWER(CAST(\"Kris\" AS STRING)), \"%\") AND LOWER(CAST(`email` AS STRING)) NOT LIKE  CONCAT(LOWER(CAST(\"Kris\" AS STRING)), \"%\")) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleWithFilterRegex(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "organisation",
		Operand: "re",
		Type:    "value",
		Value:   "^c",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM candidates")
	expected := "SELECT * FROM candidates WHERE (`organisation` REGEXP  :f_0_w_filter_0) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM candidates WHERE (`organisation` REGEXP  \"^c\") LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleBigQueryWithFilterRegex(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "organisation",
		Operand: "re",
		Type:    "value",
		Value:   "^c",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.candidates`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (REGEXP_CONTAINS(CAST(`organisation` AS STRING),  CAST(:f_0_w_filter_0 AS STRING))) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM `connect-f7e5b.staging_reporting.candidates` WHERE (REGEXP_CONTAINS(CAST(`organisation` AS STRING),  CAST(\"^c\" AS STRING))) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleWithFilterIn(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "in",
		Type:    "value",
		Value:   "3pGyisgS86yc,ImrJLEC21t2i",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "nin",
		Type:    "value",
		Value:   "3pGyisgS86yc,ImrJLEC21t2i",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "ini",
		Type:    "value",
		Value:   "3pGyisgS86yc,ImrJLEC21t2i",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "nini",
		Type:    "value",
		Value:   "3pGyisgS86yc,ImrJLEC21t2i",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM transcripts")
	expected := "SELECT * FROM transcripts WHERE (`invitation_id` IN (:f_0_w_filter_0_0 COLLATE utf8mb4_bin, :f_0_w_filter_0_1 COLLATE utf8mb4_bin) AND `invitation_id` NOT IN (:f_0_w_filter_1_0 COLLATE utf8mb4_bin, :f_0_w_filter_1_1 COLLATE utf8mb4_bin) AND `invitation_id` IN (:f_0_w_filter_2_0 COLLATE utf8mb4_general_ci, :f_0_w_filter_2_1 COLLATE utf8mb4_general_ci) AND `invitation_id` NOT IN (:f_0_w_filter_3_0 COLLATE utf8mb4_general_ci, :f_0_w_filter_3_1 COLLATE utf8mb4_general_ci)) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM transcripts WHERE (`invitation_id` IN (\"3pGyisgS86yc\" COLLATE utf8mb4_bin, \"ImrJLEC21t2i\" COLLATE utf8mb4_bin) AND `invitation_id` NOT IN (\"3pGyisgS86yc\" COLLATE utf8mb4_bin, \"ImrJLEC21t2i\" COLLATE utf8mb4_bin) AND `invitation_id` IN (\"3pGyisgS86yc\" COLLATE utf8mb4_general_ci, \"ImrJLEC21t2i\" COLLATE utf8mb4_general_ci) AND `invitation_id` NOT IN (\"3pGyisgS86yc\" COLLATE utf8mb4_general_ci, \"ImrJLEC21t2i\" COLLATE utf8mb4_general_ci)) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}

func TestPrepareNewSimpleBigQueryWithFilterIn(t *testing.T) {
	// Prepare filters
	filters := make(map[string]data.Filter)

	c1 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "in",
		Type:    "value",
		Value:   "bXkKeWW4hTZF,j1uLpVoLaals",
	}

	c2 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "nin",
		Type:    "value",
		Value:   "bXkKeWW4hTZF,j1uLpVoLaals",
	}

	c3 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "ini",
		Type:    "value",
		Value:   "bXkKeWW4hTZF,j1uLpVoLaals",
	}

	c4 := data.Criterion{
		Logic:   "and",
		Key:     "invitation_id",
		Operand: "nini",
		Type:    "value",
		Value:   "bXkKeWW4hTZF,j1uLpVoLaals",
	}

	// First Filter
	f1 := data.NewFilter()
	f1.Criterions = append(f1.Criterions, c1)
	f1.Criterions = append(f1.Criterions, c2)
	f1.Criterions = append(f1.Criterions, c3)
	f1.Criterions = append(f1.Criterions, c4)

	filters["f_0"] = *f1

	q, _ := NewSimpleQuery("SELECT * FROM `connect-f7e5b.staging_reporting.transcripts`")
	q.Flavour = QUERY_FLAVOUR_BIG_QUERY
	expected := "SELECT * FROM `connect-f7e5b.staging_reporting.transcripts` WHERE (CAST(`invitation_id` AS STRING) IN (:f_0_w_filter_0_0, :f_0_w_filter_0_1) AND CAST(`invitation_id` AS STRING) NOT IN (:f_0_w_filter_1_0, :f_0_w_filter_1_1) AND LOWER(CAST(`invitation_id` AS STRING)) IN (LOWER(:f_0_w_filter_2_0), LOWER(:f_0_w_filter_2_1)) AND LOWER(CAST(`invitation_id` AS STRING)) NOT IN (LOWER(:f_0_w_filter_3_0), LOWER(:f_0_w_filter_3_1))) LIMIT 1000 OFFSET 0"

	// Add filters
	q.Filters = filters

	// Prepare query
	q.Prepare()

	if q.Statement != expected {
		t.Errorf("Query.Prepare() failed, expected %v, got %v", expected, q.Statement)
	}

	expectedSql := "SELECT * FROM `connect-f7e5b.staging_reporting.transcripts` WHERE (CAST(`invitation_id` AS STRING) IN (\"bXkKeWW4hTZF\", \"j1uLpVoLaals\") AND CAST(`invitation_id` AS STRING) NOT IN (\"bXkKeWW4hTZF\", \"j1uLpVoLaals\") AND LOWER(CAST(`invitation_id` AS STRING)) IN (LOWER(\"bXkKeWW4hTZF\"), LOWER(\"j1uLpVoLaals\")) AND LOWER(CAST(`invitation_id` AS STRING)) NOT IN (LOWER(\"bXkKeWW4hTZF\"), LOWER(\"j1uLpVoLaals\"))) LIMIT 1000 OFFSET 0"

	if q.GetSql() != expectedSql {
		t.Errorf("Query.GetSql() failed, expected %v, got %v", expectedSql, q.GetSql())
	}
}
