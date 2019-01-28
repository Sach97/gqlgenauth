package deeplinker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// DefaultEndpoint contains endpoint URL of FCM service.
	DefaultEndpoint = "https://firebasedynamiclinks.googleapis.com/v1/shortLinks"
)

// Payload sets the payload.
type Payload struct {
	DynamicLinkInfo *DynamicLinkInfo `json:"dynamicLinkInfo"`
}

// DynamicLinkInfo sets the domain uri prefix, the link and the android info on the payload.
type DynamicLinkInfo struct {
	DomainURIPrefix string       `json:"domainUriPrefix"`
	Link            string       `json:"link"`
	AndroidInfo     *AndroidInfo `json:"androidInfo"`
}

// AndroidInfo sets the android package name on the payload.
type AndroidInfo struct {
	AndroidPackageName string `json:"androidPackageName"`
}

// FireBaseClient is interface structure
type FireBaseClient struct {
	Endpoint string
	APIKey   string
}

// Error400Response is the shape of a response in case of a code 400
type Error400Response struct {
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
func NewFireBaseClient(apiKey string) *FireBaseClient {
	return &FireBaseClient{
		Endpoint: DefaultEndpoint,
		APIKey:   apiKey,
	}
}

// GetDynamicLink creates a new firebase link
func (c *FireBaseClient) GetDynamicLink(p *Payload) {

	payload, _ := json.Marshal(p)

	//format url
	url := fmt.Sprintf("%s?key=%s", c.Endpoint, c.APIKey)

	// create request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	// add header
	req.Header.Add("content-type", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	response := new(SuccessResponse)

	if resp.StatusCode == 200 {
		json.NewDecoder(resp.Body).Decode(response)
		fmt.Println(response.PreviewLink)
	}

}
