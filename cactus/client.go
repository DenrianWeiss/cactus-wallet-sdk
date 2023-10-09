package cactus

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/DenrianWeiss/cactus-wallet-sdk/utils"
	"io"
	"net/http"
)

type Cactus struct {
	BaseUri    string
	XApiKey    string
	ApiKeyID   string
	PrivateKey *ecdsa.PrivateKey
	HttpClient *http.Client
	LogLevel   int
}

func NewCactus(baseUri string, xApiKey string, apiKeyId string, privateKey *ecdsa.PrivateKey, client *http.Client, logLevel int) *Cactus {
	if client == nil {
		client = http.DefaultClient
	}
	return &Cactus{
		BaseUri:    baseUri,
		XApiKey:    xApiKey,
		ApiKeyID:   apiKeyId,
		PrivateKey: privateKey,
		HttpClient: client,
		LogLevel:   logLevel,
	}
}

func (c *Cactus) post(path string, body map[string]interface{}) ([]byte, error) {
	// Encode body
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	// Hash Request Body
	bodyHash := sha256.Sum256(bodyBytes)
	bodyHashBase := base64.StdEncoding.EncodeToString(bodyHash[:])

	// Generate sign string
	/// Get TimeStamp
	currentTime := utils.GetCurrentGmtTime()
	/// Sign
	nonce := utils.GenerateUuid()
	signString := utils.GenerateSignString(http.MethodPost, bodyBytes, c.XApiKey, nonce, path, "", currentTime)
	// Sign
	header, err := utils.GenerateAuthorizationHeader([]byte(signString), c.ApiKeyID, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// Assemble Req
	req, err := http.NewRequest(http.MethodPost, c.BaseUri+path, bytes.NewReader(bodyBytes))
	// Add headers
	req.Header.Add("Accept", utils.Accept)
	req.Header.Add("Content-Type", utils.RequestContentType)
	req.Header.Add("x-api-key", c.XApiKey)
	req.Header.Add("x-api-nonce", nonce)
	req.Header.Add("Content-SHA256", bodyHashBase)
	req.Header.Add("Date", currentTime)
	req.Header.Add("Authorization", header)
	// Send request
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read response
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (c *Cactus) get(path string, params map[string]string) ([]byte, error) {
	// Encode url params
	paramQ := utils.EncodeGetQuery(params)
	// Generate sign string
	/// Get TimeStamp
	currentTime := utils.GetCurrentGmtTime()
	/// Sign
	nonce := utils.GenerateUuid()
	signString := utils.GenerateSignString(http.MethodGet, nil, c.XApiKey, nonce, path, paramQ, currentTime)
	// Add headers
	header, err := utils.GenerateAuthorizationHeader([]byte(signString), c.ApiKeyID, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// Assemble Req
	req, err := http.NewRequest(http.MethodGet, c.BaseUri+path+"?"+paramQ, nil)
	req.Header.Add("Accept", utils.Accept)
	req.Header.Add("Content-Type", utils.RequestContentType)
	req.Header.Add("x-api-key", c.XApiKey)
	req.Header.Add("x-api-nonce", nonce)
	req.Header.Add("Date", currentTime)
	req.Header.Add("Authorization", header)
	// Send request
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read response
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
