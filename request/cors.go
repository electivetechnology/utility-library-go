package request

import "github.com/gin-gonic/gin"

func GetCorsHeaders() map[string]string {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, OPTIONS, HEAD, POST, PUT",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	}

	return headers
}

func AddCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := GetCorsHeaders()

		for header, value := range headers {
			c.Writer.Header().Add(header, value)
		}
	}
}
