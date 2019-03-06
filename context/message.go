package context

import (
	"fmt"
)

const (
	PostMethodSupported = "only post method is allowed"
	TokenError          = "token error"
	UnauthorizedAccess  = "unauthorized access"
)

type MessageService struct {
	AppName string
}

//TODO: load messages from config or template
//TODO: strategy pattern for graphql errors

func NewMessageService(cfg *Config) *MessageService {
	return &MessageService{
		AppName: cfg.AppName,
	}
}

func (m *MessageService) CredentialsError() error {
	return fmt.Errorf("Your email and password combination does not match a %s account. Please try again.", m.AppName)
}
