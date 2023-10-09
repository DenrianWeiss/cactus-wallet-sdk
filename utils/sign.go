package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"slices"
	"time"
)

const Accept = "application/json"
const RequestContentType = "application/json"
const ServiceName = "api"

func GetCurrentGmtTime() string {
	currentTime := time.Now().UTC().Format(time.RFC1123)
	currentTime = currentTime[0:len(currentTime)-3] + "GMT"
	return currentTime
}

// GenerateSignString generates the string to be signed
// RequestMethod: GET, POST, PUT, DELETE, use the constants in http/method.go
// body: the body of the request, if it's post request, pass nil
// XApiKey: the api key
// XApiNonce: the nonce
// uri: the uri of the request
// FormattedParams: the formatted params, if there's no params, pass ""
// currentTime: the current time in RFC1123 format, using GMT TZ
func GenerateSignString(RequestMethod string, body []byte, XApiKey string, XApiNonce string, uri string, FormattedParams string, currentTime string) string {
	// Hash body
	bodyHash := sha256.Sum256(body)
	bodyHashBase := base64.StdEncoding.EncodeToString(bodyHash[:])
	if body == nil {
		bodyHashBase = ""
	}
	withoutParam := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\nx-api-key:%s\nx-api-nonce:%s\n%s", RequestMethod, Accept, bodyHashBase, RequestContentType, currentTime, XApiKey, XApiNonce, uri)
	if FormattedParams == "" {
		return withoutParam
	}
	return fmt.Sprintf("%s?%s", withoutParam, FormattedParams)
}

// Sign signs the string
// signReq: the string to be signed
func Sign(signReq []byte, key *ecdsa.PrivateKey) ([]byte, error) {
	cryptoRand := rand.Reader
	hash := sha256.Sum256(signReq)
	sign, err := ecdsa.SignASN1(cryptoRand, key, hash[:])
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// GenerateAuthorizationHeader generates the authorization header
// signReq: the string to be signed
// ApiKeyId: the api key id
// key: the private key
func GenerateAuthorizationHeader(signReq []byte, ApiKeyId string, key *ecdsa.PrivateKey) (string, error) {
	signature, err := Sign(signReq, key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s:%s", ServiceName, ApiKeyId, base64.StdEncoding.EncodeToString(signature)), nil
}

// EncodeGetParams encodes the post params, to the cactus format, like {param_name=[param_value], param_name2=[param_value2]} example:{b_id=[4a3e2fb40faa4b9d94480559ac01e8de], coin_names=[BTC,LTC], hide_no_coin_wallet=[false], total_market_order=[0]}
func EncodeGetParams(req map[string]string) string {
	result := ""
	keys := make([]string, 0)
	for k := range req {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, key := range keys {
		result += key + "=[" + req[key] + "], "
	}
	// remove the last ", "
	return "{" + result[0:len(result)-2] + "}"
}

func EncodeGetQuery(req map[string]string) string {
	result := ""
	keys := make([]string, 0)
	for k := range req {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, key := range keys {
		result += key + "=" + req[key] + "&"
	}
	// remove the last "&"
	return result[0 : len(result)-1]
}

func GenerateUuid() string {
	// Generate a random uuid
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
