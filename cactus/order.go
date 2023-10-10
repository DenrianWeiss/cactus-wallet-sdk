package cactus

import (
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
	"strconv"
)

const (
	GetFilteredOrderUrl = "/custody/v1/api/projects/%s/orders"
	GetOrderDetailsUrl  = "/custody/v1/api/projects/%s/orders/%s"
	ReplaceByFeeUrl     = "/custody/v1/api/projects/%s/orders/%s/accelerate"
	CancelOrderUrl      = "/custody/v1/api/projects/%s/orders/%s/cancel"
)

type GetFilteredOrderResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		List   []struct {
			DomainId                   string      `json:"domain_id"`
			ExchangeRate               float64     `json:"exchange_rate"`
			TimeStamp                  int64       `json:"time_stamp"`
			OrderNo                    string      `json:"order_no"`
			Applicant                  string      `json:"applicant"`
			BusinessName               string      `json:"business_name"`
			CoinName                   string      `json:"coin_name"`
			StorageType                string      `json:"storage_type"`
			WalletName                 string      `json:"wallet_name"`
			WalletCode                 string      `json:"wallet_code"`
			Amount                     int         `json:"amount"`
			OriginalAmount             int         `json:"original_amount"`
			MinerFeeRate               int         `json:"miner_fee_rate"`
			Description                string      `json:"description"`
			FromAddress                string      `json:"from_address"`
			GasPrice                   interface{} `json:"gas_price"`
			GasLimit                   interface{} `json:"gas_limit"`
			OrderDestAddressInfoVoList []struct {
				DestAddress   string `json:"dest_address"`
				MemoType      string `json:"memo_type"`
				Memo          string `json:"memo"`
				OriginBalance int    `json:"origin_balance"`
				Balance       int    `json:"balance"`
				Remark        string `json:"remark"`
			} `json:"order_dest_address_info_vo_list"`
			Status      string `json:"status"`
			InnerStatus string `json:"inner_status"`
			Bid         string `json:"bid"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

// GetFilteredOrder gets the filtered order
// bId: business id
// startTime: start time, optional
// endTime: end time, optional
// keyword: keyword, optional
// limit: limit, optional, default 10
// offset: offset, optional, default 0
// sortByTime: sort by time, optional, default desc
// applicant: applicant who made the request, optional
// coinName: coin name, optional
// chainName: chain name, optional
// walletName: wallet name, optional
// status: status, optional
func (c *Cactus) GetFilteredOrder(
	bId string,
	applicant []string,
	coinName []constants.CactusToken,
	chainName []constants.ChainName,
	walletName []string,
	status []string,
	keyword string,
	sortByTime constants.OrderType,
	startTime int,
	endTime int,
	offset int,
	limit int,
) (*GetFilteredOrderResp, error) {
	params := map[string]string{}
	if offset != 0 {
		params["offset"] = fmt.Sprintf("%d", offset)
	}
	if limit != 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}
	if keyword != "" {
		params["keyword"] = keyword
	}
	if startTime != 0 {
		params["start_time"] = fmt.Sprintf("%d", startTime)
	}
	if endTime != 0 {
		params["end_time"] = fmt.Sprintf("%d", endTime)
	}
	if sortByTime != constants.OrderTypeNotUsed {
		params["sort_by_time"] = strconv.Itoa(int(sortByTime))
	}
	if len(applicant) > 0 {
		applicantString := ""
		for _, app := range applicant {
			applicantString += app + ","
		}
		applicantString = applicantString[:len(applicantString)-1]
		params["applicant"] = applicantString
	}
	if len(coinName) > 0 {
		coinNameString := ""
		for _, coin := range coinName {
			coinNameString += string(coin) + ","
		}
		coinNameString = coinNameString[:len(coinNameString)-1]
		params["coin_name"] = coinNameString
	}
	if len(chainName) > 0 {
		chainNameString := ""
		for _, chain := range chainName {
			chainNameString += string(chain) + ","
		}
		chainNameString = chainNameString[:len(chainNameString)-1]
		params["chain_name"] = chainNameString
	}
	if len(walletName) > 0 {
		walletNameString := ""
		for _, wallet := range walletName {
			walletNameString += wallet + ","
		}
		walletNameString = walletNameString[:len(walletNameString)-1]
		params["wallet_name"] = walletNameString
	}
	if len(status) > 0 {
		statusString := ""
		for _, stat := range status {
			statusString += stat + ","
		}
		statusString = statusString[:len(statusString)-1]
		params["status"] = statusString
	}
	path := fmt.Sprintf(GetFilteredOrderUrl, bId)
	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}
	var getFilteredOrderResp GetFilteredOrderResp
	err = json.Unmarshal(resp, &getFilteredOrderResp)
	if err != nil {
		return nil, err
	}
	return &getFilteredOrderResp, nil
}

type GetOrderDetailsResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderWalletInfo struct {
			BusinessName               string      `json:"business_name"`
			CoinName                   string      `json:"coin_name"`
			WalletType                 string      `json:"wallet_type"`
			StorageType                string      `json:"storage_type"`
			WalletName                 string      `json:"wallet_name"`
			WalletCode                 string      `json:"wallet_code"`
			Timestamp                  int64       `json:"timestamp"`
			OrderNo                    string      `json:"order_no"`
			Applicant                  string      `json:"applicant"`
			Amount                     int         `json:"amount"`
			ExchangeRate               int         `json:"exchange_rate"`
			Value                      interface{} `json:"value"`
			MinerFeeRate               interface{} `json:"miner_fee_rate"`
			Description                string      `json:"description"`
			Status                     string      `json:"status"`
			InnerStatus                string      `json:"inner_status"`
			MinerFee                   int64       `json:"miner_fee"`
			FromAddress                string      `json:"from_address"`
			GasPrice                   int64       `json:"gas_price"`
			GasLimit                   int         `json:"gas_limit"`
			OrderDestAddressInfoVoList []struct {
				DestAddress   string `json:"dest_address"`
				MemoType      string `json:"memo_type"`
				Memo          string `json:"memo"`
				OriginBalance int    `json:"origin_balance"`
				Balance       int    `json:"balance"`
				Remark        string `json:"remark"`
			} `json:"order_dest_address_info_vo_list"`
			Bid string `json:"bid"`
		} `json:"order_wallet_info"`
		TxInfoModels []struct {
			TxType      string `json:"tx_type"`
			BlockHeight int    `json:"block_height"`
			TxSize      int    `json:"tx_size"`
			TxHash      string `json:"tx_hash"`
			GasPrice    int64  `json:"gas_price"`
			GasLimit    int    `json:"gas_limit"`
			MinerFee    string `json:"miner_fee"`
		} `json:"tx_info_models"`
		ConsolidationTxInfoModels []struct {
			TxType      string `json:"tx_type"`
			BlockHeight int    `json:"block_height"`
			TxSize      int    `json:"tx_size"`
			TxHash      string `json:"tx_hash"`
			GasPrice    int64  `json:"gas_price"`
			GasLimit    int    `json:"gas_limit"`
			MinerFee    string `json:"miner_fee"`
		} `json:"consolidation_tx_info_models"`
		MinerFeeTxInfoModels []struct {
			TxType      string      `json:"tx_type"`
			BlockHeight int         `json:"block_height"`
			TxSize      int         `json:"tx_size"`
			TxHash      string      `json:"tx_hash"`
			GasPrice    int64       `json:"gas_price"`
			GasLimit    interface{} `json:"gas_limit"`
			MinerFee    string      `json:"miner_fee"`
		} `json:"miner_fee_tx_info_models"`
		PartialFailed  []interface{} `json:"partial_failed"`
		PartialSuccess []interface{} `json:"partial_success"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

func (c *Cactus) GetOrderDetails(bId string, orderNo string) (*GetOrderDetailsResp, error) {
	path := fmt.Sprintf(GetOrderDetailsUrl, bId, orderNo)
	resp, err := c.get(path, nil)
	if err != nil {
		return nil, err
	}
	var getOrderDetailsResp GetOrderDetailsResp
	err = json.Unmarshal(resp, &getOrderDetailsResp)
	if err != nil {
		return nil, err
	}
	return &getOrderDetailsResp, nil
}

type ReplaceByFeeResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Data       int64  `json:"data"` // GasPrice
	Successful bool   `json:"successful"`
}

