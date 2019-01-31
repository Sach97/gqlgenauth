package context

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//Config holds our config structure
type Config struct {
	AppName string

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
	ConfirmationEmailTemplate  string
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
	}
}

//TODO: strategy for loading from env variable
