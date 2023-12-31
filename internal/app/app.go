package app

import (
	"context"
	"fishing-store/internal/api/handler"
	"fishing-store/internal/api/httpMiddleware"
	"fishing-store/internal/app/config"
	"fishing-store/internal/entity"
	"fishing-store/internal/repository"
	"fishing-store/internal/service"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4/pgxpool"
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

	db := newDataBase(cfg.DataBase)
	err := db.Ping(ctx)
	if err != nil {
		panic(err.Error())
	}

	// Repository
	categoryRepository := repository.NewCategoryRepository(db)
	manufacturerRepository := repository.NewManufacturerRepository(db)
	productRepository := repository.NewProductRepository(db)
	supplyRepository := repository.NewSupplyRepository(db)
	saleRepository := repository.NewSaleRepository(db)

	// Service
	categoryService := service.NewCategoryService(categoryRepository, productRepository)
	manufacturerService := service.NewManufacturerService(manufacturerRepository, productRepository)
	productService := service.NewProductService(productRepository)
	supplyService := service.NewSupplyService(supplyRepository, productRepository)
	saleService := service.NewSaleService(saleRepository, productRepository)

	// API
	middleware := httpMiddleware.NewMiddleware(logger)
	handler.RegisterCategoryHandlers(router, categoryService, middleware)
	handler.RegisterManufacturerHandlers(router, manufacturerService, middleware)
	handler.RegisterProductHandlers(router, productService, middleware)
	handler.RegisterSupplyHandlers(router, supplyService, middleware)
	handler.RegisterSaleHandlers(router, saleService, middleware)

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

func newDataBase(cfg config.DataBase) *pgxpool.Pool {

	pgxConfig, err := pgxpool.ParseConfig(
		"postgres" + "://" +
			cfg.Username + ":" + cfg.Password + "@" + cfg.Address + "/" + cfg.DBName + cfg.Params)
	if err != nil {
		panic(err.Error())
	}
	//pgLogger := zapadapter.NewLogger(logger)
	//config.ConnConfig.Logger = pgLogger

	db, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		panic(err.Error())
	}

	return db
}

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
