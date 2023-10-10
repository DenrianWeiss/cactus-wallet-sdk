package constants

type WalletType string

const (
	WalletTypeMixed      WalletType = "MIXED_ADDRESS"
	WalletTypeSegregated WalletType = "SEGREGATED_ADDRESS"
	WalletTypeDapp       WalletType = "DAPP_ADDRESS"
	WalletTypeDappCustom WalletType = "DAPP_CUSTOM_ADDRESS"
)

type StorageType string

const (
	StorageTypeCold StorageType = "PRIME_COLD_LV1"
	StorageTypeHot  StorageType = "PRIME_HOT"
)

type WalletFilterType string

const (
	WalletFilterTypeAll  WalletFilterType = ""
	WalletFilterTypeHot  WalletFilterType = "HOT"
	WalletFilterTypeCold WalletFilterType = "COLD"
)

type OrderType int

const (
	OrderTypeDesc OrderType = iota
	OrderTypeAsc

	OrderTypeNotUsed OrderType = -1
)

type FeeLevelType string

const (
	FeeLevelTypeLow    FeeLevelType = "LOW"
	FeeLevelTypeNormal FeeLevelType = "NORMAL"
	FeeLevelTypeHigh   FeeLevelType = "HIGHER"
	FeeLevelTypeCustom FeeLevelType = "CUSTOM"
)

type MemoType string

const (
	MemoTypeNone MemoType = "MEMO_NONE"
	MemoTypeText MemoType = "MEMO_TEXT"
	MemoTypeHash MemoType = "MEMO_HASH"
)

type TxType string

// WITHDRAW, DEPOSIT, PREFUND_DEPOSIT, MINER_REWARD, FEE_PAYMENT, ADVANCED_FEE, CONSOLIDATION, DEPOSIT_ROLLBACK, WITHDRAW_CREDIT, WITHDRAW_FAILED, WITHDRAW_CANCELED, WITHDRAW_REFUND
const (
	TxTypeWithdraw         TxType = "WITHDRAW"
	TxTypeDeposit          TxType = "DEPOSIT"
	TxTypePrefundDeposit   TxType = "PREFUND_DEPOSIT"
	TxTypeMinerReward      TxType = "MINER_REWARD"
	TxTypeFeePayment       TxType = "FEE_PAYMENT"
	TxTypeAdvancedFee      TxType = "ADVANCED_FEE"
	TxTypeConsolidation    TxType = "CONSOLIDATION"
	TxTypeDepositRollback  TxType = "DEPOSIT_ROLLBACK"
	TxTypeWithdrawCredit   TxType = "WITHDRAW_CREDIT"
	TxTypeWithdrawFailed   TxType = "WITHDRAW_FAILED"
	TxTypeWithdrawCanceled TxType = "WITHDRAW_CANCELED"
	TxTypeWithdrawRefund   TxType = "WITHDRAW_REFUND"
	TxTypeRebase           TxType = "REBASE"
	TxTypeNone             TxType = ""
)

type ReplaceByFeeLevel string

const (
	ReplaceByFeeLevel1 ReplaceByFeeLevel = "LEVEL1"
	ReplaceByFeeLevel2 ReplaceByFeeLevel = "LEVEL2"
	ReplaceByFeeLevel3 ReplaceByFeeLevel = "LEVEL3"
	ReplaceByFeeCustom ReplaceByFeeLevel = "CUSTOM"
)

type SignatureVersion string

const (
	SignatureVersionV4       SignatureVersion = "V4"
	SignatureVersionPersonal SignatureVersion = "personalSign"
)
