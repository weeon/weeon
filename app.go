package weeon

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/weeon/weeon/transport"
	"github.com/weeon/weeon/transport/http"
)

type AppInterface interface {
	Run() error
}

type Config struct {
	Service    string
	HTTPRouter func(r *gin.Engine)
}

type App struct {
	config  *Config
	servers []transport.Server
}

func New(ctx context.Context, cfg *Config) AppInterface {
	return &App{
		config: cfg,
	}
}

func (a *App) Run() error {
	a.setupTransport()

	// wait os signal
	a.waitSignal()
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	return nil
}

func (a *App) setupTransport() {
	a.servers = make([]transport.Server, 0)
	if a.config.HTTPRouter != nil {
		a.servers = append(a.servers, http.NewGinServer(a.config.HTTPRouter))
	}

	for _, s := range a.servers {
		go func(s transport.Server) {
			if err := s.Run(); err != nil {
				panic(err)
			}
		}(s)
	}
}

func (a *App) waitSignal() {
	// wait os signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	err := a.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
