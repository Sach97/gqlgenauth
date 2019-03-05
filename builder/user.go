package builder

import (
	"errors"
	"time"

	"github.com/Sach97/ninshoo/context"
	"github.com/Sach97/ninshoo/model"

	"github.com/dgrijalva/jwt-go"
)

type BuilderService struct {
	cfg *context.Config
}

//TODO: eventually connect to DB to get roles

func NewBuilderService(cfg *context.Config) *BuilderService {
	return &BuilderService{
		cfg: cfg,
	}
}

type CustomClaims struct {
	StandardClaims jwt.StandardClaims
	HasuraClaims   HasuraClaims `json:"https://hasura.io/jwt/claims"`
}

//hacky but I did'nt find better wsay to do this
func (c CustomClaims) Valid() error {
	return errors.New("")

}

func (b *BuilderService) BuildCustomClaims(user *model.User) *CustomClaims {
	now := time.Now()
	expires := now.Add(24 * time.Hour * 30) //TODO:get expiration from config
	hasuraClaims := HasuraClaimsBuilder.
		AddRole("user").
		DefaultRole("user").
		UserID(user.ID).
		Custom("custom-value").
		Build()
	//TODO: find a way to iterate over roles
	standardClaims := StandardClaimsBuilder.
		Subject(user.ID).
		ExpiresAt(expires.Unix()).
		Build()
	return &CustomClaims{StandardClaims: standardClaims, HasuraClaims: hasuraClaims}
}
