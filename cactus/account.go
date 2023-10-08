package cactus

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func (c *Cactus) GetEthAccounts(chainId int, account ...string) (resp *GetEthAccountsResp, err error) {
	params := map[string]string{
		"chainId": strconv.Itoa(chainId),
	}
	if len(account) > 0 {
		params["account"] = account[0]
	}
	get, err := c.get("/eth-accounts", params)
	if err != nil {
		return nil, err
	}
	response := &GetEthAccountsResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// CreateTransaction creates a transaction
// chainId: required
// req: transaction details
func (c *Cactus) CreateTransaction(chainId string, req *SendTransactionReq) (resp *SendTransactionResp, err error) {
	r, _ := json.Marshal(req)
	post, err := c.post("/transactions", map[string]string{"chainId": chainId}, bytes.NewReader(r))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(post, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransaction returns transaction by hash or other conditions
// chainId: required
// from: optional, filter from address
// custodianTxId(transactionId): optional, filtering transactions by transactionId, this transactionId is the custodian_transactionId returned by Cactus Custody.
// transactionHash: optional, filter by transaction hash
func (c *Cactus) GetTransaction(chainId string, from string, custodianTxId string, transactionHash string) (resp *GetTransactionResp, err error) {
	params := map[string]string{
		"chainId": chainId,
	}
	if from != "" {
		params["from"] = from
	}
	if custodianTxId != "" {
		params["custodianTxId"] = custodianTxId
	}
	if transactionHash != "" {
		params["transactionHash"] = transactionHash
	}
	get, err := c.get("/transactions", params)
	if err != nil {
		return nil, err
	}
	response := &GetTransactionResp{}
	err = json.Unmarshal(get, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
