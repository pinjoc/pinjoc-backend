package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/storage/sqlc"
)

var MonthToInt = map[string]int32{
	"JAN": 1,
	"FEB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

type Service struct {
	CLOB interface {
		GetCLOB(ctx context.Context, payload model.OrderBookPayload) ([]sqlc.GetCLOBRow, error)
		GetBestRate(ctx context.Context, payload model.OrderBookPayload) (float64, error)
		GetAvailableToken(ctx context.Context) ([]sqlc.GetAvailableTokenRow, error)
		RandomUpdate(ctx context.Context) error
	}

	Tokenized interface {
		GetAllToken(ctx context.Context) ([]sqlc.GetAllTokenRow, error)
		GetToken(ctx context.Context, payload model.TokenizedPayload) ([]sqlc.GetTokenRow, error)
		GetBestPrice(ctx context.Context, payload model.TokenizedPayload) (float64, error)
		RandomUpdate(ctx context.Context) error
		RandomVolume(ctx context.Context) error
	}
}

func NewService(db *pgxpool.Pool) Service {
	return Service{
		CLOB: &ClobService{
			q: sqlc.New(db),
		},
		Tokenized: &TokenizedService{
			q: sqlc.New(db),
		},
	}
}
