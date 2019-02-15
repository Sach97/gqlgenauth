package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	gcontext "github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/db"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/jwt"
	"github.com/Sach97/gqlgenauth/auth/mailer"
	"github.com/Sach97/gqlgenauth/auth/tokenizer"
	"github.com/Sach97/gqlgenauth/auth/user"
	"github.com/Sach97/gqlgenauth/auth/utils"
	gqlgen_todos_auth "github.com/Sach97/gqlgenauth/examples/gqlgen-todos-auth"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// // Context Stuffs
	cfg := gcontext.LoadConfig(".")

	// // Token stuffs
	RedisClient := tokenizer.NewRedisClient(cfg)
	t := tokenizer.Tokenizer{RedisClient}

	// // Mail stuffs
	m := mailer.NewMailer(cfg)

	// // Firebase STUFFS
	d := deeplinker.NewFireBaseClient(cfg)

	//DB STUFFS
	sql := db.Strategy(db.DriverSQL{Name: "postgres"})

	s, err := sql.OpenDB(cfg)
	if err != nil {
		panic(err)
	}

	//Log stuffs
	l := utils.NewLoggerService(cfg)

	//JWT stuffs
	a := jwt.NewAuthService(cfg)

	//Message service stuffs
	msg := gcontext.NewMessageService(cfg)

	// User service stuffs
	u := user.NewUserService(msg, s, l, a, &t, m, d)
	// credentials := model.UserCredentials{Email: "sacha.arbonel@hotmail.fr", Password: "secretpassword"}
	// signup := u.Signup(&credentials)
	// fmt.Println(signup)
	//ctx = context.WithValue(ctx, "config", cfg)
	//setServices(ctx,Services{
	//	cfg : cfg,
	//  log : log,
	//  db : db,
	//})
	//ctx = context.WithValue(ctx, "log", l) //
	//ctx = context.WithValue(ctx, "userService", u)

	// user, _ := ctx.Value("userService").(*user.Service)
	// instructions := user.Signup(&model.UserCredentials{Email: "sacha.arbonel@hotmail.fr", Password: "secretpassword"})
	// fmt.Println(instructions)
	//ctx = user.SetUserService(ctx, u)
	//ctx = context.WithValue(ctx, "authService", a)
	//ctx = context.WithValue(ctx, "dbService", s)

	//TODO: find a better way to do this like auth
	// r := chi.NewRouter()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen_todos_auth.NewExecutableSchema(gqlgen_todos_auth.Config{Resolvers: &gqlgen_todos_auth.Resolver{
		UserService: u,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
