module github.com/vulcanize/eth-statediff-service

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.11
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
)

replace github.com/ethereum/go-ethereum v1.9.11 => github.com/vulcanize/go-ethereum v1.9.11-statediff-0.0.5