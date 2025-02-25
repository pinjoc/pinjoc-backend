package api

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/pinjoc/pinjoc-backend/internal/env"
	"github.com/pinjoc/pinjoc-backend/internal/handler"
	"github.com/pinjoc/pinjoc-backend/internal/service"
	"github.com/pinjoc/pinjoc-backend/lib/config"
	"github.com/pinjoc/pinjoc-backend/lib/db/postgres"
	"github.com/robfig/cron/v3"
)

func InitRandom(svc service.Service) {
	c := cron.New(cron.WithSeconds())

	// Jalankan setiap 30 detik
	_, err := c.AddFunc("*/30 * * * * *", func() {
		log.Println("Running RandomUpdate...")
		svc.CLOB.RandomUpdate(context.Background())
		svc.Tokenized.RandomUpdate(context.Background())
		svc.Tokenized.RandomVolume(context.Background())
	})

	if err != nil {
		log.Fatalf("Failed to initialize cron job: %v", err)
	}

	c.Start() // Mulai cron job
}

func InitConfig() config.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	return config.Config{
		AddrHttp: ":8080",
		DbConfig: config.DBConfig{
			Addr:         env.GetString("DB_ADDR", ""),
			MaxOpenConns: 5,
			MaxIdleConns: 5,
			MaxIdleTime:  "10m",
		},
	}
}

func InitDatabase(dbCfg config.DBConfig) (*pgxpool.Pool, error) {
	conn, err := postgres.New(dbCfg.Addr, dbCfg.MaxOpenConns, dbCfg.MaxIdleConns, dbCfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Success Connected Database")
	return conn, nil
}

func InitServer() {
	cfg := InitConfig()

	db, err := InitDatabase(cfg.DbConfig)
	if err != nil {
		log.Fatalf("failed start server %v", err)
	}

	handler := handler.NewHandler(db)
	appConfig := AppConfig{
		h: handler,
		c: cfg,
	}

	app := NewApp(appConfig)
	app.RegisterRoute()

	svc := service.NewService(db)
	InitRandom(svc)

	if err := app.Run(); err != nil {
		log.Fatalf("failed start server %v", err)
	}
}
