package auth

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// secretKey is used to sign & validate jwt tokens
var secretKey string

// init func to get the secret key from env
func init() {
	secretKey = os.Getenv("BLOG_SECRET_KEY")
}

// CreateToken will create the signed jwt token with given user claims
func CreateToken(username string) (*string, error) {
	claims := jwt.MapClaims{}

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["authorized"] = true

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// ValidateToken will validate the given jwt token and responds with username
func ValidateToken(tokenStr string) (username string, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username, ok = claims["username"].(string)
		if !ok {
			return "", fmt.Errorf("couldn't find user claims in the token")
		}

		return username, nil
	}

	return "", fmt.Errorf("unable to fetch claims from the token")
}
