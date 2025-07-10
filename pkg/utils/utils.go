package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

// GenerateRandomID 生成随机ID
func GenerateRandomID(prefix string) string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
	}

	return fmt.Sprintf("%s-%s", prefix, hex.EncodeToString(b))
}

// SolToLamports 将SOL转换为lamports
func SolToLamports(sol float64) uint64 {
	lamports := sol * 1e9
	return uint64(lamports)
}

// LamportsToSol 将lamports转换为SOL
func LamportsToSol(lamports uint64) float64 {
	return float64(lamports) / 1e9
}

// CalculateTransactionFee 计算交易费用
func CalculateTransactionFee(numSignatures, numInstructions int) uint64 {
	// 这里是示例代码，实际费用计算需要根据Solana的费用模型
	baseFee := uint64(5000) // 基础费用，单位lamports
	signatureFee := uint64(numSignatures) * 1000
	instructionFee := uint64(numInstructions) * 2000

	return baseFee + signatureFee + instructionFee
}

// WaitWithTimeout 带超时的等待
func WaitWithTimeout(callback func() bool, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		if callback() {
			return true
		}

		time.Sleep(100 * time.Millisecond)
	}

	return false
}