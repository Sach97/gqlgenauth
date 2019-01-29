package context

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//Config holds our config structure
type Config struct {
	AppName string

	//DB config
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	FirebaseApiKey        string
	AndroidPackageName    string
	DomainURIPrefix       string
	ConfirmationEndpoint  string
	ResetPasswordEndpoint string
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

		DBHost:     config.Get("db.host").(string),
		DBPort:     config.Get("db.port").(string),
		DBUser:     config.Get("db.user").(string),
		DBPassword: config.Get("db.password").(string),
		DBName:     config.Get("db.dbname").(string),

		FirebaseApiKey:        os.Getenv("FIREBASE"),
		AndroidPackageName:    config.Get("deeplinker.androidPackageName").(string),
		DomainURIPrefix:       config.Get("deeplinker.domainURIPrefix").(string),
		ConfirmationEndpoint:  config.Get("deeplinker.confirmationEndpoint").(string),
		ResetPasswordEndpoint: config.Get("deeplinker.resetPasswordEndpoint").(string),
	}
}

//TODO: strategy for loading from env variable
