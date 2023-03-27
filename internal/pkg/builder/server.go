package builder

import (
	"go-template/internal/pkg/config"
	"net/http"
	"os"

	sentrygin "github.com/getsentry/sentry-go/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newServer(config *config.HttpConfig) (*gin.Engine, error) {
	gin.SetMode(config.Mode)

	server := gin.Default()
	server.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// just enable datadog if $DD_ENABLED is not empty
	if os.Getenv("DD_ENABLED") != "" {
		server.Use(gintrace.Middleware(os.Getenv("DD_SERVICE")))
	}

	setCORS(server)
	server.GET("/ping", func(c *gin.Context) { c.AbortWithStatus(http.StatusOK) })

	return server, nil
}

func setCORS(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods(http.MethodOptions)
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("x-request-id")
	corsConfig.AddAllowHeaders("X-Request-Id")
	engine.Use(cors.New(corsConfig))
}
