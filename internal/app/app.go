package app

import (
	"EffectiveMobile/config"
	v1 "EffectiveMobile/internal/controller/http/v1"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	log *zap.Logger
	c   config.Config
}

func NewApp(log *zap.Logger, config config.Config) (*App, error) {
	return &App{
		log: log,
		c:   config,
	}, nil
}

func (a *App) Run() {

	r := gin.Default()
	v1.NewRouter(r, a.log)

	// Graceful shutdown mechanism
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.c.HTTP.Host, a.c.HTTP.Port),
		Handler: r,
	}

	go func() {
		a.log.Info(fmt.Sprintf("Starting on port %s...", a.c.HTTP.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Fatal("Could not listen", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.log.Warn("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.log.Fatal("Server forced to shutdown:", zap.Error(err))
	}
	// End of graceful shutdown mechanism

}
