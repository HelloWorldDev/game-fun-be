package response

import (
	"github.com/shopspring/decimal"
)

type SwapRouteData struct {
	Quote        Quote           `json:"quote"`
	RawTx        RawTx           `json:"raw_tx"`
	AmountInUSD  decimal.Decimal `json:"amountInUSD"`
	AmountOutUSD decimal.Decimal `json:"amountOutUSD"`
	PlatformType string          `form:"platform_type" binding:"required,oneof=pump raydium game g_external g_points" example:"pump"`
	JitoOrderID  interface{}     `json:"jito_order_id"`
	// 预计获得积分
	EstimatedPoints decimal.Decimal `json:"estimated_points"` // 预计获得积分
}

type Quote struct {
	InputMint            string          `json:"inputMint"`
	InAmount             decimal.Decimal `json:"inAmount"`
	InDecimals           uint8           `json:"inDecimals"`
	OutDecimals          uint8           `json:"outDecimals"`
	OutputMint           string          `json:"outputMint"`
	OutAmount            decimal.Decimal `json:"outAmount"`
	OtherAmountThreshold string          `json:"otherAmountThreshold"`
	SwapMode             string          `json:"swapMode"`
	SlippageBps          string          `json:"slippageBps"`
	PlatformFee          uint64          `json:"platformFee"`
	PriceImpactPct       string          `json:"priceImpactPct"`
	RoutePlan            []RoutePlan     `json:"routePlan"`
	TimeTaken            string          `json:"timeTaken"`
}

type RoutePlan struct {
	SwapInfo SwapInfo `json:"swapInfo"`
	Percent  int      `json:"percent"`
}

type SwapInfo struct {
	AmmKey     string          `json:"ammKey"`
	Label      string          `json:"label"`
	InputMint  string          `json:"inputMint"`
	OutputMint string          `json:"outputMint"`
	InAmount   decimal.Decimal `json:"inAmount"`
	OutAmount  decimal.Decimal `json:"outAmount"`
	FeeAmount  float64         `json:"feeAmount"`
	FeeMint    string          `json:"feeMint"`
}

type RawTx struct {
	SwapTransaction      string `json:"swapTransaction"`
	LastValidBlockHeight int    `json:"lastValidBlockHeight"`
	RecentBlockhash      string `json:"recentBlockhash"`
}
