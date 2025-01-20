package crczp

import (
	"context"
	"net/http"
	"time"
)

// Client struct stores information for authentication to the CRCZP API.
// All functions are methods of this struct
type Client struct {
	// Endpoint of the CRCZP instance to connect to. For example `https://your.crczp.ex`.
	Endpoint string

	// ClientID used by the CRCZP instance OIDC provider.
	ClientID string

	// HTTPClient which is used to do requests.
	HTTPClient *http.Client

	// Bearer Token which is used for authentication to the CRCZP instance. Is set by NewClient function.
	Token string

	// Time when Token expires, used to refresh it automatically when required. Is set by NewClient function.
	// Is used only with CRCZP instances using Keycloak OIDC provider.
	TokenExpiryTime time.Time

	// Username of the user to login as.
	Username string

	// Password of the user to login as.
	Password string

	// How many times should a failed HTTP request be retried. There is a delay of 100ms before the first retry.
	// The delay is doubled before each following retry.
	RetryCount int
}

// NewClientWithToken creates and returns a Client which uses an already created Bearer token.
func NewClientWithToken(endpoint, clientId, token string) (*Client, error) {
	client := Client{
		Endpoint:   endpoint,
		ClientID:   clientId,
		HTTPClient: http.DefaultClient,
		Token:      token,
	}

	return &client, nil
}

// NewClient creates and returns a Client which uses username and password for authentication.
// The username and password is used to login to Keycloak of the CRCZP instance.
func NewClient(endpoint, clientId, username, password string) (*Client, error) {
	client := Client{
		Endpoint:   endpoint,
		ClientID:   clientId,
		HTTPClient: http.DefaultClient,
		Username:   username,
		Password:   password,
	}
	err := client.authenticateKeycloak(context.Background())
	if err != nil {
		return nil, err
	}
	return &client, nil
}
