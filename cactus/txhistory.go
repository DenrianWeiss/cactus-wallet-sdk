package cactus

import (
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/cactus-wallet-sdk/constants"
	"strconv"
)

const (
	GetWalletTransactionSummaryUrl = "/custody/v1/api/projects/%s/wallets/%s/tx-summaries"
	GetTransactionDetailsUrl       = "/custody/v1/api/projects/%s/wallets/%s/tx-details"
	EditTransactionRemarkUrl       = "/custody/v1/api/projects/%s/wallets/%s/details/%s"
)

type GetWalletTransactionSummaryResp struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Successful interface{} `json:"successful"`
	Data       struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		List   []struct {
			WalletCode      string `json:"wallet_code"`
			WalletType      string `json:"wallet_type"`
			CoinName        string `json:"coin_name"`
			OrderNo         string `json:"order_no"`
			BlockHeight     int    `json:"block_height"`
			TxId            string `json:"tx_id"`
			TxType          string `json:"tx_type"`
			Amount          int    `json:"amount"`
			WalletBalance   int    `json:"wallet_balance"`
			RemarkDetail    string `json:"remark_detail"`
			TxTimeStamp     int64  `json:"tx_time_stamp"`
			CreateTimeStamp int64  `json:"create_time_stamp"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}

// GetWalletTransactionHistory gets the wallet transaction history
// bId: business id
// walletCode: wallet code
// coinName: coin name
// txTypes: transaction types, see TxType, optional
// addresses: addresses, optional
// offset: offset, optional, default 0
// limit: limit, optional, default 10
// createTimeOrder: create time order, optional, default desc
// startTime: start time, optional
// endTime: end time, optional
func (c *Cactus) GetWalletTransactionHistory(
	bId string,
	walletCode string,
	coinName constants.CactusToken,
	txTypes []constants.TxType,
	addresses []string,
	offset int,
	limit int,
	createTimeOrder constants.OrderType,
	startTime int64,
	endTime int64,
) (*GetWalletTransactionSummaryResp, error) {
	query := map[string]string{}
	if offset != 0 {
		query["offset"] = fmt.Sprintf("%d", offset)
	}
	if limit != 0 {
		query["limit"] = fmt.Sprintf("%d", limit)
	}
	if coinName != "" {
		query["coin_name"] = string(coinName)
	}
	if createTimeOrder != constants.OrderTypeNotUsed {
		query["create_time_order"] = strconv.Itoa(int(createTimeOrder))
	}
	if startTime != 0 {
		query["start_time"] = fmt.Sprintf("%d", startTime)
	}
	if endTime != 0 {
		query["end_time"] = fmt.Sprintf("%d", endTime)
	}
	if len(txTypes) > 0 {
		txTypesString := ""
		for _, txType := range txTypes {
			txTypesString += string(txType) + ","
		}
		txTypesString = txTypesString[:len(txTypesString)-1]
		query["tx_types"] = txTypesString
	}
	if len(addresses) > 0 {
		addressesString := ""
		for _, address := range addresses {
			addressesString += address + ","
		}
		addressesString = addressesString[:len(addressesString)-1]
		query["addresses"] = addressesString
	}
	url := fmt.Sprintf(GetWalletTransactionSummaryUrl, bId, walletCode)
	resp, err := c.get(url, query)
	if err != nil {
		return nil, err
	}
	var getWalletTransactionSummaryResp GetWalletTransactionSummaryResp
	err = json.Unmarshal(resp, &getWalletTransactionSummaryResp)
	if err != nil {
		return nil, err
	}
	return &getWalletTransactionSummaryResp, nil
}

type GetTransactionDetailsResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		List   []struct {
			Id             int                  `json:"id"`
			DomainId       string               `json:"domain_id"`
			DomainName     interface{}          `json:"domain_name"`
			DomainCode     interface{}          `json:"domain_code"`
			KycNumber      interface{}          `json:"kyc_number"`
			ServiceType    interface{}          `json:"service_type"`
			WalletCode     string               `json:"wallet_code"`
			WalletType     constants.WalletType `json:"wallet_type"`
			CoinName       string               `json:"coin_name"`
			OrderNo        interface{}          `json:"order_no"`
			BlockHeight    int                  `json:"block_height"`
			ConfirmRatio   interface{}          `json:"confirm_ratio"`
			TxId           string               `json:"tx_id"`
			TxSize         int                  `json:"tx_size"`
			TxType         constants.TxType     `json:"tx_type"`
			WithdrawAmount interface{}          `json:"withdraw_amount"`
			GasPrice       interface{}          `json:"gas_price"`
			GasLimit       interface{}          `json:"gas_limit"`
			TxFee          int                  `json:"tx_fee"`
			TxFeeRate      interface{}          `json:"tx_fee_rate"`
			TxFeeType      interface{}          `json:"tx_fee_type"`
			MinerReward    interface{}          `json:"miner_reward"`
			MinerFee       interface{}          `json:"miner_fee"`
			DepositAmount  int                  `json:"deposit_amount"`
			WalletBalance  int                  `json:"wallet_balance"`
			ExtendedInfo   struct {
				DomainCoinBalance int64       `json:"domain_coin_balance"`
				Attachments       interface{} `json:"attachments"`
			} `json:"extended_info"`
			TxStatus     string      `json:"tx_status"`
			RemarkDetail interface{} `json:"remark_detail"`
			Vins         []struct {
				Address  string      `json:"address"`
				Idx      int         `json:"idx"`
				Tag      interface{} `json:"tag"`
				Amount   interface{} `json:"amount"`
				Balance  interface{} `json:"balance"`
				IsChange int         `json:"is_change"`
				Desc     interface{} `json:"desc"`
			} `json:"vins"`
			Vouts []struct {
				Address  string      `json:"address"`
				Idx      int         `json:"idx"`
				Tag      interface{} `json:"tag"`
				Amount   int         `json:"amount"`
				Balance  int         `json:"balance"`
				IsChange int         `json:"is_change"`
				Desc     string      `json:"desc"`
			} `json:"vouts"`
			TxTimeStamp     int64  `json:"tx_time_stamp"`
			CreateTimeStamp int64  `json:"create_time_stamp"`
			Bid             string `json:"bid"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
	Successful bool `json:"successful"`
}

