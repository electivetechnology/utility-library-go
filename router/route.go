package router

import (
	"github.com/electivetechnology/utility-library-go/auth"
	"github.com/gin-gonic/gin"
	"os"
)

type Route struct {
	Method          []string
	Path            string
	Handler         gin.HandlerFunc
	IsAuthenticated bool
	Authenticator   func(r Route) gin.HandlerFunc
}

// List of available routes/endpoints
// [Methods], Path, handler
var routes = []Route{}
var isEnabled = "true"

func init() {
	isEnabled = os.Getenv("ACL_CLIENT_ENABLED")
}

// RegisterRoute allows to add new Route object to list of engine endpoints
func RegisterRoute(route Route) {
	routes = append(routes, route)
}

func addRoute(route Route, f func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes) gin.IRoutes {
	log.Printf("Registering new endpoint: %s %s", route.Method, route.Path)

	if route.IsAuthenticated && isEnabled == "true" {
		log.Printf("Checking if request is authenticated")
		if route.Authenticator != nil {
			return f(route.Path, route.Authenticator(route), route.Handler)
		} else {
			return f(route.Path, auth.IsAuthenticated(), route.Handler)
		}
	}

	return f(route.Path, route.Handler)
}
