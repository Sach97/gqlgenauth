package middleware

import (
	"context"
	"net/http"

	"github.com/Sach97/gqlgenauth/auth/builder"
	auth "github.com/Sach97/gqlgenauth/auth/jwt"
	"github.com/go-chi/jwtauth"
)

type Chi struct {
	AuthService *auth.AuthService
}

func (c *Chi) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		tokenString := jwtauth.TokenFromHeader(req) //TODO: remove this depedency
		token, err := c.AuthService.ValidateJWT(tokenString, &builder.CustomClaims{})
		if err != nil {
			ctx = context.WithValue(ctx, "error", err) //TODO: refactor this
		}
		if token != nil {
			ctx = context.WithValue(ctx, "claims", token.Claims)
		}
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func (c *Chi) GetUserID(ctx context.Context) string {
	claims := ctx.Value("claims").(*builder.CustomClaims)
	sub := claims.StandardClaims.Subject

	return sub
}

//usage
//routerStrategy := middleware.Startegy{Chi{}}
//r := chi.NewRouter()
//r.Use(routerStrategy.AuthMiddleware)
// routerStrategy.GetUserID()
