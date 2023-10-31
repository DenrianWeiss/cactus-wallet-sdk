package cactus

import (
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
)

const (
	EstimateWithdrawalFeeUrl = "/custody/v1/api/projects/%s/estimate-miner-fee"
	CreateWithdrawalOrderUrl = "/custody/v1/api/projects/%s/order/create"
	GetWithdrawalFeeRangeUrl = "/custody/v1/api/customize-fee-rate/range"
	GetWithdrawalRateUrl     = "/custody/v1/api/recommend-fee-rate/list"
)

type EstimateWithdrawalFeeResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Data       int    `json:"data"`
	Successful bool   `json:"successful"`
}

type WithdrawalArgsFeeReq struct {
	FromAddress         string                 `json:"from_address,omitempty"`
	FromWalletCode      string                 `json:"from_wallet_code"`
	CoinName            constants.CactusToken  `json:"coin_name,omitempty"`
	OrderNo             string                 `json:"order_no,omitempty"`
	Description         string                 `json:"description,omitempty"`
	FeeRateLevel        constants.FeeLevelType `json:"fee_rate_level,omitempty"`
	FeeRate             float64                `json:"fee_rate,omitempty"`
	DestAddressItemList DestAddressItem        `json:"dest_address_item_list"`
}

type DestAddressItem struct {
	Amount          int                `json:"amount"`
	DestAddress     string             `json:"dest_address"`
	MemoType        string             `json:"memo_type,omitempty"`
	Memo            constants.MemoType `json:"memo,omitempty"`
	IsAllWithdrawal bool               `json:"is_all_withdrawal"`
	Remark          string             `json:"remark,omitempty"`
}

// EstimateWithdrawalFee estimates the withdrawal fee
// bId: business id
// req: request body
func (c *Cactus) EstimateWithdrawalFee(bId string, req WithdrawalArgsFeeReq) (*EstimateWithdrawalFeeResp, error) {
	path := fmt.Sprintf(EstimateWithdrawalFeeUrl, bId)
	// serialize req and deserialize to map
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var reqMap map[string]interface{}
	err = json.Unmarshal(reqBytes, &reqMap)
	if err != nil {
		return nil, err
	}
	resp, err := c.post(path, reqMap)
	if err != nil {
		return nil, err
	}
	var estimateWithdrawalFeeResp EstimateWithdrawalFeeResp
	err = json.Unmarshal(resp, &estimateWithdrawalFeeResp)
	if err != nil {
		return nil, err
	}
	return &estimateWithdrawalFeeResp, nil
}

type CreateWithdrawalOrderResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		OrderNo string `json:"OrderNo"`
	} `json:"data"`
}

// CreateWithdrawOrder creates the withdrawal order
// bId: business id
// req: request body
func (c *Cactus) CreateWithdrawOrder(bId string, req WithdrawalArgsFeeReq) (*CreateWithdrawalOrderResp, error) {
	path := fmt.Sprintf(CreateWithdrawalOrderUrl, bId)
	// serialize req and deserialize to map
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var reqMap map[string]interface{}
	err = json.Unmarshal(reqBytes, &reqMap)
	if err != nil {
		return nil, err
	}
	resp, err := c.post(path, reqMap)
	if err != nil {
		return nil, err
	}
	var createWithdrawalOrderResp CreateWithdrawalOrderResp
	err = json.Unmarshal(resp, &createWithdrawalOrderResp)
	if err != nil {
		return nil, err
	}
	return &createWithdrawalOrderResp, nil
}

type GetWithdrawalFeeRangeResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		MaxFeeRate int `json:"maxFeeRate"`
		MinFeeRate int `json:"minFeeRate"`
	} `json:"data"`
}

func (c *Cactus) GetWithdrawalFeeRange(coinName constants.CactusToken) (*GetWithdrawalFeeRangeResp, error) {
	params := map[string]string{
		"coin_name": string(coinName),
	}
	resp, err := c.get(GetWithdrawalFeeRangeUrl, params)
	if err != nil {
		return nil, err
	}
	var getWithdrawalFeeRangeResp GetWithdrawalFeeRangeResp
	err = json.Unmarshal(resp, &getWithdrawalFeeRangeResp)
	if err != nil {
		return nil, err
	}
	return &getWithdrawalFeeRangeResp, nil
}

type GetWithdrawalRateResp struct {
	Code       int     `json:"code"`
	Message    string  `json:"message"`
	Successful bool    `json:"successful"`
	Data       []int64 `json:"data"`
}

func (c *Cactus) GetWithdrawalRate(coinName constants.CactusToken) (*GetWithdrawalRateResp, error) {
	params := map[string]string{
		"coin_name": string(coinName),
	}
	resp, err := c.get(GetWithdrawalRateUrl, params)
	if err != nil {
		return nil, err
	}
	var getWithdrawalRateResp GetWithdrawalRateResp
	err = json.Unmarshal(resp, &getWithdrawalRateResp)
	if err != nil {
		return nil, err
	}
	return &getWithdrawalRateResp, nil
}
