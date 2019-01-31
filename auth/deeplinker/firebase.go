package deeplinker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sach97/gqlgenauth/auth/context"
)

const (
	// DefaultEndpoint contains endpoint URL of firebase short link service.
	DefaultEndpoint = "https://firebasedynamiclinks.googleapis.com/v1/shortLinks"
)

// Payload sets the payload.
type Payload struct {
	DynamicLinkInfo *DynamicLinkInfo `json:"dynamicLinkInfo"`
}

// DynamicLinkInfo sets the domain uri prefix, the link and the android info on the payload.
type DynamicLinkInfo struct {
	DomainURIPrefix string       `json:"domainUriPrefix"`
	Link            *string      `json:"link"`
	AndroidInfo     *AndroidInfo `json:"androidInfo"`
}

// AndroidInfo sets the android package name on the payload.
type AndroidInfo struct {
	AndroidPackageName string `json:"androidPackageName"`
}

// FireBaseClient is interface structure
type FireBaseClient struct {
	Endpoint              string
	APIKey                string
	AndroidPackageName    string
	DomainURIPrefix       string
	ConfirmationEndpoint  string
	ResetPasswordEndpoint string
}

// ErrorResponse is the shape of a response in case of a code 400
type ErrorResponse struct {
	Code    string
	Message string
	Status  string
}

// SuccessResponse is the shape of a response in case of a code 200
type SuccessResponse struct {
	Shortlink   string
	Warning     []Warning
	PreviewLink string
}

// Warning is the struct of warning messages
type Warning struct {
	WarningCode    string
	WarningMessage string
}

// NewFireBaseClient instantiate a new Firebase Client
func NewFireBaseClient(config *context.Config) *FireBaseClient {
	if config.FirebaseApiKey == "" {
		panic("You must set your API key")
	}

	return &FireBaseClient{
		Endpoint:              DefaultEndpoint,
		APIKey:                config.FirebaseAPIKey,
		AndroidPackageName:    config.AndroidPackageName,
		DomainURIPrefix:       config.DomainURIPrefix,
		ConfirmationEndpoint:  config.ConfirmationEndpoint,
		ResetPasswordEndpoint: config.ResetPasswordEndpoint,
	}
}

type Response struct {
	SuccessResponse *SuccessResponse
	ErrorResponse   *ErrorResponse
}

// GetDynamicLink creates a new firebase link
func (c *FireBaseClient) GetDynamicLink(token string, confirm bool) (string, error) {

	var link string
	if confirm {
		link = fmt.Sprintf("%s?token=%s", c.ConfirmationEndpoint, token)
	} else {
		link = fmt.Sprintf("%s?token=%s", c.ResetPasswordEndpoint, token)
	}
	androidInfo := AndroidInfo{AndroidPackageName: c.AndroidPackageName}
	dynamicLinkInfo := DynamicLinkInfo{DomainURIPrefix: c.DomainURIPrefix, Link: &link, AndroidInfo: &androidInfo}

	p := Payload{DynamicLinkInfo: &dynamicLinkInfo}
	fmt.Println(c.APIKey)
	payload, _ := json.Marshal(p)

	//format url
	url := fmt.Sprintf("%s?key=%s", c.Endpoint, c.APIKey)

	// create request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	// add header
	req.Header.Add("content-type", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	// if resp.StatusCode == 200 {
	// 	json.NewDecoder(resp.Body).Decode(successRes)
	// 	return &Response{
	// 		SuccessResponse: successRes,
	// 	}
	// }
	// errRes := new(ErrorResponse)
	// json.NewDecoder(resp.Body).Decode(successRes)
	// return &Response{
	// 	ErrorResponse: errRes,
	// }

	if resp.StatusCode == 200 {
		successRes := new(SuccessResponse)
		json.NewDecoder(resp.Body).Decode(successRes)
		return successRes.Shortlink, nil
	} else {
		errRes := new(ErrorResponse)
		json.NewDecoder(resp.Body).Decode(errRes)
		return "", fmt.Errorf("Code :%s, Status: %s, Message: %s", errRes.Code, errRes.Status, errRes.Message)
	}

}
