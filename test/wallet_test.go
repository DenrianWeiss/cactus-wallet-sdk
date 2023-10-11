package test

import "testing"

func TestCreateWallet(t *testing.T) {
	// Get client.
	client := NewClient()
	resp, err := client.CreateWallet(GetBusinessId(), "DEFI", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
