package jwt

import (
	"fmt"
	"time"

	"github.com/Sach97/gqlgenauth/auth/context"
	jwt "github.com/dgrijalva/jwt-go"
)

//AuthService holds our auth struct
type AuthService struct {
	appName             *string
	signedSecret        *string
	expiredTimeInSecond *time.Duration
}

//NewAuthService instantiates a new AuthService
func NewAuthService(config *context.Config) *AuthService {
	return &AuthService{&config.AppName, &config.JWTSecret, &config.JWTExpireIn}
}

//SignJWT signs a new jwt
func (a *AuthService) SignJWT(customMapClaims interface{}) (*string, error) {
	mapClaims, _ := customMapClaims.(jwt.MapClaims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString([]byte(*a.signedSecret))
	return &tokenString, err
}

//ValidateJWT validates a new jwt
func (a *AuthService) ValidateJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("	unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(*a.signedSecret), nil
	})
	return token, err
}