package api

import (
	"math/big"
	"meta-aggregator/internal/pkg/utils/eth"
	"strconv"
	"strings"

	"meta-aggregator/internal/pkg/constant"
	aggregatorerrors "meta-aggregator/internal/pkg/errors"

	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

// Faucets
type (
	FindRouteRequest struct {
		TokenIn    string `form:"tokenIn" binding:"required"`
		TokenOut   string `form:"tokenOut" binding:"required"`
		AmountIn   string `form:"amountIn" binding:"required"`
		SaveGas    string `form:"saveGas"`
		Dexes      string `form:"dexes"`
		GasInclude string `form:"gasInclude"`
		GasPrice   string `form:"gasPrice"`
		Debug      string `form:"debug"`
	}

	FindEncodedRouteRequest struct {
		FindRouteRequest
		EncodedRequestParams
	}

	EncodedRequestParams struct {
		SlippageTolerance string `form:"slippageTolerance"`
		ChargeFeeBy       string `form:"chargeFeeBy"`
		FeeReceiver       string `form:"feeReceiver"`
		IsInBps           string `form:"isInBps"`
		FeeAmount         string `form:"feeAmount"`
		Deadline          string `form:"deadline"`
		To                string `form:"to"`
		ClientData        string `form:"clientData"`
		Referral          string `form:"referral"`
	}
)

func (r *FindEncodedRouteRequest) Validate() *errors.DomainError {
	if err := r.ValidateAmountIn(); err != nil {
		return err
	}

	if err := r.ValidateTokens(); err != nil {
		return err
	}

	if err := r.ValidateTo(); err != nil {
		return err
	}

	if err := r.ValidateFeeReceiver(); err != nil {
		return err
	}

	if err := r.ValidateFeeAmount(); err != nil {
		return err
	}

	if err := r.ValidateChargeFeeBy(); err != nil {
		return err
	}

	if err := r.ValidateSlippageTolerance(); err != nil {
		return err
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateAmountIn() *errors.DomainError {
	amountInBi, ok := new(big.Int).SetString(r.AmountIn, 10)
	if !ok || amountInBi.Cmp(constant.Zero) <= 0 {
		return errors.NewDomainErrorInvalid(nil, "amountIn")
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateTokens() *errors.DomainError {
	if strings.ToLower(r.TokenIn) == strings.ToLower(r.TokenOut) {
		return aggregatorerrors.NewDomainErrTokensAreIdentical()
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateTo() *errors.DomainError {
	if r.To == "" {
		return errors.NewDomainErrorRequired(nil, "to")
	}

	if !eth.ValidateAddress(r.To) {
		return errors.NewDomainErrorInvalid(nil, "to")
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateFeeReceiver() *errors.DomainError {
	if r.FeeReceiver == "" {
		return nil
	}

	if !eth.ValidateAddress(r.FeeReceiver) {
		return errors.NewDomainErrorInvalid(nil, "feeReceiver")
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateFeeAmount() *errors.DomainError {
	if r.FeeAmount == "" {
		return nil
	}

	_, ok := new(big.Int).SetString(r.FeeAmount, 10)
	if !ok {
		return errors.NewDomainErrorInvalid(nil, "feeAmount")
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateChargeFeeBy() *errors.DomainError {
	if r.FeeAmount == "" {
		return nil
	}

	if r.ChargeFeeBy != constant.ChargeFeeByCurrencyIn && r.ChargeFeeBy != constant.ChargeFeeByCurrencyOut {
		return errors.NewDomainErrorInvalid(nil, "chargeFeeBy")
	}

	return nil
}

func (r *FindEncodedRouteRequest) ValidateSlippageTolerance() *errors.DomainError {
	if r.SlippageTolerance == "" {
		return nil
	}

	slippageTolerance, err := strconv.ParseInt(r.SlippageTolerance, 10, 64)
	if err != nil {
		return nil
	}

	if slippageTolerance < 0 || slippageTolerance > constant.MaximumSlippage {
		return errors.NewDomainErrorOutOfRange(nil, "slippageTolerance")
	}

	return nil
}
