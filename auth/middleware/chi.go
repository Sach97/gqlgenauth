package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/Sach97/gqlgenauth/auth/jwt"
	"github.com/Sach97/gqlgenauth/auth/user"
	"github.com/go-chi/jwtauth"
)

type Chi struct {
	AuthService *jwt.AuthService
}

func (c Chi) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenString := jwtauth.TokenFromHeader(req)

		token, err := c.AuthService.ValidateJWT(tokenString, &user.MyCustomClaims{})
		if err != nil || !token.Valid { //
			fmt.Println("we are here")
			fmt.Errorf("Token is not valid", err)
		}
		ctx := req.Context()
		ctx = context.WithValue(ctx, "error", err) //TODO: solve this
		ctx = context.WithValue(ctx, "claims", token.Claims)
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (c Chi) GetUserID(ctx context.Context) (string, error) {
	claims := ctx.Value("claims").(*user.MyCustomClaims)
	sub := claims.StandardClaims.Subject
	userID, err := base64.StdEncoding.DecodeString(string(sub))
	return string(userID), err
}

//usage
//routerStrategy := middleware.Startegy{Chi{}}
//r := chi.NewRouter()
//r.Use(routerStrategy.AuthMiddleware)
// routerStrategy.GetUserID()
