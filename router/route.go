package router

import "github.com/gin-gonic/gin"

type Route struct {
	Method  []string
	Path    string
	Handler gin.HandlerFunc
}

// List of available routes/endpoints
// [Methods], Path, handler (use isAuthorized(handler) middleware for authorisation)
var routes = []Route{}

// RegisterRoute allows to add new Route object to list of engine endpoints
func RegisterRoute(route Route) {
	routes = append(routes, route)
}

func addRoute(route Route, f func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes) gin.IRoutes {
	log.Printf("Registering new endpoint: %s %s", route.Method, route.Path)

	return f(route.Path, route.Handler)
}
