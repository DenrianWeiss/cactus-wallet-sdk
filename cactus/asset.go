package cactus

import "encoding/json"

const (
	GetTotalAssetNotionalValueUrl   = "/custody/v1/api/history-asset"
	GetCurrentAssetNotionalValueUrl = "/custody/v1/api/asset"
)

type GetTotalAssetNotionalValueResp struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Successful interface{} `json:"successful"`
	Data       struct {
		HistoryAssetResult []struct {
			CreateTime     string `json:"create_time"`
			MarketValue    string `json:"market_value"`
			MarketValueCny string `json:"market_value_cny"`
		} `json:"history_asset_result"`
	} `json:"data"`
}

// GetTotalAssetNotionalValue gets total asset notional value
// bId: business id, optional, query all if not provided
func (c *Cactus) GetTotalAssetNotionalValue(bId string) (*GetTotalAssetNotionalValueResp, error) {
	var params map[string]string = nil
	if bId != "" {
		params = map[string]string{
			"b_id": bId,
		}
	}
	resp, err := c.get(GetTotalAssetNotionalValueUrl, params)
	if err != nil {
		return nil, err
	}
	var getTotalAssetNotionalValueResp GetTotalAssetNotionalValueResp
	err = json.Unmarshal(resp, &getTotalAssetNotionalValueResp)
	if err != nil {
		return nil, err
	}
	return &getTotalAssetNotionalValueResp, nil
}

type GetCurrentAssetNotionalValueResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ColdMarketValue     float64 `json:"cold_market_value"`
		ColdMarketValueCny  float64 `json:"cold_market_value_cny"`
		HotMarketValue      float64 `json:"hot_market_value"`
		HotMarketValueCny   float64 `json:"hot_market_value_cny"`
		TotalMarketValue    float64 `json:"total_market_value"`
		TotalMarketValueCny float64 `json:"total_market_value_cny"`
		Coins               []struct {
			CoinName  string  `json:"coin_name"`
			Amount    int64   `json:"amount"`
			Value     float64 `json:"value"`
			ValueCny  float64 `json:"value_cny"`
			StoreType string  `json:"store_type"`
		} `json:"coins"`
	} `json:"data"`
	Successful interface{} `json:"successful"`
}

func (c *Cactus) GetCurrentAssetNotionalValue(bId string) (*GetCurrentAssetNotionalValueResp, error) {
	var params map[string]string = nil
	if bId != "" {
		params = map[string]string{
			"b_id": bId,
		}
	}
	resp, err := c.get(GetCurrentAssetNotionalValueUrl, params)
	if err != nil {
		return nil, err
	}
	var getCurrentAssetNotionalValueResp GetCurrentAssetNotionalValueResp
	err = json.Unmarshal(resp, &getCurrentAssetNotionalValueResp)
	if err != nil {
		return nil, err
	}
	return &getCurrentAssetNotionalValueResp, nil
}
