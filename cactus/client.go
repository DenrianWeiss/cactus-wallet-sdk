package cactus

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Cactus struct {
	BaseUri    string
	ApiKey     string
	HttpClient *http.Client
	LogLevel   int
	cachedJwt  string
}

func NewCactus(baseUri string, apiKey string, client *http.Client, logLevel int) *Cactus {
	if client == nil {
		client = http.DefaultClient
	}
	return &Cactus{
		BaseUri:    baseUri,
		ApiKey:     apiKey,
		HttpClient: client,
		LogLevel:   logLevel,
	}
}

func (c *Cactus) Log(l string, level int) {
	if level > c.LogLevel {
		log.Println(l)
	}
}

func (c *Cactus) GetJwt() (string, CatcusError) {
	req := &CreateTokenReq{
		GrantType:    GrantTypeRefreshToken,
		RefreshToken: c.ApiKey,
	}
	r, _ := json.Marshal(req)
	post, err := c.HttpClient.Post(c.BaseUri+"/token", "application/json", bytes.NewReader(r))
	if err != nil {
		return "", err
	}
	defer post.Body.Close()
	resp := &CreateTokenResp{}
	err = json.NewDecoder(post.Body).Decode(resp)
	if err != nil || resp == nil || resp.AccessToken == "" {
		return "", CatcusError(errors.New(ErrorCreatingJwt))
	}
	c.cachedJwt = resp.AccessToken
	return resp.AccessToken, nil
}

func (c *Cactus) getCachedJwt() (string, CatcusError) {
	if c.cachedJwt == "" {
		return c.GetJwt()
	}
	return c.cachedJwt, nil
}

func (c *Cactus) get(endpoint string, params map[string]string) ([]byte, CatcusError) {
	// Assemble query string
	queryString := ""
	for k, v := range params {
		queryString += k + "=" + v + "&"
	}
	// Remove last &
	queryString = queryString[:len(queryString)-1]
	// Assemble request
	requestPath := c.BaseUri + endpoint + "?" + queryString
	// Add http header
	request, err := http.NewRequest(http.MethodGet, requestPath, nil)
	if err != nil {
		return nil, err
	}
	// Add Content-Type
	request.Header.Add("Content-Type", "application/json")
	// Add JWT
	jwt, err := c.GetJwt()
	request.Header.Add("Authorization", "Bearer "+jwt)
	// Send request
	response, err := c.HttpClient.Do(request)
	// Handle JWT Expired
	if response != nil && response.StatusCode == http.StatusForbidden {
		// Refresh JWT
		jwt, err = c.GetJwt()
		// Let's try again
		// Delete old JWT
		request.Header.Del("Authorization")
		request.Header.Add("Authorization", "Bearer "+jwt)
		response, err = c.HttpClient.Do(request)
		if err != nil {
			return nil, CatcusError(err)
		}
	}
	if err != nil {
		return nil, CatcusError(err)
	}
	defer response.Body.Close()
	// Read response
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, CatcusError(err)
	}
	return all, nil
}
func (c *Cactus) post(endpoint string, uriParams map[string]string, body interface{}) ([]byte, CatcusError) {
	// Assemble query string
	queryString := ""
	for k, v := range uriParams {
		queryString += k + "=" + v + "&"
	}
	// Remove last &
	queryString = queryString[:len(queryString)-1]
	// Assemble request
	requestPath := c.BaseUri + endpoint + "?" + queryString
	// Add http header
	request, err := http.NewRequest(http.MethodPost, requestPath, nil)
	if err != nil {
		return nil, err
	}
	// Add Content-Type
	request.Header.Add("Content-Type", "application/json")
	// Add JWT
	jwt, err := c.GetJwt()
	request.Header.Add("Authorization", "Bearer "+jwt)
	// Add body
	r, err := json.Marshal(body)
	if err != nil {
		return nil, CatcusError(err)
	}
	request.Body = io.NopCloser(bytes.NewReader(r))
	// Send request
	response, err := c.HttpClient.Do(request)
	// Handle JWT Expired
	if response != nil && response.StatusCode == http.StatusForbidden {
		// Refresh JWT
		jwt, err = c.GetJwt()
		// Let's try again
		// Delete old JWT
		request.Header.Del("Authorization")
		request.Header.Add("Authorization", "Bearer "+jwt)
		response, err = c.HttpClient.Do(request)
		if err != nil {
			return nil, CatcusError(err)
		}
	}
	if err != nil {
		return nil, CatcusError(err)
	}
	defer response.Body.Close()
	// Read response
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, CatcusError(err)
	}
	return all, nil
}
