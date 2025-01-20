package crczp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (c *Client) authenticateKeycloak(ctx context.Context) (err error) {
	query := url.Values{}
	query.Add("username", c.Username)
	query.Add("password", c.Password)
	query.Add("client_id", c.ClientID)
	query.Add("grant_type", "password")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/keycloak/realms/CRCZP/protocol/openid-connect/token",
		c.Endpoint), strings.NewReader(query.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err2 := res.Body.Close()
		// If there was an error already, I assume it is more important
		if err == nil {
			err = err2
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode == http.StatusNotFound || res.StatusCode == http.StatusMethodNotAllowed {
		return &Error{ResourceName: "CRCZP Keycloak endpoint", Err: ErrNotFound}
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication to Keycloak failed, status: %d, body: %s", res.StatusCode, body)
	}

	result := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}

	c.Token = result.AccessToken
	c.TokenExpiryTime = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)

	return
}

func (c *Client) refreshToken(ctx context.Context) error {
	if !c.TokenExpiryTime.IsZero() && time.Now().Add(10*time.Second).After(c.TokenExpiryTime) {
		return c.authenticateKeycloak(ctx)
	}
	return nil
}
