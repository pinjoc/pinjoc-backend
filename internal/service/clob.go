package service

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/storage/sqlc"
)

type ClobService struct {
	q *sqlc.Queries
}

func (s *ClobService) GetCLOB(ctx context.Context, payload model.OrderBookPayload) ([]sqlc.GetCLOBRow, error) {
	monthInt := MonthToInt[payload.Month]
	return s.q.GetCLOB(ctx, sqlc.GetCLOBParams{
		Address:   payload.CollateralAddress,
		Address_2: payload.DebtTokenAddress,
		Month:     monthInt,
		Year:      payload.Year,
	})
}

func (s *ClobService) GetBestRate(ctx context.Context, payload model.OrderBookPayload) (float64, error) {
	monthInt := MonthToInt[payload.Month]
	return s.q.GetBestRate(ctx, sqlc.GetBestRateParams{
		Address:   payload.CollateralAddress,
		Address_2: payload.DebtTokenAddress,
		Month:     monthInt,
		Year:      payload.Year,
	})
}

func (s *ClobService) GetAvailableToken(ctx context.Context) ([]sqlc.GetAvailableTokenRow, error) {
	return s.q.GetAvailableToken(ctx)
}

func (s *ClobService) updateAvailableToken(ctx context.Context, payload model.UpdatePayload) error {
	return s.q.UpdateTokenAvailable(ctx, sqlc.UpdateTokenAvailableParams{
		AvailableToken: payload.AvailableToken,
		ID:             payload.ID,
	})
}

func (s *ClobService) randomAmount() int32 {
	randAmount := rand.IntN(501) + 500
	updateType := rand.IntN(2)
	if updateType == 0 {
		return int32(randAmount)
	}

	return -int32(randAmount)
}

func (s *ClobService) RandomUpdate(ctx context.Context) error {
	id, err := s.q.GetRandomOrder(ctx)
	if err != nil {
		return err
	}

	amount := s.randomAmount()
	log.Println("CLOB Available update", id, amount)
	return s.updateAvailableToken(ctx, model.UpdatePayload{
		ID:             id,
		AvailableToken: amount,
	})
}

func (s *ClobService) UpdateAvailabe(ctx context.Context, payload model.UpdateAvailabe) (int32, error) {
	rateStr := fmt.Sprintf("%f", payload.Rate)
	rate := pgtype.Numeric{}
	if err := rate.Scan(rateStr); err != nil {
		return 0, err
	}
	return s.q.UpdateAvailable(ctx, sqlc.UpdateAvailableParams{
		AvailableToken: payload.AvailableToken,
		Address:        payload.CollateralAddress,
		Address_2:      payload.DebtTokenAddress,
		Month:          MonthToInt[payload.Month],
		Year:           payload.Year,
		OrderType:      payload.OrderType,
		Rate:           rate,
	})
}
