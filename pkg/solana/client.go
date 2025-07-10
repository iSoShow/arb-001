package solana

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

// Client 是Solana区块链的客户端
type Client struct {
	rpcClient *rpc.Client
	wsClient  *ws.Client
}

// NewClient 创建一个新的Solana客户端
func NewClient(rpcURL, wsURL string) (*Client, error) {
	if rpcURL == "" {
		return nil, fmt.Errorf("RPC URL is required")
	}

	rpcClient := rpc.New(rpcURL)

	var wsClient *ws.Client
	var err error
	if wsURL != "" {
		wsClient, err = ws.Connect(context.Background(), wsURL)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
		}
	}

	return &Client{
		rpcClient: rpcClient,
		wsClient:  wsClient,
	}, nil
}

// GetBalance 获取账户余额
func (c *Client) GetBalance(pubkey solana.PublicKey) (uint64, error) {
	balance, err := c.rpcClient.GetBalance(
		context.Background(),
		pubkey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance.Value, nil
}

// GetTransaction 获取交易详情
func (c *Client) GetTransaction(signature string) (*solana.Transaction, error) {
	tx, err := c.rpcClient.GetTransaction(
		context.Background(),
		signature,
		&rpc.GetTransactionOpts{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return tx.Transaction, nil
}

// SendTransaction 发送交易
func (c *Client) SendTransaction(tx *solana.Transaction) (string, error) {
	sig, err := c.rpcClient.SendTransaction(
		context.Background(),
		tx,
	)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	return sig, nil
}

// SubscribeTransactions 订阅新交易
func (c *Client) SubscribeTransactions(ctx context.Context, callback func(*solana.Transaction)) (*ws.Subscription, error) {
	if c.wsClient == nil {
		return nil, fmt.Errorf("WebSocket client is not initialized")
	}

	// 这里是示例代码，实际实现需要根据Solana的WebSocket API
	// 实现交易订阅逻辑
	return nil, fmt.Errorf("not implemented")
}

// PrivateKey 表示Solana私钥
type PrivateKey struct {
	key solana.PrivateKey
}

// LoadPrivateKeyFromFile 从文件加载私钥
func LoadPrivateKeyFromFile(path string) (*PrivateKey, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	var privateKeyBytes []byte
	if err := json.Unmarshal(data, &privateKeyBytes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
	}

	privateKey := solana.PrivateKey(privateKeyBytes)

	return &PrivateKey{
		key: privateKey,
	}, nil
}

// PublicKey 获取对应的公钥
func (p *PrivateKey) PublicKey() solana.PublicKey {
	return p.key.PublicKey()
}