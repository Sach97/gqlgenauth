package tmp

import (
	"context"

	"github.com/Sach97/gqlgenauth/examples/gqlgen-todos-auth"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, email string, password string) (gqlgen_todos_auth.Instructions, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (gqlgen_todos_auth.AuthPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) VerifyToken(ctx context.Context, token string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*gqlgen_todos_auth.User, error) {
	panic("not implemented")
}
