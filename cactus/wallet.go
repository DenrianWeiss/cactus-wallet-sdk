package cactus

import (
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
	"strconv"
)

const (
	GetSingleWalletInfoUri = "/custody/v1/api/projects/%s/wallets/%s"
	GetWalletListUrl       = "/custody/v1/api/wallets"
	GetCoinInfoUrl         = "/custody/v1/api/coin-infos"
	GetChainInfoUrl        = "/custody/v1/api/chain-infos"
	CreateWalletUrl        = "/custody/v1/api/projects/%s/wallets/create"
)

type GetSingleWalletInfoResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		DomainId              string               `json:"domain_id"`
		BId                   string               `json:"b_id"`
		WalletCode            string               `json:"wallet_code"`
		WalletName            string               `json:"wallet_name"`
		CoinName              string               `json:"coin_name"`
		WalletType            constants.WalletType `json:"wallet_type"`
		StorageType           string               `json:"storage_type"`
		AvailableAmount       int                  `json:"available_amount"`
		FreezeAmount          int                  `json:"freeze_amount"`
		TotalAmount           int                  `json:"total_amount"`
		UsdTotalMarket        float64              `json:"usd_total_market"`
		CnyTotalMarket        float64              `json:"cny_total_market"`
		CoinStatus            string               `json:"coin_status"`
		ChineseReasonOfStatus string               `json:"chinese_reason_of_status"`
		EnglishReasonOfStatus string               `json:"english_reason_of_status"`
		NormalAddressLimit    int                  `json:"normal_address_limit"`
		NormalAddressNum      int                  `json:"normal_address_num"`
		CreateTime            int64                `json:"create_time"`
	} `json:"data"`
}

// GetSingleWalletInfo gets the single wallet info
// bId: the business id
// walletCode: the wallet code
// coinName: the coin name, optional, see constants/tokens.go
func (c *Cactus) GetSingleWalletInfo(bId string, walletCode string, coinName constants.CactusToken) (*GetSingleWalletInfoResp, error) {
	path := fmt.Sprintf(GetSingleWalletInfoUri, bId, walletCode)
	param := map[string]string{}
	if coinName != "" {
		param["coin_name"] = string(coinName)
	}
	resp, err := c.get(path, param)
	if err != nil {
		return nil, err
	}
	var getSingleWalletInfoResp GetSingleWalletInfoResp
	err = json.Unmarshal(resp, &getSingleWalletInfoResp)
	if err != nil {
		return nil, err
	}
	return &getSingleWalletInfoResp, nil
}

type GetWalletListResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
	Data       struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		List   []struct {
			DomainId              string                `json:"domain_id"`
			BId                   string                `json:"b_id"`
			WalletCode            string                `json:"wallet_code"`
			WalletName            string                `json:"wallet_name"`
			CoinName              string                `json:"coin_name"`
			ContractAddress       interface{}           `json:"contract_address"`
			WalletType            constants.WalletType  `json:"wallet_type"`
			StorageType           constants.StorageType `json:"storage_type"`
			AvailableAmount       int                   `json:"available_amount"`
			FreezeAmount          int                   `json:"freeze_amount"`
			TotalAmount           int                   `json:"total_amount"`
			UsdTotalMarket        float64               `json:"usd_total_market"`
			CnyTotalMarket        float64               `json:"cny_total_market"`
			CoinStatus            string                `json:"coin_status"`
			ChineseReasonOfStatus string                `json:"chinese_reason_of_status"`
			EnglishReasonOfStatus string                `json:"english_reason_of_status"`
			NormalAddressLimit    int                   `json:"normal_address_limit"`
			NormalAddressNum      int                   `json:"normal_address_num"`
			CreateTime            int64                 `json:"create_time"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}

// GetWalletList gets the wallet list
// All args are optional
func (c *Cactus) GetWalletList(bId string,
	walletFilterType constants.WalletFilterType,
	hideNoCoinWallet bool,
	coinNames []constants.CactusToken,
	walletTypes constants.WalletType,
	keyword string,
	defiWalletCode string,
	mainWalletCode string,
	chain constants.ChainName,
	totalMarketOrder constants.OrderType,
	createTimeOrder constants.OrderType,
	offset int,
	limit int) (*GetWalletListResp, error) {
	params := map[string]string{}
	// Probe every arg and if present, add it to the params map
	if bId != "" {
		params["b_id"] = bId
	}
	if walletFilterType != "" {
		params["type"] = string(walletFilterType)
	}
	if hideNoCoinWallet {
		params["hide_no_coin_wallet"] = "true"
	}
	if len(coinNames) > 0 {
		// Create comma seperated string of coin names
		coinNamesString := ""
		for _, coinName := range coinNames {
			coinNamesString += string(coinName) + ","
		}
		// Remove trailing comma
		coinNamesString = coinNamesString[:len(coinNamesString)-1]
		params["coin_names"] = coinNamesString
	}
	if walletTypes != "" {
		params["wallet_types"] = string(walletTypes)
	}
	if keyword != "" {
		params["keyword"] = keyword
	}
	if defiWalletCode != "" {
		params["defi_wallet_code"] = defiWalletCode
	}
	if mainWalletCode != "" {
		params["main_wallet_code"] = mainWalletCode
	}
	if chain != "" {
		params["chain"] = string(chain)
	}
	if totalMarketOrder != constants.OrderTypeNotUsed {
		params["total_market_order"] = strconv.Itoa(int(totalMarketOrder))
	}
	if createTimeOrder != constants.OrderTypeNotUsed {
		params["create_time_order"] = strconv.Itoa(int(createTimeOrder))
	}
	if offset != 0 {
		params["offset"] = strconv.Itoa(offset)
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	resp, err := c.get(GetWalletListUrl, params)
	if err != nil {
		return nil, err
	}
	var getWalletListResp GetWalletListResp
	err = json.Unmarshal(resp, &getWalletListResp)
	if err != nil {
		return nil, err
	}
	return &getWalletListResp, nil
}

type GetCoinInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		CactusSymbol       string `json:"cactus_symbol"`
		Symbol             string `json:"symbol"`
		Chain              string `json:"chain"`
		CactusChain        string `json:"cactus_chain"`
		Decimals           string `json:"decimals"`
		ContractAddress    string `json:"contract_address"`
		DepositBlockNumber string `json:"deposit_block_number"`
		ConfirmBlockNumber string `json:"confirm_block_number"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

func (c *Cactus) GetCoinInfo(cactusToken string, token string) (*GetCoinInfoResp, error) {
	params := map[string]string{}
	if cactusToken != "" {
		params["cactus_symbol"] = cactusToken
	}
	if token != "" {
		params["symbol"] = token
	}
	resp, err := c.get(GetCoinInfoUrl, params)
	if err != nil {
		return nil, err
	}
	var getCoinInfoResp GetCoinInfoResp
	err = json.Unmarshal(resp, &getCoinInfoResp)
	if err != nil {
		return nil, err
	}
	return &getCoinInfoResp, nil
}

type GetChainInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Chain              string `json:"chain"`
		FullName           string `json:"full_name"`
		MainCoin           string `json:"main_coin"`
		EvmChain           bool   `json:"evm_chain"`
		SupportEip1559     bool   `json:"support_eip1559"`
		ConfirmBlockNumber int    `json:"confirm_block_number"`
		MinerBlockNumber   int    `json:"miner_block_number"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

// GetChainInfo gets the chain info
func (c *Cactus) GetChainInfo(chain string, fullName string) (*GetChainInfoResp, error) {
	params := map[string]string{}
	if chain != "" {
		params["chain"] = chain
	}
	if fullName != "" {
		params["full_name"] = fullName
	}
	resp, err := c.get(GetChainInfoUrl, params)
	if err != nil {
		return nil, err
	}
	var getChainInfoResp GetChainInfoResp
	err = json.Unmarshal(resp, &getChainInfoResp)
	if err != nil {
		return nil, err
	}
	return &getChainInfoResp, nil
}

type CreateWalletResp struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Data       []string `json:"data"`
	Successful bool     `json:"successful"`
}

// CreateWallet creates a defi wallet
// bId: the business id
// walletType: the wallet type, only "DEFI" is allowed
// number: the number of wallets to create
func (c *Cactus) CreateWallet(bId string, walletType string, number int) (*CreateWalletResp, error) {
	path := fmt.Sprintf(CreateWalletUrl, bId)
	params := map[string]interface{}{
		"wallet_type": walletType,
		"number":      number,
	}
	resp, err := c.post(path, params)
	if err != nil {
		return nil, err
	}
	var createWalletResp CreateWalletResp
	err = json.Unmarshal(resp, &createWalletResp)
	if err != nil {
		return nil, err
	}
	return &createWalletResp, nil
}
