package builder

import (
	"time"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/jwt"
	"github.com/Sach97/gqlgenauth/auth/model"
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

func (b *BuilderService) BuildCustomClaims(user *model.User) *jwt.CustomClaims {
	now := time.Now()
	expires := now.Add(24 * time.Hour * 30) //TODO:get expiration from config
	customClaims := HasuraClaimsBuilder.
		AddRole("editor").
		AddRole("user").
		DefaultRole("user").
		OrgID(user.ID).
		Custom("custom-value").
		Build()
	//TODO: find a way to iterate over roles
	claims := CustomClaimsBuilder.
		Subject(user.ID).
		ExpiresAt(expires.Unix()).
		Issuer("test").
		Build("https://hasura.io/jwt/claims", customClaims)
	return &claims
}
