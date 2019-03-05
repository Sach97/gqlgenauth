package ninshoo_hasura

import (
	"context"

	"github.com/Sach97/ninshoo/model"
	"github.com/Sach97/ninshoo/user"
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

func (r *mutationResolver) Signup(ctx context.Context, email string, password string) (*Instructions, error) {
	instructions := r.UserService.Signup(&model.UserCredentials{Email: email, Password: password})
	return &Instructions{
		Text: instructions,
	}, nil
	//TODO: clean this in gqlgen.yml to point to an Instruction struct
}
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*AuthPayload, error) {
	token, err := r.UserService.Login(&model.UserCredentials{Email: email, Password: password})
	return &AuthPayload{
		Token: token,
	}, err
	//TODO: clean this to point to an AuthPayload struct
}

func (r *mutationResolver) VerifyToken(ctx context.Context, token string) (bool, error) {
	return r.UserService.VerifyUserToken(token)

}

//TODO: remove me resolver
type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "world", nil
}
