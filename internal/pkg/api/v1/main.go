package api

import (
	"meta-aggregator/internal/pkg/service"
	"meta-aggregator/pkg/context"

	t "github.com/KyberNetwork/kyberswap-error/pkg/transformers"
	"github.com/gin-gonic/gin"
)

type RouteApi struct {
	routeSvc service.IRouteSvc
}

func newRouteApi(routeSvc service.IRouteSvc) *RouteApi {
	return &RouteApi{
		routeSvc,
	}
}

func (a *RouteApi) setupRoute(rg *gin.RouterGroup) {
	rg.GET("/route/encode", handleFindRoute())
}

func (a *RouteApi) setupRestrictedRoute(rg *gin.RouterGroup) {
}

func handleFindRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.New(c).WithLogPrefix("route-api")

		request := FindEncodedRouteRequest{}
		if err := ctx.ShouldBindQuery(&request); err != nil {
			apiErr := t.RestTransformerInstance().ValidationErrToRestAPIErr(err)

			ctx.Errorf("%v", apiErr)
			ctx.AbortWithStatusJSON(apiErr.HttpStatus, apiErr)
			return
		}

		if err := request.Validate(); err != nil {
			apiErr := t.RestTransformerInstance().DomainErrToRestAPIErr(err)

			ctx.Errorf("%v", apiErr)
			ctx.AbortWithStatusJSON(apiErr.HttpStatus, apiErr)
			return
		}
	}
}
