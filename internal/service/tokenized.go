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

func (s *TokenizedService) UpdateAmount(ctx context.Context, payload model.UpdatePayload) error {
	return s.q.TokenAmount(ctx, sqlc.TokenAmountParams{
		Amount: payload.AvailableToken,
		ID:     payload.ID,
	})
}

func (s *TokenizedService) UpdateVolume(ctx context.Context, payload model.UpdatePayload) error {
	return s.q.TokenVolume(ctx, sqlc.TokenVolumeParams{
		Volume: payload.AvailableToken,
		ID:     payload.ID,
	})
}

func (s *TokenizedService) randomAmount() int32 {
	randAmount := rand.IntN(51) + 50
	updateType := rand.IntN(2)
	if updateType == 0 {
		return int32(randAmount)
	}

	return -int32(randAmount)
}

func (s *TokenizedService) randomVolume() int32 {
	randAmount := rand.IntN(501) + 2000
	updateType := rand.IntN(2)
	if updateType == 0 {
		return int32(randAmount)
	}

	return -int32(randAmount)
}

func (s *TokenizedService) RandomUpdate(ctx context.Context) error {
	id, err := s.q.GetRandomToken(ctx)
	if err != nil {
		return err
	}

	amount := s.randomAmount()
	log.Println("Token Amount update", id, amount)
	return s.UpdateAmount(ctx, model.UpdatePayload{
		ID:             id,
		AvailableToken: amount,
	})
}

func (s *TokenizedService) RandomVolume(ctx context.Context) error {
	id, err := s.q.GetRandomToken(ctx)
	if err != nil {
		return err
	}

	volume := s.randomVolume()
	log.Println("Token volume update", id, volume)
	return s.UpdateVolume(ctx, model.UpdatePayload{
		ID:             id,
		AvailableToken: volume,
	})
}
