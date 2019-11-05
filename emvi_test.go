package emvi

import (
	"testing"
)

const (
	testClientId     = "jEPnWfxPS6JuEVHFS3q4"
	testClientSecret = "AJaJjQ0EJflce8kbYqKj8bQO3ZLZznxcmXecootMIMCLFf8hzyx44YKKuhiqBTmw"
	testClientOrga   = "leco"
)

var (
	testConfig = &Config{"https://auth.emvi-integration.com", "https://api.emvi-integration.com"}
)

func TestNewClientRefreshToken(t *testing.T) {
	client := getTestClient()

	if err := client.refreshToken(); err != nil {
		t.Fatalf("Token must be refreshed, but was: %v", err)
	}

	if client.TokenType != "Bearer" ||
		client.AccessToken == "" ||
		client.ExpiresIn == 0 {
		t.Fatalf("Client data not as expected: %v", client)
	}
}

func TestClient_GetOrganization(t *testing.T) {
	client := getTestClient()
	result, err := client.GetOrganization()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.Id == "" ||
		result.Name == "" ||
		result.NameNormalized == "" {
		t.Fatalf("Result not as expected: %v", result)
	}
}

func getTestClient() *Client {
	return NewClient(testClientId, testClientSecret, testClientOrga, testConfig)
}
