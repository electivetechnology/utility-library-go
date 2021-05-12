package router

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PORT = 80
)

type Router struct {
	Logger logger.Logging
	Engine *gin.Engine
	Port   int // Port to run engin on (defaults to 80)
}

var log logger.Logging

// NewRouter returns a new blank Router instance.
func NewRouter() *Router {
	// Add generic logger
	log = logger.NewLogger("router")

	// Assign default port
	port := DEFAULT_PORT

	// Check if port has been reconfigured by ENV
	p := os.Getenv("ROUTER_PORT")
	log.Printf("Read router port from env: %s", p)
	if p != "" {
		port, _ = strconv.Atoi(p)
	}
	log.Printf("Router configured to run on port %d", port)

	// Setup Router
	r := Router{
		Logger: log,
		Engine: SetupEngine(),
		Port:   port,
	}

	return &r
}

func (r *Router) Run() {
	r.Logger.Fatalf("%v", r.Engine.Run(":"+strconv.Itoa(r.Port)))
}
