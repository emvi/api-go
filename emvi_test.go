package emvi

import (
	"testing"
)

const (
	// this is our API test organization
	testClientId     = "n0QizLaKCH1RoBB3sYWK"
	testClientSecret = "P5S7cjSFw2QDTjLffzOHMrCnFkOhmAJ6fQPzW2JLP2C18RWKlOysOdYNK1mTr3WD"
	testClientOrga   = "api-test"
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

func TestClient_FindArticles(t *testing.T) {
	client := getTestClient()
	articles, count, err := client.FindArticles("test", nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(articles) != 2 || count != 2 {
		t.Fatalf("Result not as expected: %v %v", len(articles), count)
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

func TestClient_GetArticle(t *testing.T) {
	client := getTestClient()
	article, content, authors, err := client.GetArticle("z0odQAGdLO", "lMK14kgOv8", 0)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if article.Id != "z0odQAGdLO" || content == nil || len(authors) != 1 {
		t.Fatalf("Result not as expected: %v %v %v", article, content, authors)
	}
}

func TestClient_GetLanguages(t *testing.T) {
	client := getTestClient()
	result, err := client.GetLanguages()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Result not as expected: %v", result)
	}
}

func getTestClient() *Client {
	return NewClient(testClientId, testClientSecret, testClientOrga, testConfig)
}
