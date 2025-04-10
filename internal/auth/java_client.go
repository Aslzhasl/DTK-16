package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JavaAuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewJavaAuthClient(baseURL string) *JavaAuthClient {
	return &JavaAuthClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *JavaAuthClient) VerifyUser(token string) (UserInfo, error) {
	reqBody := map[string]string{"token": token}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return UserInfo{}, fmt.Errorf("marshal error")
	}

	req, _ := http.NewRequest("POST", c.baseURL+"/auth/validate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return UserInfo{}, fmt.Errorf("auth service error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserInfo{}, fmt.Errorf("auth failed")
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return UserInfo{}, fmt.Errorf("invalid response format")
	}

	return userInfo, nil
}
