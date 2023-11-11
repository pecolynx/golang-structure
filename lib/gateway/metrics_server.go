package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/pecolynx/golang-structure/lib/domain"
	"github.com/pecolynx/golang-structure/lib/log"
)

const readHeaderTimeout = time.Duration(30) * time.Second

const (
	LibGatewayContextKey domain.ContextKey = "lib_gateway"
)

func MetricsServerProcess(ctx context.Context, port int, gracefulShutdownTimeSec int) error {
	router := gin.New()
	router.Use(gin.Recovery())

	httpServer := http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	router.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	logger := log.GetLoggerFromContext(ctx, LibGatewayContextKey)
	logger.InfoContext(ctx, fmt.Sprintf("metrics server listening at %v", httpServer.Addr))

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.InfoContext(ctx, fmt.Sprintf("failed to ListenAndServe. err: %v", err))
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		gracefulShutdownTime1 := time.Duration(gracefulShutdownTimeSec) * time.Second
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), gracefulShutdownTime1)
		defer shutdownCancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.InfoContext(ctx, fmt.Sprintf("Server forced to shutdown. err: %v", err))
			return err
		}
		return nil
	case err := <-errCh:
		return err
	}
}
