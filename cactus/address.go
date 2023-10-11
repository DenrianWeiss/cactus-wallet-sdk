package cactus

import (
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
)

const (
	ApplyNewAddressUrl        = "/custody/v1/api/projects/%s/wallets/%s/addresses/apply"
	GetSingleAddressUrl       = "/custody/v1/api/projects/%s/wallets/%s/addresses/%s"
	GetAddressListUrl         = "/custody/v1/api/projects/%s/wallets/%s/addresses"
	EditAddressDescriptionUrl = "/custody/v1/api/projects/%s/wallets/%s/addresses/%s"
	VerifyAddressFormat       = "/custody/v1/api/addresses/type/check"
)

// ApplyNewAddressResp is the response of ApplyNewAddress
// Data is an array of addresses
type ApplyNewAddressResp struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Data       []string `json:"data"`
	Successful bool     `json:"successful"`
}

// ApplyNewAddress applies new address for a wallet
// bId: business id
// walletCode: wallet code
// coinName: coin name, optional
// addressNum: number of addresses to apply
// addressType: address type, only accept "NORMAL_ADDRESS"
func (c *Cactus) ApplyNewAddress(bId string, walletCode string, coinName constants.CactusToken, addressNum int, addressType string) (*ApplyNewAddressResp, error) {
	req := map[string]interface{}{
		"address_num":  addressNum,
		"address_type": addressType,
		"b_id":         bId,
		"wallet_code":  walletCode,
	}
	if coinName != "" {
		req["coin_name"] = string(coinName)
	}
	resp, err := c.post(ApplyNewAddressUrl, req)
	if err != nil {
		return nil, err
	}
	var applyNewAddressResp ApplyNewAddressResp
	err = json.Unmarshal(resp, &applyNewAddressResp)
	if err != nil {
		return nil, err
	}
	return &applyNewAddressResp, nil
}

type GetSingleAddressResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		Address          string      `json:"address"`
		AddressStorage   string      `json:"address_storage"`
		AddressType      string      `json:"address_type"`
		AvailableAmount  interface{} `json:"available_amount"`
		CoinName         string      `json:"coin_name"`
		BchAddressFormat interface{} `json:"bch_address_format"`
		BId              string      `json:"b_id"`
		Description      string      `json:"description"`
		DomainId         string      `json:"domain_id"`
		FreezeAmount     interface{} `json:"freeze_amount"`
		TotalAmount      int         `json:"total_amount"`
		WalletCode       string      `json:"wallet_code"`
		WalletType       string      `json:"wallet_type"`
	} `json:"data"`
}

// GetSingleAddress gets the single address info
// bId: business id
// walletCode: wallet code
// coinName: coin name, optional
// address: address
func (c *Cactus) GetSingleAddress(bId string, walletCode string, coinName constants.CactusToken, address string) (*GetSingleAddressResp, error) {
	query := map[string]string{}
	if coinName != "" {
		query["coin_name"] = string(coinName)
	}
	url := fmt.Sprintf(GetSingleAddressUrl, bId, walletCode, address)
	resp, err := c.get(url, query)
	if err != nil {
		return nil, err
	}
	var getSingleAddressResp GetSingleAddressResp
	err = json.Unmarshal(resp, &getSingleAddressResp)
	if err != nil {
		return nil, err
	}
	return &getSingleAddressResp, nil
}

type GetAddressListResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
		List   []struct {
			Address          string      `json:"address"`
			AddressStorage   string      `json:"address_storage"`
			AddressType      string      `json:"address_type"`
			AvailableAmount  interface{} `json:"available_amount"`
			CoinName         string      `json:"coin_name"`
			BchAddressFormat interface{} `json:"bch_address_format"`
			BId              string      `json:"b_id"`
			Description      string      `json:"description"`
			DomainId         string      `json:"domain_id"`
			FreezeAmount     interface{} `json:"freeze_amount"`
			TotalAmount      int         `json:"total_amount"`
			WalletCode       string      `json:"wallet_code"`
			WalletType       string      `json:"wallet_type"`
		} `json:"list"`
	} `json:"data"`
}

// GetAddressList gets the address list
// bId: business id
// walletCode: wallet code
// coinName: coin name, optional
// hideNoCoinAddress: hide addresses with no coin, optional
// keyword: keyword, optional
// offset: offset, optional, default 0
// limit: limit, optional, default 10
func (c *Cactus) GetAddressList(bId string, walletCode string, coinName constants.CactusToken, hideNoCoinAddress bool, keyword string, offset int, limit int) (*GetAddressListResp, error) {
	query := map[string]string{}
	if hideNoCoinAddress {
		query["hide_no_coin_address"] = "true"
	}
	if offset != 0 {
		query["offset"] = fmt.Sprintf("%d", offset)
	}
	if limit != 0 {
		query["limit"] = fmt.Sprintf("%d", limit)
	}
	if coinName != "" {
		query["coin_name"] = string(coinName)
	}
	if keyword != "" {
		query["keyword"] = keyword
	}
	url := fmt.Sprintf(GetAddressListUrl, bId, walletCode)
	resp, err := c.get(url, query)
	if err != nil {
		return nil, err
	}
	var getAddressListResp GetAddressListResp
	err = json.Unmarshal(resp, &getAddressListResp)
	if err != nil {
		return nil, err
	}
	return &getAddressListResp, nil
}

type EditAccountDescriptionResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

// EditAddressDescription edits the address description
// bId: business id
// walletCode: wallet code
// address: address
// description: description
func (c *Cactus) EditAddressDescription(bId string, walletCode string, address string, description string) (*EditAccountDescriptionResp, error) {
	url := fmt.Sprintf(EditAddressDescriptionUrl, bId, walletCode, address)
	req := map[string]interface{}{
		"description": description,
	}
	resp, err := c.post(url, req)
	if err != nil {
		return nil, err
	}
	var editAccountDescriptionResp EditAccountDescriptionResp
	err = json.Unmarshal(resp, &editAccountDescriptionResp)
	if err != nil {
		return nil, err
	}
	return &editAccountDescriptionResp, nil
}

// VerifyAddressResp verifies the address format
// data: bad address list
type VerifyAddressResp struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Successful bool     `json:"successful"`
	Data       []string `json:"data"`
}

// VerifyAddress verifies the address format
// coinName: coin name
// addresses: address list
func (c *Cactus) VerifyAddress(coinName constants.CactusToken, addresses []string) (*VerifyAddressResp, error) {
	params := map[string]interface{}{
		"coin_name": coinName,
		"addresses": addresses,
	}
	resp, err := c.post(VerifyAddressFormat, params)
	if err != nil {
		return nil, err
	}
	var verifyAddressResp VerifyAddressResp
	err = json.Unmarshal(resp, &verifyAddressResp)
	if err != nil {
		return nil, err
	}
	return &verifyAddressResp, nil
}
