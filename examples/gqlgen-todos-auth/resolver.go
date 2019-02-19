package gqlgen_todos_auth

import (
	"context"

	"github.com/Sach97/gqlgenauth/auth/middleware"
	"github.com/Sach97/gqlgenauth/auth/model"
	"github.com/Sach97/gqlgenauth/auth/user"
)

type Resolver struct {
	UserService    *user.Service
	RouterStrategy *middleware.RouterStrategy
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
	//TODO: clean this in gqlgen.yml to point to an Instruction struct
}
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (AuthPayload, error) {
	token, err := r.UserService.Login(&model.UserCredentials{Email: email, Password: password})
	return AuthPayload{
		Token: token,
	}, err
	//TODO: clean this to point to an AuthPayload struct
}

func (r *mutationResolver) VerifyToken(ctx context.Context, token string) (bool, error) {
	return r.UserService.VerifyUserToken(token)
	//TODO
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	userID, err := r.RouterStrategy.GetUserID(ctx)
	//if userid exists retrive user
	return r.UserService.FindByID(userID)
}
