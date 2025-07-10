package listener

import (
	"context"
	"fmt"
	"sync"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/solana-arb-bot/internal/config"
	"github.com/yourusername/solana-arb-bot/pkg/solana"
)

// Transaction 表示监听到的交易
type Transaction struct {
	Signature string
	Amount    float64
	Program   string
	Data      []byte
	Accounts  []string
}

// TransactionHandler 是处理监听到的交易的回调函数
type TransactionHandler func(tx *Transaction)

// Listener 监听Solana链上的交易
type Listener struct {
	config    config.ListenerConfig
	client    *solana.Client
	handler   TransactionHandler
	mu        sync.Mutex
	subscribed bool
}

// NewListener 创建一个新的交易监听器
func NewListener(config config.ListenerConfig, client *solana.Client) (*Listener, error) {
	if client == nil {
		return nil, fmt.Errorf("solana client is required")
	}

	return &Listener{
		config:  config,
		client:  client,
		handler: nil,
	}, nil
}

// SetTransactionHandler 设置交易处理回调
func (l *Listener) SetTransactionHandler(handler TransactionHandler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handler = handler
}

// Start 开始监听交易
func (l *Listener) Start(ctx context.Context) error {
	l.mu.Lock()
	if l.subscribed {
		l.mu.Unlock()
		return fmt.Errorf("listener already started")
	}
	l.subscribed = true
	l.mu.Unlock()

	logrus.Info("Starting transaction listener")

	// 设置交易监听过滤器
	// 这里需要根据Solana的API实现具体的监听逻辑
	// 以下是示例代码

	// 监听新确认的交易
	subscription, err := l.client.SubscribeTransactions(ctx, func(tx *solana.Transaction) {
		// 处理交易
		l.processTransaction(tx)
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe to transactions: %w", err)
	}

	// 等待上下文取消
	<-ctx.Done()

	// 取消订阅
	if err := subscription.Unsubscribe(); err != nil {
		return fmt.Errorf("failed to unsubscribe: %w", err)
	}

	l.mu.Lock()
	l.subscribed = false
	l.mu.Unlock()

	return nil
}

// processTransaction 处理监听到的交易
func (l *Listener) processTransaction(tx *solana.Transaction) {
	// 检查交易是否满足我们的条件
	// 1. 检查交易金额是否超过最小阈值
	// 2. 检查交易是否涉及目标程序

	// 这里是示例代码，实际实现需要根据Solana的交易结构进行解析
	amount := 0.0 // 从交易中提取金额
	program := "" // 从交易中提取程序ID

	// 检查金额是否满足条件
	if amount < l.config.MinTransactionAmount {
		return
	}

	// 检查程序是否是目标程序
	targetFound := false
	for _, targetProgram := range l.config.TargetPrograms {
		if program == targetProgram {
			targetFound = true
			break
		}
	}

	if !targetFound {
		return
	}

	// 创建交易对象
	transaction := &Transaction{
		Signature: tx.Signatures[0].String(),
		Amount:    amount,
		Program:   program,
		// 其他字段根据需要填充
	}

	// 调用处理回调
	l.mu.Lock()
	handler := l.handler
	l.mu.Unlock()

	if handler != nil {
		handler(transaction)
	}
}