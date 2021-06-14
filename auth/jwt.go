package auth

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtAuthenticator struct {
}

func NewJwtAuthenticator() JwtAuthenticator {
	return JwtAuthenticator{}
}

func (a JwtAuthenticator) Authentiicate(c *gin.Context) (SecurityToken, error) {
	t, err := GetJwtToken(c)

	if err != nil {
		log.Printf("Received arro authenticating JWT %v", err)
		return t, errors.New("Could not validate JWT token")
	}

	// Mark token as Authenticated
	log.Printf("Sucessfully authenticated request")
	t.Authenticated = true

	return t, nil
}

func GetJwtToken(c *gin.Context) (Token, error) {
	jWtToken := NewToken()

	// Get Public Key path
	publicKeyPath := os.Getenv("PUBLIC_KEY_PATH")
	log.Printf("Env public key path configured to: %s", publicKeyPath)
	if publicKeyPath == "" {
		publicKeyPath = "jwt/public.pem"
	}
	log.Printf("Setting public key path as: %s", publicKeyPath)

	// Get authorisation header for this request
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return jWtToken, errors.New("Authorization header not present")
	}

	if strings.HasPrefix(authHeader, "Bearer") {
		log.Printf("Got Bearer token")

		// Load Public key
		log.Printf("Loading public key")
		keyData, _ := ioutil.ReadFile(publicKeyPath)
		key, _ := jwt.ParseRSAPublicKeyFromPEM(keyData)

		// get JWT token sting from header
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		// Set raw token
		jWtToken.RawToken = tokenString

		// convert jwt string into token object (this will also validate token)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return key, nil
		})
		log.Printf("Successfully parsed jwt token: %v", token)

		if err != nil {
			return jWtToken, err
		}

		log.Printf("Setting claims for this token")
		claims := token.Claims.(jwt.MapClaims)

		// Set username
		if claims["username"] != "" {
			jWtToken.Username = claims["username"].(string)
			log.Printf("Got username claim: %s", jWtToken.Username)
		}

		// Set username
		if claims["organisation"] != "" {
			jWtToken.Organisation = claims["organisation"].(string)
			log.Printf("Got organisation claim: %s", jWtToken.Organisation)
		}
	}

	return jWtToken, nil
}
