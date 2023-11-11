package log

import (
	"context"
	"log/slog"

	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liblog "github.com/pecolynx/golang-structure/lib/log"
)

const (
	LibGatewayLoggerContextKey    libdomain.ContextKey = "lib_gateway"
	AppGORMLoggerContextKey       libdomain.ContextKey = "app_gorm"
	AppServiceLoggerContextKey    libdomain.ContextKey = "app_service"
	AppGatewayLoggerContextKey    libdomain.ContextKey = "app_gateway"
	AppControllerLoggerContextKey libdomain.ContextKey = "app_controller"
	AppTraceLoggerContextKey      libdomain.ContextKey = "app_trace"
)

var (
	LoggerKeys = []libdomain.ContextKey{
		LibGatewayLoggerContextKey,
		AppGORMLoggerContextKey,
		AppServiceLoggerContextKey,
		AppGatewayLoggerContextKey,
		AppControllerLoggerContextKey,
		AppTraceLoggerContextKey,
	}
	// Loggers     map[libdomain.ContextKey]*slog.Logger = make(map[libdomain.ContextKey]*slog.Logger)
	// LogHandlers map[slog.Level]slog.Handler           = make(map[slog.Level]slog.Handler)
	// lock        sync.Mutex
)

// func init() {
// 	for _, level := range []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError} {
// 		LogHandlers[level] = &LogHandler{Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
// 			Level: level,
// 		})}
// 	}

// 	for _, key := range LoggerKeys {
// 		Loggers[key] = slog.New(LogHandlers[slog.LevelWarn])
// 	}
// }

func InitLogger(ctx context.Context) context.Context {
	for _, key := range LoggerKeys {
		if _, ok := liblog.Loggers[key]; !ok {
			liblog.Loggers[key] = slog.New(liblog.LogHandlers[liblog.DefaultLogLevel])
		}
		ctx = context.WithValue(ctx, key, liblog.Loggers[key])
	}
	return ctx
}

// func SetLogLevel(logLevel slog.Level) {
// 	for _, key := range LoggerKeys {
// 		if _, ok := Loggers[key]; ok {
// 			setLogLevel(key, logLevel)
// 		}
// 	}
// }

// func setLogLevel(contextKey libdomain.ContextKey, logLevel slog.Level) {
// 	Loggers[contextKey] = slog.New(LogHandlers[logLevel])
// }

// // GetLoggerFromContext Gets the logger from context
// func GetLoggerFromContext(ctx context.Context, key libdomain.ContextKey) *slog.Logger {
// 	if ctx == nil {
// 		panic("nil context")
// 	}

// 	logger, ok := ctx.Value(key).(*slog.Logger)
// 	if ok {
// 		return logger
// 	}

// 	lock.Lock()
// 	defer lock.Unlock()

// 	if _, ok := Loggers[key]; !ok {
// 		Loggers[key] = slog.New(&LogHandler{Handler: slog.NewJSONHandler(os.Stdout, nil)})
// 		Loggers[key].WarnContext(ctx, fmt.Sprintf("logger not found. logger: %s", key))
// 	}

// 	return Loggers[key]
// }