// ReplaceByFee replaces the order by fee
// bId: business id
// orderNo: order no
// level: replace by fee level
// gasPrice: gas price, required if level is custom
func (c *Cactus) ReplaceByFee(bId string, orderNo string, level constants.ReplaceByFeeLevel, gasPrice float64) (*ReplaceByFeeResp, error) {
	params := map[string]interface{}{
		"level": level,
	}
	if gasPrice != 0 {
		params["gas_price"] = gasPrice
	}
	path := fmt.Sprintf(ReplaceByFeeUrl, bId, orderNo)
	resp, err := c.post(path, params)
	if err != nil {
		return nil, err
	}
	var replaceByFeeResp ReplaceByFeeResp
	err = json.Unmarshal(resp, &replaceByFeeResp)
	if err != nil {
		return nil, err
	}
	return &replaceByFeeResp, nil
}

type CancelOrderResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Data       int64  `json:"data"` // GasPrice
	Successful bool   `json:"successful"`
}

// CancelOrder cancels the order
// bId: business id
// orderNo: order no
func (c *Cactus) CancelOrder(bId string, orderNo string, level constants.ReplaceByFeeLevel, gasPrice float64) (*CancelOrderResp, error) {
	path := fmt.Sprintf(CancelOrderUrl, bId, orderNo)
	param := map[string]interface{}{
		"level": level,
	}
	if gasPrice != 0 {
		param["gas_price"] = gasPrice
	}
	resp, err := c.post(path, param)
	if err != nil {
		return nil, err
	}
	var cancelOrderResp CancelOrderResp
	err = json.Unmarshal(resp, &cancelOrderResp)
	if err != nil {
		return nil, err
	}
	return &cancelOrderResp, nil
}
