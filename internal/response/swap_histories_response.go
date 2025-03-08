package response

// SwapHistoriesResponse 表示交易历史记录的响应结构。
type SwapHistoriesResponse struct {
	TransactionHistories []TransactionHistory `json:"transaction_histories"` // 交易历史记录列表
	HasMore              bool                 `json:"has_more"`              // 是否还有更多数据
}

// TransactionHistory 表示单笔交易历史记录。
type TransactionHistory struct {
	IsBuy     bool   `json:"is_buy"`     // 是否为买入
	Payer     string `json:"payer"`      // 支付方地址
	Signature string `json:"signature"`  // 交易签名
	BlockTime string `json:"block_time"` // 区块时间
	// Index        int    `json:"index"`         // 交易索引
	TokenAmount  string `json:"token_amount"`  // 代币数量
	NativeAmount string `json:"native_amount"` // 原生代币数量
}
