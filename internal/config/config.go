package config

import (
	"github.com/spf13/viper"
)

// Config 包含应用程序的所有配置
type Config struct {
	Solana   SolanaConfig   `mapstructure:"solana"`
	Jito     JitoConfig     `mapstructure:"jito"`
	Listener ListenerConfig `mapstructure:"listener"`
	Trader   TraderConfig   `mapstructure:"trader"`
}

// SolanaConfig 包含Solana连接配置
type SolanaConfig struct {
	RPCURL string `mapstructure:"rpc_url"`
	WSURL  string `mapstructure:"ws_url"`
}

// JitoConfig 包含Jito服务器配置
type JitoConfig struct {
	ServerURL string `mapstructure:"server_url"`
	UUID      string `mapstructure:"uuid"`
}

// ListenerConfig 包含交易监听配置
type ListenerConfig struct {
	MinTransactionAmount float64  `mapstructure:"min_transaction_amount"`
	TargetPrograms      []string `mapstructure:"target_programs"`
}

// TraderConfig 包含交易执行配置
type TraderConfig struct {
	PrivateKeyPath      string  `mapstructure:"private_key_path"`
	MaxTransactionAmount float64 `mapstructure:"max_transaction_amount"`
}

// Load 从指定路径加载配置文件
func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}