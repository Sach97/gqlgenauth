package ninshoo

import (
	"github.com/Sach97/ninshoo/builder"
	"github.com/Sach97/ninshoo/context"
	"github.com/Sach97/ninshoo/db"
	"github.com/Sach97/ninshoo/deeplinker"
	"github.com/Sach97/ninshoo/jwt"
	"github.com/Sach97/ninshoo/mailer"
	"github.com/Sach97/ninshoo/tokenizer"
	"github.com/Sach97/ninshoo/user"
	"github.com/Sach97/ninshoo/utils"
)

func NewNinshoo() *user.Service {
	// // Context Stuffs
	cfg := context.LoadConfig(".")

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

	err = m.Ping()
	if err != nil {
		l.Errorf("Can't ping smtp port : %s", err)
	}

	//JWT stuffs
	auth := jwt.NewAuthService(cfg)

	//Message service stuffs
	msg := context.NewMessageService(cfg)

	//Builder service stuffs
	b := builder.NewBuilderService(cfg)

	// User service stuffs
	return user.NewUserService(msg, s, l, auth, &t, m, d, b)
}
