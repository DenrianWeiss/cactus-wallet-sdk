package cactus

type CreateTokenReq struct {
	GrantType    string `json:"grantType"`
	RefreshToken string `json:"refreshToken"`
}

type CreateTokenResp struct {
	AccessToken string `json:"jwt"`
}

type GetEthAccountsResp struct {
	Name             string   `json:"name"`
	Address          string   `json:"address"`
	Labels           []string `json:"labels"`
	Balance          string   `json:"balance"`
	ChainId          int      `json:"chainId"`
	CustodianDetails struct {
		DomainId  string `json:"domainId"`
		ProjectId string `json:"projectId"`
		WalletId  string `json:"walletId"`
	} `json:"custodianDetails"`
}

type SendTransactionReq struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	GasLimit             string `json:"gasLimit"`
	Value                string `json:"value"`
	Data                 string `json:"data"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
}

type SendTransactionResp struct {
	TransactionStatus      string `json:"transactionStatus"`
	TransactionHash        string `json:"transactionHash"`
	CustodianTransactionId string `json:"custodian_transactionId"`
	GasPrice               string `json:"gasPrice"`
	MaxFeePerGas           string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas   string `json:"maxPriorityFeePerGas"`
	GasLimit               string `json:"gasLimit"`
	Nonce                  string `json:"nonce"`
	From                   string `json:"from"`
	Signature              string `json:"signature"`
}

type GetTransactionEntry struct {
	TransactionStatus      string `json:"transactionStatus"`
	TransactionHash        string `json:"transactionHash"`
	CustodianTransactionId string `json:"custodian_transactionId"`
	GasPrice               string `json:"gasPrice"`
	MaxFeePerGas           string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas   string `json:"maxPriorityFeePerGas"`
	GasLimit               string `json:"gasLimit"`
	Nonce                  string `json:"nonce"`
	From                   string `json:"from"`
	Signature              string `json:"signature"`
}

type GetTransactionResp []GetTransactionEntry
