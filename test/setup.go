package test

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/DenrianWeiss/cactus-wallet-sdk/cactus"
	"io"
	"net/http"
	"os"
)

const pkcs8PrivateHeader = "-----BEGIN PRIVATE KEY-----"
const pkcs8PrivateFooter = "-----END PRIVATE KEY-----"

func GetPkcs8PrivateKey() *ecdsa.PrivateKey {
	fileName := os.Getenv("PRIVATE_KEY_FILE")
	if fileName == "" {
		fileName = "private_key.pem"
	}
	// Read Private Key
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// Parse Private Key
	// Remove anything before and including "-----BEGIN PRIVATE KEY-----"
	index := bytes.Index(content, []byte(pkcs8PrivateHeader))
	if index == -1 {
		panic("Invalid private key file")
	}
	content = content[index+len(pkcs8PrivateHeader):]
	// Remove Footer
	index = bytes.Index(content, []byte(pkcs8PrivateFooter))
	if index == -1 {
		panic("Invalid private key file")
	}
	content = content[:index]
	content = bytes.TrimPrefix(content, []byte("\n"))
	content = bytes.TrimSuffix(content, []byte("\n"))
	// Decode base64
	decodedLength := base64.StdEncoding.DecodedLen(len(content))
	decoded := make([]byte, decodedLength)
	_, err = base64.StdEncoding.Decode(decoded, content)
	if err != nil {
		panic(err)
	}
	// X509 Load private key
	privateKey, err := x509.ParsePKCS8PrivateKey(decoded)
	if err != nil {
		panic(err)
	}
	// Try casting to ecdsa.PrivateKey
	ecdsaPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		panic("Invalid private key file")
	}
	return ecdsaPrivateKey
}

func NewClient() *cactus.Cactus {
	// First Load Up Private Key
	privateKey := GetPkcs8PrivateKey()
	apiKey := os.Getenv("API_KEY")
	apiKeyId := os.Getenv("API_KEY_ID")
	// Create Client
	client := cactus.NewCactus("https://api.mycactus.dev", apiKey, apiKeyId, privateKey, http.DefaultClient, 0)
	return client
}

func GetBusinessId() string {
	return os.Getenv("BUSINESS_ID")
}
