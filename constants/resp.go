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
