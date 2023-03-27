package errors

import (
	"meta-aggregator/internal/pkg/constant"

	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

func NewDomainErrTokensAreIdentical() *errors.DomainError {
	return errors.NewDomainError(
		constant.DomainErrCodeTokensAreIdentical,
		constant.DomainErrMsgTokensAreIdentical,
		[]string{"tokenIn", "tokenOut"},
		nil,
	)
}
