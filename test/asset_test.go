package test

import "testing"

func TestGetTotalAssetNotionalValue(t *testing.T) {
	// Get client.
	client := NewClient()
	resp, err := client.GetCurrentAssetNotionalValue("")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
