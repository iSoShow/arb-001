# Solana 套利监听机器人

这是一个用于监听Solana链上大额交易并通过Jito服务器发送跟单交易的套利机器人。

## 功能特点

- 实时监听Solana链上的大额交易
- 自动分析交易数据，识别套利机会
- 通过Jito服务器发送跟单交易，减少MEV损失
- 高效的交易执行策略

## 安装要求

- Go 1.20 或更高版本
- Solana CLI

## 配置说明

在 `config.yaml` 文件中配置以下参数：

```yaml
solana:
  rpc_url: "https://api.mainnet-beta.solana.com"
  ws_url: "wss://api.mainnet-beta.solana.com"

jito:
  server_url: "https://jito-server.example.com"
  uuid: "your-jito-uuid"

listener:
  min_transaction_amount: 1000  # SOL
  target_programs: [
    "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin",  # Serum DEX
    "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4"   # Jupiter
  ]

trader:
  private_key_path: "./keypair.json"
  max_transaction_amount: 100  # SOL