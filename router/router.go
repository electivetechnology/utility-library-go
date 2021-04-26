package router

import (
	"strconv"

	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/gin-gonic/gin"
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

	// Setup Router
	r := Router{
		Logger: log,
		Engine: SetupEngine(),
		Port:   80,
	}

	return &r
}

func (r *Router) Run() {
	r.Logger.Fatalf("%v", r.Engine.Run(":"+strconv.Itoa(r.Port)))
}
