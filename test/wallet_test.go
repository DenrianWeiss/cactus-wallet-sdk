package test

import (
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	// Get client.
	client := NewClient()
	resp, err := client.CreateWallet(GetBusinessId(), "DEFI", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestGetChainInfo(t *testing.T) {
	// Get client.
	client := NewClient()
	resp, err := client.GetChainInfo(string(constants.ChainNameETH), "")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
