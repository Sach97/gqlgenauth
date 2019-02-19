package gqlgen_todos_auth

import (
	"context"

	"github.com/Sach97/gqlgenauth/auth/model"
	"github.com/Sach97/gqlgenauth/auth/user"
)

type Resolver struct {
	UserService *user.Service
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, email string, password string) (Instructions, error) {
	instructions := r.UserService.Signup(&model.UserCredentials{Email: email, Password: password})
	return Instructions{
		Text: instructions,
	}, nil
}
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (AuthPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) VerifyToken(ctx context.Context, token string) (bool, error) {
	return r.UserService.VerifyUserToken(token)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	panic("not implemented")
}
