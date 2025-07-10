package models

import (
	"time"

	"github.com/gagliardetto/solana-go"
)

// TransactionType 表示交易类型
type TransactionType string

const (
	// TransactionTypeSwap 表示交换交易
	TransactionTypeSwap TransactionType = "swap"
	// TransactionTypeTransfer 表示转账交易
	TransactionTypeTransfer TransactionType = "transfer"
	// TransactionTypeLiquidity 表示流动性操作
	TransactionTypeLiquidity TransactionType = "liquidity"
	// TransactionTypeUnknown 表示未知类型
	TransactionTypeUnknown TransactionType = "unknown"
)

// Transaction 表示一个Solana交易
type Transaction struct {
	Signature    string          `json:"signature"`
	BlockTime    time.Time       `json:"block_time"`
	Type         TransactionType `json:"type"`
	Amount       float64         `json:"amount"`
	Fee          float64         `json:"fee"`
	Program      string          `json:"program"`
	FromAddress  string          `json:"from_address"`
	ToAddress    string          `json:"to_address"`
	TokenAddress string          `json:"token_address,omitempty"`
	RawData      []byte          `json:"raw_data,omitempty"`
}

// ArbitrageOpportunity 表示套利机会
type ArbitrageOpportunity struct {
	ID              string    `json:"id"`
	DetectedAt      time.Time `json:"detected_at"`
	SourceTx        string    `json:"source_tx"`
	ExpectedProfit  float64   `json:"expected_profit"`
	ExecutionStatus string    `json:"execution_status"`
	ExecutedTx      string    `json:"executed_tx,omitempty"`
	ExecutedAt      time.Time `json:"executed_at,omitempty"`
	ActualProfit    float64   `json:"actual_profit,omitempty"`
}