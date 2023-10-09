package utils

import (
	"testing"
)

const GenerateGetSign = "GET\n" +
	"application/json\n" +
	"\n" +
	"application/json\n" +
	"Tue, 03 Mar 2020 12:26:57 GMT\n" +
	"x-api-key:X5SGmgTAoYaVw1t7oD2p82pHgf0eNNVw3wxYGgM2\n" +
	"x-api-nonce:36dbe33ed529455cb0638eef0f5f59e3\n" +
	"/custody/v1/api/wallets?{b_id=[4a3e2fb40faa4b9d94480559ac01e8de], coin_names=[BTC,LTC], hide_no_coin_wallet=[false], total_market_order=[0]}"

const GeneratePostSign = "POST\n" +
	"application/json\n" +
	"3cLLd5MmUAMM2BneR7eT0NV9AZ4TUJ2F7xy31krmInQ=\n" +
	"application/json\n" +
	"Tue, 03 Mar 2020 13:26:57 GMT\n" +
	"x-api-key:X5SGmgTAoYaVw1t7oD2p82pHgf0eNNVw3wxYGgM2\n" +
	"x-api-nonce:36dbe33ed529455cb0638eef0f5f59e3\n" +
	"/custody/v1/api/projects/4a3e2fb40faa4b9d94480559ac01e8de/order/create"

const GenerateGetParams = "{b_id=[4a3e2fb40faa4b9d94480559ac01e8de], coin_names=[BTC,LTC], hide_no_coin_wallet=[false], total_market_order=[0]}"

func TestGenerateSignStringGet(t *testing.T) {
	resp := GenerateSignString("GET",
		nil,
		"X5SGmgTAoYaVw1t7oD2p82pHgf0eNNVw3wxYGgM2",
		"36dbe33ed529455cb0638eef0f5f59e3",
		"/custody/v1/api/wallets",
		"{b_id=[4a3e2fb40faa4b9d94480559ac01e8de], coin_names=[BTC,LTC], hide_no_coin_wallet=[false], total_market_order=[0]}",
		"Tue, 03 Mar 2020 12:26:57 GMT")
	if resp != GenerateGetSign {
		t.Errorf("GenerateSignString() = %v, want %v", resp, GenerateGetSign)
	}
}

// Todo: Generate POST sign string - cactus does not provide example for this.

func TestEncodeGetParams(t *testing.T) {
	params := map[string]string{
		"b_id":                "4a3e2fb40faa4b9d94480559ac01e8de",
		"coin_names":          "BTC,LTC",
		"hide_no_coin_wallet": "false",
		"total_market_order":  "0",
	}
	resp := EncodeGetParams(params)
	if resp != GenerateGetParams {
		t.Errorf("EncodeGetParams() = %v, want %v", resp, GenerateGetParams)
	}
}
