package kypo

import (
	"net/http"
	"time"
)

type Client struct {
	Endpoint        string
	ClientID        string
	HTTPClient      *http.Client
	Token           string
	TokenExpiryTime time.Time
	Username        string
	Password        string
}

func NewClientWithToken(endpoint, clientId, token string) (*Client, error) {
	client := Client{
		Endpoint:   endpoint,
		ClientID:   clientId,
		HTTPClient: http.DefaultClient,
		Token:      token,
	}

	return &client, nil
}

func NewClient(endpoint, clientId, username, password string) (*Client, error) {
	client := Client{
		Endpoint:   endpoint,
		ClientID:   clientId,
		HTTPClient: http.DefaultClient,
		Username:   username,
		Password:   password,
	}
	err := client.authenticate()
	if err != nil {
		return nil, err
	}
	return &client, nil
}
