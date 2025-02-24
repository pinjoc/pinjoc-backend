package service

import (
	"context"
	"log"
	"math/rand/v2"

	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/storage/sqlc"
)

var MonthToStr = map[int32]string{
	1:  "JAN",
	2:  "FEB",
	3:  "MAR",
	4:  "APR",
	5:  "MAY",
	6:  "JUN",
	7:  "JUL",
	8:  "AUG",
	9:  "SEP",
	10: "OCT",
	11: "NOV",
	12: "DEC",
}

type ClobService struct {
	q *sqlc.Queries
}

func (s *ClobService) GetCLOB(ctx context.Context, payload model.OrderBookPayload) ([]sqlc.GetCLOBRow, error) {
	return s.q.GetCLOB(ctx, sqlc.GetCLOBParams{
		Address:   payload.CollateralAddress,
		Address_2: payload.DebtTokenAddress,
		Month:     payload.Month,
		Year:      payload.Year,
	})
}

func (s *ClobService) GetBestRate(ctx context.Context, payload model.OrderBookPayload) (float64, error) {
	return s.q.GetBestRate(ctx, sqlc.GetBestRateParams{
		Address:   payload.CollateralAddress,
		Address_2: payload.DebtTokenAddress,
		Month:     payload.Month,
		Year:      payload.Year,
	})
}

func (s *ClobService) GetAvailableToken(ctx context.Context) ([]sqlc.GetAvailableTokenRow, error) {
	return s.q.GetAvailableToken(ctx)
}

func (s *ClobService) UpdateAvailable(ctx context.Context, payload model.UpdatePayload) error {
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
	log.Println("Random update", id, amount)
	return s.UpdateAvailable(ctx, model.UpdatePayload{
		ID:             id,
		AvailableToken: amount,
	})
}
