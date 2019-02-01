package context

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

//Config holds our config structure
type Config struct {
	AppName string

	//JWT
	JWTSecret   string
	JWTExpireIn time.Duration

	//SMTP
	SMTPIdentity string
	SMTPUsername string
	SMTPPassword string
	SMTPHost     string
	SMTPAddress  string

	//DB
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	//Firebase
	FirebaseAPIKey        string
	AndroidPackageName    string
	DomainURIPrefix       string
	ConfirmationEndpoint  string
	ResetPasswordEndpoint string

	//UserService
	ConfirmationEmailTemplate string

	DebugMode bool
	LogFormat string
}

//LoadConfig load the config from path
func LoadConfig(path string) *Config {
	config := viper.New()
	config.SetConfigName("Config")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error context file: %s \n", err)
	}

	return &Config{
		AppName: config.Get("app-name").(string),

		//JWT

		JWTSecret:   config.Get("auth.jwt-secret").(string),
		JWTExpireIn: config.GetDuration("auth.jwt-expire-in"),

		//SMTP
		SMTPIdentity: config.Get("smtp.identity").(string),
		SMTPUsername: config.Get("smtp.username").(string),
		SMTPPassword: os.Getenv("SMTPPASSWORD"),
		SMTPHost:     config.Get("smtp.host").(string),
		SMTPAddress:  config.Get("smtp.address").(string),

		//DB
		DBHost:     config.Get("db.host").(string),
		DBPort:     config.Get("db.port").(string),
		DBUser:     config.Get("db.user").(string),
		DBPassword: config.Get("db.password").(string),
		DBName:     config.Get("db.dbname").(string),

		//Firebase
		FirebaseAPIKey:        os.Getenv("FIREBASE"),
		AndroidPackageName:    config.Get("deeplinker.androidPackageName").(string),
		DomainURIPrefix:       config.Get("deeplinker.domainURIPrefix").(string),
		ConfirmationEndpoint:  config.Get("deeplinker.confirmationEndpoint").(string),
		ResetPasswordEndpoint: config.Get("deeplinker.resetPasswordEndpoint").(string),

		//Log
		DebugMode: config.Get("log.debug-mode").(bool),
		LogFormat: config.Get("log.log-format").(string),
	}
}

//TODO: strategy for loading from env variable
