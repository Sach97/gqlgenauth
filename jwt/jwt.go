package jwt

import (
	"fmt"
	"time"

	"github.com/Sach97/ninshoo/builder"
	"github.com/Sach97/ninshoo/context"
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
	if config.JWTSecret == "" {
		panic("You must fill JWTSECRET env variable with your jwt secret ")
	}
	return &AuthService{&config.AppName, &config.JWTSecret, &config.JWTExpireIn}
}

//SignJWT signs a new jwt
func (a *AuthService) SignJWT(customClaims *builder.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(*a.signedSecret))
	return tokenString, err
}

//ValidateJWT validates a token string
func (a *AuthService) ValidateJWT(tokenString string, customClaims *builder.CustomClaims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(*a.signedSecret), nil
	})
	return token, err
}

//TODO: strategy pattern for other algorithms than HMAC
//TODO: strategy pattern stateful authentification
