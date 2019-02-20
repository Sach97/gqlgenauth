package jwt

import (
	"errors"
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

type CustomClaims struct {
	StandardClaims jwt.StandardClaims
	CustomClaimsI  map[string]interface{}
}

//hacky but I did'nt find better wsay to do this
func (c CustomClaims) Valid() error {
	return errors.New("")

}

//SignJWT signs a new jwt
func (a *AuthService) SignJWT(customClaims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(*a.signedSecret))
	return tokenString, err
}

//ValidateJWT validates a token string
func (a *AuthService) ValidateJWT(tokenString string, customClaims *CustomClaims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("	unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*a.signedSecret), nil
	})
	return token, err
}
