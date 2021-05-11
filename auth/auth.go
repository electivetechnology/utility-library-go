package auth

import (
	"net/http"

	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/gin-gonic/gin"
)

var log logger.Logging
var authenticators map[string]Authenticator

func init() {
	// Add generic logger
	log = logger.NewLogger("auth")

	// Create holder for authenticators
	authenticators = make(map[string]Authenticator)

	// Add JWT Authenticator
	jwtAuth := NewJwtAuthenticator()
	authenticators["jwt"] = jwtAuth
}

func IsAuthenticated() gin.HandlerFunc {
	log.Printf("Checking if request is authenticated")

	return func(c *gin.Context) {
		st, _ := c.Get("SecurityToken")

		token := st.(SecurityToken)

		if token.IsAuthenticated() == false {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func CheckAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Checking authentication details for this request")
		token := NewToken()

		for name, authenticator := range authenticators {
			log.Printf("Trying authentication using %s as authenticator", name)
			t, err := authenticator.Authentiicate(c)

			if err == nil {
				token = t.(Token)
			}
		}

		c.Set("SecurityToken", token)
	}
}
