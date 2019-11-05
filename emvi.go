package emvi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	defaultAuthHost = "https://auth.emvi.com"
	defaultApiHost  = "https://api.emvi.com"

	grantType              = "client_credentials"
	authenticationEndpoint = "/api/v1/auth/token"
	searchArticlesEndpoint = "/api/v1/search/article"
	searchListsEndpoint    = "/api/v1/search/list"
	searchTagsEndpoint     = "/api/v1/search/tag"
	searchAllEndpoint      = "/api/v1/search"
	organizationEndpoint   = "/api/v1/organization"
	articleEndpoint        = "/api/v1/article/{id}"
	articleHistoryEndpoint = "/api/v1/article/{id}/history"
	languagesEndpoint      = "/api/v1/lang"
	languageEndpoint       = "/api/v1/lang/{id}"
	pinnedEndpoint         = "/api/v1/pin"
	listEndpoint           = "/api/v1/articlelist/{id}"
	listEntriesEndpoint    = "/api/v1/articlelist/{id}/entry"
	tagEndpoint            = "/api/v1/tag/{name}"
)

// Client connects to the Emvi client API.
type Client struct {
	ClientId     string
	ClientSecret string
	Organization string
	AuthHost     string
	ApiHost      string
	TokenType    string
	AccessToken  string
	ExpiresIn    int // TTL in seconds
}

// Config is used for advanced client configuration.
type Config struct {
	AuthHost string
	ApiHost  string
}

// NewClient returns a new Client instance.
// For clientId and clientSecret, use the keys generated in the administration by Emvi.
// For organization pass the organization subdomain (e.g. https://my-orga.emvi.com/ -> my-orga).
// The config object is optional.
func NewClient(clientId, clientSecret, organization string, config *Config) *Client {
	if config == nil {
		config = &Config{defaultAuthHost, defaultApiHost}
	}

	if config.AuthHost == "" {
		config.AuthHost = defaultAuthHost
	}

	if config.ApiHost == "" {
		config.ApiHost = defaultApiHost
	}

	return &Client{ClientId: clientId,
		ClientSecret: clientSecret,
		Organization: organization,
		AuthHost:     config.AuthHost,
		ApiHost:      config.ApiHost}
}

func (client *Client) refreshToken() error {
	body := struct {
		GrantType    string `json:"grant_type"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{grantType, client.ClientId, client.ClientSecret}
	bodyJson, err := json.Marshal(&body)

	if err != nil {
		return err
	}

	c := http.Client{}
	resp, err := c.Post(client.AuthHost+authenticationEndpoint, "application/json", bytes.NewBuffer(bodyJson))

	if err != nil {
		return err
	}

	respJson := struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&respJson); err != nil {
		return err
	}

	client.TokenType = respJson.TokenType
	client.AccessToken = respJson.AccessToken
	client.ExpiresIn = respJson.ExpiresIn

	return nil
}

func (client *Client) GetOrganization() (*Organization, error) {
	result := new(Organization)

	if err := client.performGet(client.ApiHost+organizationEndpoint, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) performGet(url string, respObj interface{}) error {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	client.setRequestHeader(req)
	c := http.Client{}
	resp, err := c.Do(req)

	if err != nil {
		return err
	}

	// retry on 401
	if resp.StatusCode == http.StatusUnauthorized {
		if err := client.refreshToken(); err != nil {
			return err
		}

		client.setRequestHeader(req)
		resp, err = c.Do(req)

		if err != nil {
			return err
		}
	}

	// check status is ok
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("received status code %d on request: %s", resp.StatusCode, string(body)))
	}

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&respObj); err != nil {
		return err
	}

	return nil
}

func (client *Client) setRequestHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+client.AccessToken)
	req.Header.Set("Organization", client.Organization)
	req.Header.Set("Client", client.ClientId)
}
