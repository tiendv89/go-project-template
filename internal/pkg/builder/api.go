package builder

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/pkg/config"
)

type IRunner interface {
	Run() error
}

type apiBuilder struct {
	server *gin.Engine
}

func NewApiBuilder() (IRunner, error) {
	server, _ := newServer(&config.Instance().Http)
	return &apiBuilder{server: server}, nil
}

func (f *apiBuilder) Run() error {
	return f.server.Run(config.Instance().Http.BindAddress)
}
