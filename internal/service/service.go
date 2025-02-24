package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/storage/sqlc"
)

type Service struct {
	CLOB interface {
		GetCLOB(ctx context.Context, payload model.OrderBookPayload) ([]sqlc.GetCLOBRow, error)
		GetBestRate(ctx context.Context, payload model.OrderBookPayload) (float64, error)
		GetAvailableToken(ctx context.Context) ([]sqlc.GetAvailableTokenRow, error)
		RandomUpdate(ctx context.Context) error
	}
}

func NewService(db *pgxpool.Pool) Service {
	return Service{
		CLOB: &ClobService{
			q: sqlc.New(db),
		},
	}
}
