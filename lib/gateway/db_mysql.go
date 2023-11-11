package gateway

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	slog_gorm "github.com/orandin/slog-gorm"
)

type gormLogger struct {
	logger *slog.Logger
}

func (l *gormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	l.logger.InfoContext(ctx, s, args...)
}

func (l *gormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.logger.WarnContext(ctx, s, args...)
}

func (l *gormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	l.logger.ErrorContext(ctx, s, args...)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	if err != nil {
		l.logger.ErrorContext(ctx, sql, slog.Duration("elapsed", elapsed))
		return
	}
	l.logger.DebugContext(ctx, sql, slog.Duration("elapsed", elapsed))
}

func OpenMySQL(username, password, host string, port int, database string, logger *slog.Logger) (*gorm.DB, error) {
	logger.Info("INFO")
	logger.Debug("DEBUG")
	logger.Warn("WARN")
	c := mysql.Config{
		DBName:               database,
		User:                 username,
		Passwd:               password,
		Addr:                 fmt.Sprintf("%s:%d", host, port),
		Net:                  "tcp",
		ParseTime:            true,
		MultiStatements:      true,
		Params:               map[string]string{"charset": "utf8"},
		Collation:            "utf8mb4_unicode_ci",
		AllowNativePasswords: true,
		Loc:                  time.UTC,
		// Loc:                  jst,
	}
	dsn := c.FormatDSN()
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&multiStatements=true", username, password, host, port, database)
	return gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithLogger(logger), // Optional, use slog.Default() by default
			slog_gorm.WithTraceAll(),     // trace all messages
			// slog_gorm.SetLogLevel(DefaultLogType, slog.Level(32)), // Define the default logging level
		),
		// Logger: gorm_logrus.New(),
		// Logger: &gormLogger{
		// 	logger: logger,
		// },
	})
}

func MigrateMySQLDB(db *gorm.DB, sqlFS embed.FS) error {
	driverName := "mysql"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return migrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
	})
}
