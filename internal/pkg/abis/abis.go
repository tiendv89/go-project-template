package abis

import (
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	Faucet abi.ABI
	Reward abi.ABI
)

func init() {
	builder := []struct {
		ABI  *abi.ABI
		data []byte
	}{
		{&Faucet, faucet},
		{&Reward, reward},
	}

	for _, b := range builder {
		var err error
		*b.ABI, err = abi.JSON(bytes.NewReader(b.data))
		if err != nil {
			panic(err)
		}
	}
}
