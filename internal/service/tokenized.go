package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/storage/sqlc"
)

type TokenizedService struct {
	q *sqlc.Queries
}

func (s *TokenizedService) GetAllToken(ctx context.Context) ([]sqlc.GetAllTokenRow, error) {
	return s.q.GetAllToken(ctx)
}

func (s *TokenizedService) GetToken(ctx context.Context, payload model.TokenizedPayload) ([]sqlc.GetTokenRow, error) {
	monthInt := MonthToInt[payload.Month]
	rateStr := fmt.Sprintf("%f", payload.Rate)
	rate := pgtype.Numeric{}
	if err := rate.Scan(rateStr); err != nil {
		return nil, err
	}
	return s.q.GetToken(ctx, sqlc.GetTokenParams{
		Address:   payload.QouteToken,
		Address_2: payload.BaseToken,
		Month:     monthInt,
		Year:      payload.Year,
		Rate:      rate,
	})
}

func (s *TokenizedService) GetBestPrice(ctx context.Context, payload model.TokenizedPayload) (float64, error) {
	monthInt := MonthToInt[payload.Month]
	rateStr := fmt.Sprintf("%f", payload.Rate)
	rate := pgtype.Numeric{}
	if err := rate.Scan(rateStr); err != nil {
		return 0, err
	}
	return s.q.GetBasePrice(ctx, sqlc.GetBasePriceParams{
		Address:   payload.QouteToken,
		Address_2: payload.BaseToken,
		Month:     monthInt,
		Year:      payload.Year,
		Rate:      rate,
	})
}
