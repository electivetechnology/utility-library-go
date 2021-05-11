package auth

import "github.com/gin-gonic/gin"

type Authenticator interface {
	Authentiicate(*gin.Context) (SecurityToken, error)
}
