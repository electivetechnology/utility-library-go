package router

import (
	"os"

	"github.com/electivetechnology/utility-library-go/auth"
	"github.com/gin-gonic/gin"
)

// SetupEngine sets up router and returns its instance of the gin.Engine.
func SetupEngine(handler gin.HandlerFunc) *gin.Engine {
	// Force log's color
	gin.ForceConsoleColor()

	gin.SetMode(os.Getenv("GIN_MODE"))
	engine := gin.Default()
	engine.Use(handler)
	engine.Use(gin.Recovery())

	// Add Cors
	var isCorsEnabled bool
	if os.Getenv("ENABLE_CORS") == "true" {
		log.Printf("CORS enabled")
		isCorsEnabled = true
		engine.Use(AddCors())
	}

	// Use the checkAuthentication middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	engine.Use(auth.CheckAuthentication())

	// Load endpoints
	loadEndpoints(engine, routes, isCorsEnabled)

	return engine
}

func loadEndpoints(engine *gin.Engine, endpoints []Route, isCorsEnabled bool) *gin.Engine {
	log.Printf("Engine is loading endpoints..")

	// Load route collection
	for _, route := range endpoints {
		registerEndpoint(route, engine)

		// Add CORS [OPTIONS] version of this endpoint
		if isCorsEnabled {
			optionsRoute := route
			optionsRoute.Method = []string{"OPTIONS"}
			optionsRoute.Handler = NoContent
			log.Printf("Created OPTIONS copy of the route for CORS %s %s", optionsRoute.Method, optionsRoute.Path)
			registerEndpoint(optionsRoute, engine)
		}
	}

	log.Printf("Endpoints loaded successfully")

	return engine
}

func registerEndpoint(route Route, engine *gin.Engine) {
	methods := map[string]func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes{
		"GET":     engine.GET,
		"HEAD":    engine.HEAD,
		"OPTIONS": engine.OPTIONS,
		"PATCH":   engine.PATCH,
		"POST":    engine.POST,
		"PUT":     engine.PUT,
		"DELETE":  engine.DELETE,
	}

	for _, method := range route.Method {
		// Get routes that are already registered
		registeredRoutes := engine.Routes()
		if len(registeredRoutes) > 0 {
			var routeExist bool = false
			// Check if endpoint is already registered
			for _, r := range registeredRoutes {
				if r.Path == route.Path && r.Method == method {
					routeExist = true
				}
			}

			// Register endpoint
			if !routeExist {
				addRoute(route, methods[method])
			}
		} else {
			addRoute(route, methods[method])
		}
	}
}
