package app

import (
	"context"
	"fishing-store/internal/app/config"
	"fishing-store/internal/entity"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nanmu42/gzip"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strconv"
)

func Run(ctx context.Context, cfg *config.Config) {
	closer := newCloser()
	logger := newLogger()
	router := chi.NewRouter()

	server := newServer(cfg.Server, router)
	closer.Add(server.Shutdown)

	go func() {
		logger.DPanic("ListenAndServe", zap.Any("Error", server.ListenAndServe()))
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		logger.Error("Close err", zap.Error(err))
	}
}

func newServer(cfg config.Server, router http.Handler) *http.Server {
	return &http.Server{
		Handler:        gzip.DefaultHandler().WrapHandler(router),
		Addr:           ":" + strconv.Itoa(cfg.Port),
		WriteTimeout:   cfg.WriteTimeout,
		ReadTimeout:    cfg.ReadTimeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}
}

//func newDataBase(cfg config.DataBase) *sqlx.DB {
//
//	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
//
//	db, err := sqlx.ConnectContext(ctx,
//		"mysql",
//		cfg.Username+":"+cfg.Password+"@"+cfg.Address+"/"+cfg.DBName+cfg.Params)
//	if err != nil {
//		panic(err)
//	}
//
//	db.SetConnMaxLifetime(cfg.MaxConnLifetime)
//	db.SetConnMaxIdleTime(cfg.MaxConnIdleTime)
//	db.SetMaxOpenConns(cfg.MaxOpenCons)
//	db.SetMaxIdleConns(cfg.MaxIdleCons)
//
//	return db
//}

func newLogger() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func newCloser() *entity.Closer {
	return &entity.Closer{}
}
