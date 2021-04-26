package router

import (
	"testing"
)

type TestRegisterRouteItem struct {
	route Route
}

func TestRegisterRoute(t *testing.T) {
	r1 := Route{}
	r1.Method = []string{"HEAD", "OPTIONS"}
	r1.Path = "/v1/status"
	r1.Handler = NoContent

	r2 := Route{}
	r2.Method = []string{"GET"}
	r2.Path = "/v1/foo"
	r2.Handler = NoContent

	testData := []TestRegisterRouteItem{
		{r1},
	}

	items := 0
	for _, item := range testData {
		RegisterRoute(item.route)
		items++
	}

	if len(routes) != items {
		t.Errorf("RegisterRoute() failed, expected number of registered endpoints to be %d, got %d", items, len(routes))
	}
}
