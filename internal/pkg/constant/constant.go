package constant

import (
	"math/big"
)

var BONE = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
var BoneFloat, _ = new(big.Float).SetString("1000000000000000000")

var One *big.Int
var Zero *big.Int
var Two *big.Int
var Three *big.Int
var Four *big.Int
var Five *big.Int
var Ten *big.Float
var BipBase *big.Int
var tenPowDecimals []*big.Float
var tenPowInt []*big.Int
var LowReserve *big.Float

// MaximumPriceUsd Assume that no token is going to have price > this value
const MaximumPriceUsd = 200000

const MinLiquidityUsd = 200
const MajorThresholdUsd = 100
const maxDecimals = 60

// ForceUpdateThreshold is the timeframe (in second) that we force something to update
const ForceUpdateThreshold = 3600

type AmountRange int64

const (
	OutOfRange AmountRange = iota
	TenPercentage
	TwentyPercentage
	ThirtyPercentage
	FortyPercentage
	FiftyPercentage
	SixtyPercentage
	SeventyPercentage
	EightyPercentage
	NinetyPercentage
	HundredPercentage
)

const MinAmount = "MinAmount"
const MaxAmount = "MaxAmount"

//const TenPercentage = "Ten"
//const TwentyPercentage = "Twenty"
//const ThirtyPercentage = "Thirty"
//const FortyPercentage = "Forty"
//const FiftyPercentage = "Fifty"
//const SixtyPercentage = "Sixty"
//const SeventyPercentage = "Seventy"
//const EightyPercentage = "Eighty"
//const NinetyPercentage = "Ninety"
//const OutOfRange = ""

func TenPowDecimals(decimal uint8) *big.Float {
	if decimal < maxDecimals {
		return tenPowDecimals[decimal]
	}
	return new(big.Float).Quo(Ten, new(big.Float).SetInt64(int64(decimal)))
}

func TenPowInt(decimal uint8) *big.Int {
	if decimal < maxDecimals {
		return tenPowInt[decimal]
	}
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
}

func init() {
	Zero = big.NewInt(0)
	One = big.NewInt(1)
	Two = big.NewInt(2)
	Three = big.NewInt(3)
	Four = big.NewInt(4)
	Five = big.NewInt(5)
	Ten = new(big.Float).SetFloat64(10)
	BipBase = big.NewInt(10000)
	tenPowDecimals = make([]*big.Float, maxDecimals)
	tenPowDecimals[0] = new(big.Float).SetFloat64(1)
	tenPowInt = make([]*big.Int, maxDecimals)
	tenPowInt[0] = big.NewInt(1)
	for i := 1; i < maxDecimals; i++ {
		tenPowDecimals[i] = new(big.Float).Mul(tenPowDecimals[i-1], Ten)
		tenPowInt[i] = new(big.Int).Mul(tenPowInt[i-1], big.NewInt(10))
	}
	LowReserve = new(big.Float).SetFloat64(1e-6)
}