// GetTransactionDetails gets the transaction details
// bId: business id
// walletCode: wallet code
// coinName: coin name, optional
// txTypes: transaction types, optional
// addresses: addresses, optional
// id: wallet detail item id, optional
// txId: transaction hash, optional
// orderNo: order number, optional
func (c *Cactus) GetTransactionDetails(
	bId string,
	walletCode string,
	coinName constants.CactusToken,
	txTypes []constants.TxType,
	addresses []string,
	id string,
	txId string,
	orderNo string,
	offset int,
	limit int,
	createTimeOrder constants.OrderType,
	startTime int64,
	endTime int64,
) (*GetTransactionDetailsResp, error) {
	query := map[string]string{}
	if offset != 0 {
		query["offset"] = fmt.Sprintf("%d", offset)
	}
	if limit != 0 {
		query["limit"] = fmt.Sprintf("%d", limit)
	}
	if coinName != "" {
		query["coin_name"] = string(coinName)
	}
	if createTimeOrder != constants.OrderTypeNotUsed {
		query["create_time_order"] = strconv.Itoa(int(createTimeOrder))
	}
	if startTime != 0 {
		query["start_time"] = fmt.Sprintf("%d", startTime)
	}
	if endTime != 0 {
		query["end_time"] = fmt.Sprintf("%d", endTime)
	}
	if len(txTypes) > 0 {
		txTypesString := ""
		for _, txType := range txTypes {
			txTypesString += string(txType) + ","
		}
		txTypesString = txTypesString[:len(txTypesString)-1]
		query["tx_types"] = txTypesString
	}
	if len(addresses) > 0 {
		addressesString := ""
		for _, address := range addresses {
			addressesString += address + ","
		}
		addressesString = addressesString[:len(addressesString)-1]
		query["addresses"] = addressesString
	}
	if id != "" {
		query["id"] = id
	}
	if txId != "" {
		query["tx_id"] = txId
	}
	if orderNo != "" {
		query["order_no"] = orderNo
	}
	url := fmt.Sprintf(GetTransactionDetailsUrl, bId, walletCode)
	resp, err := c.get(url, query)
	if err != nil {
		return nil, err
	}
	var getTransactionDetailsResp GetTransactionDetailsResp
	err = json.Unmarshal(resp, &getTransactionDetailsResp)
	if err != nil {
		return nil, err
	}
	return &getTransactionDetailsResp, nil
}

type EditTransactionRemarkResp struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}

// EditTransactionRemark edits the transaction remark
// bId: business id
// walletCode: wallet code
// id: wallet detail item id
// remark: remark
func (c *Cactus) EditTransactionRemark(bId string, walletCode string, id string, remark string) (*EditTransactionRemarkResp, error) {
	url := fmt.Sprintf(EditTransactionRemarkUrl, bId, walletCode, id)
	req := map[string]interface{}{
		"remark": remark,
	}
	resp, err := c.post(url, req)
	if err != nil {
		return nil, err
	}
	var editTransactionRemarkResp EditTransactionRemarkResp
	err = json.Unmarshal(resp, &editTransactionRemarkResp)
	if err != nil {
		return nil, err
	}
	return &editTransactionRemarkResp, nil
}
