package cactus

import (
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
	"strconv"
)

const (
	CreateContractOrderUrl       = "/custody/v1/api/projects/%s/contract/call"
	CreateSignOrderUrl           = "/custody/v1/api/projects/%s/wallets/%s/signatures"
	GetDefiTransactionHistoryUrl = "/custody/v1/api/projects/%s/wallets/%s/contract/orders"
	GetDefiTransactionDetailsUrl = "/custody/v1/api/projects/%s/wallets/%s/contract/orders/%s"
)

type CreateContractOrderReq struct {
	OrderNo              string              `json:"order_no"`
	FromWalletCode       string              `json:"from_wallet_code"`
	FromAddress          string              `json:"from_address"`
	ToAddress            string              `json:"to_address"`
	Amount               string              `json:"amount"`
	Chain                constants.ChainName `json:"chain"`
	ContractData         string              `json:"contract_data"`
	GasPriceLevel        string              `json:"gas_price_level,omitempty"`
	GasPrice             int64               `json:"gas_price,omitempty"`
	MaxFeePerGas         int64               `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas int64               `json:"max_priority_fee_per_gas,omitempty"`
	GasLimit             int                 `json:"gas_limit"`
	Description          string              `json:"description"`
}

type CreateContractOrderResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderNo string `json:"OrderNo"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

// CreateContractOrder creates a contract order
// bId: business id
// req: request body
func (c *Cactus) CreateContractOrder(bId string, req CreateContractOrderReq) (*CreateContractOrderResp, error) {
	path := fmt.Sprintf(CreateContractOrderUrl, bId)
	var reqMap map[string]interface{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBytes, &reqMap)
	if err != nil {
		return nil, err
	}
	resp, err := c.post(path, reqMap)
	if err != nil {
		return nil, err
	}
	var createContractOrderResp CreateContractOrderResp
	err = json.Unmarshal(resp, &createContractOrderResp)
	if err != nil {
		return nil, err
	}
	return &createContractOrderResp, nil
}

type CreateSignOrderReq struct {
	Address          string                     `json:"address"`
	SignatureVersion constants.SignatureVersion `json:"signature_version"`
	Payload          interface{}                `json:"payload"`
	Chain            constants.ChainName        `json:"chain"`
	OrderNo          string                     `json:"order_no"`
	Description      string                     `json:"description,omitempty"`
}

type CreateSignOrderResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		OrderNo string `json:"OrderNo"`
	} `json:"data"`
}

func (c *Cactus) CreateSignOrder(bId string, walletCode string, req CreateSignOrderReq) (*CreateSignOrderResp, error) {
	path := fmt.Sprintf(CreateSignOrderUrl, bId, walletCode)
	var reqMap map[string]interface{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBytes, &reqMap)
	if err != nil {
		return nil, err
	}
	resp, err := c.post(path, reqMap)
	if err != nil {
		return nil, err
	}
	var createSignOrderResp CreateSignOrderResp
	err = json.Unmarshal(resp, &createSignOrderResp)
	if err != nil {
		return nil, err
	}
	return &createSignOrderResp, nil
}

type GetTransactionHistoryResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		List   []struct {
			TimeStamp       int64       `json:"time_stamp"`
			MinerFee        int64       `json:"miner_fee"`
			OrderNo         string      `json:"order_no"`
			ContractAddress string      `json:"contract_address"`
			GasPrice        int64       `json:"gas_price"`
			GasLimit        int         `json:"gas_limit"`
			ContractData    string      `json:"contract_data"`
			Applicant       string      `json:"applicant"`
			Status          string      `json:"status"`
			Amount          int64       `json:"amount"`
			Description     interface{} `json:"description"`
			TxId            string      `json:"tx_id"`
			DepositTrans    []struct {
				Amount   int64  `json:"amount"`
				CoinName string `json:"coin_name"`
			} `json:"deposit_trans"`
			WithdrawTrans []struct {
				Amount   int64  `json:"amount"`
				CoinName string `json:"coin_name"`
			} `json:"withdraw_trans"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

// GetTransactionHistory gets the transaction history
// bId: business id
// walletCode: wallet code
// keyword: keyword, optional
// sortByTime: ASC or DESC, optional
// status: status, optional
// chain: chain
// startTime: start time, optional
// limit: limit, optional
// offset: offset, optional
func (c *Cactus) GetTransactionHistory(bId string, walletCode string, keyword string, sortByTime string, status string, chain constants.ChainName, startTime int64, limit int, offset int) (*GetTransactionHistoryResp, error) {
	query := map[string]string{
		"chain": string(chain),
	}
	if keyword != "" {
		query["keyword"] = keyword
	}
	if sortByTime != "" {
		query["sort_by_time"] = sortByTime
	}
	if status != "" {
		query["status"] = status
	}
	if startTime != 0 {
		query["start_time"] = strconv.FormatInt(startTime, 10)
	}
	if limit != 0 {
		query["limit"] = strconv.Itoa(limit)
	}
	if offset != 0 {
		query["offset"] = strconv.Itoa(offset)
	}
	url := fmt.Sprintf(GetDefiTransactionHistoryUrl, bId, walletCode)
	resp, err := c.get(url, query)
	if err != nil {
		return nil, err
	}
	var getTransactionHistoryResp GetTransactionHistoryResp
	err = json.Unmarshal(resp, &getTransactionHistoryResp)
	if err != nil {
		return nil, err
	}
	return &getTransactionHistoryResp, nil
}

type GetDefiTransactionDetailsResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TimeStamp        int64       `json:"time_stamp"`
		MinerFee         int64       `json:"miner_fee"`
		OrderNo          string      `json:"order_no"`
		ContractAddress  string      `json:"contract_address"`
		ContractFunction string      `json:"contract_function"`
		GasPrice         int64       `json:"gas_price"`
		GasLimit         int         `json:"gas_limit"`
		ContractData     string      `json:"contract_data"`
		Applicant        string      `json:"applicant"`
		Status           string      `json:"status"`
		Amount           int         `json:"amount"`
		TxId             string      `json:"tx_id"`
		Description      interface{} `json:"description"`
		DepositTrans     []struct {
			Amount   int64  `json:"amount"`
			CoinName string `json:"coin_name"`
		} `json:"deposit_trans"`
		WithdrawTrans []interface{} `json:"withdraw_trans"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

func (c *Cactus) GetDefiTransactionDetails(bId string, walletCode string, orderNo string) (*GetDefiTransactionDetailsResp, error) {
	url := fmt.Sprintf(GetDefiTransactionDetailsUrl, bId, walletCode, orderNo)
	resp, err := c.get(url, nil)
	if err != nil {
		return nil, err
	}
	var getDefiTransactionDetailsResp GetDefiTransactionDetailsResp
	err = json.Unmarshal(resp, &getDefiTransactionDetailsResp)
	if err != nil {
		return nil, err
	}
	return &getDefiTransactionDetailsResp, nil
}
