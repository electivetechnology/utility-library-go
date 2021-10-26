package router

import "github.com/gin-gonic/gin"

var headers map[string]string

func AddDeaultHeaders() {
	log.Printf("Adding default CORS Headers")
	defaultHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, OPTIONS, HEAD, POST, PUT, PATCH, DELETE",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	}

	if headers == nil {
		headers = make(map[string]string)
	}

	for key, header := range defaultHeaders {
		log.Printf("Adding default Header %s", key)
		headers[key] = header
	}
}

func AddCorsHeader(key string, value string) {
	if headers == nil {
		headers = make(map[string]string)
	}

	headers[key] = value
}

func AddCors() gin.HandlerFunc {
	// Add Default CORS headers
	AddDeaultHeaders()

	return func(c *gin.Context) {
		for header, value := range headers {
			log.Printf("Adding CORS Header %s as %s", header, value)
			c.Writer.Header().Add(header, value)
		}
	}
}
