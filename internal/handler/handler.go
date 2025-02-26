package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pinjoc/pinjoc-backend/internal/service"
)

type Handler struct {
	CLOB interface {
		GetCLOB(ctx *fiber.Ctx) error
		GetAvailableToken(ctx *fiber.Ctx) error
		GetBestRate(ctx *fiber.Ctx) error
		UpdateAvailabeToken(ctx *fiber.Ctx) error
	}

	Tokenized interface {
		GetToken(ctx *fiber.Ctx) error
		GetAllToken(ctx *fiber.Ctx) error
		GetBestPrice(ctx *fiber.Ctx) error
		UpdateAmount(ctx *fiber.Ctx) error
	}
}

func NewHandler(db *pgxpool.Pool) Handler {
	return Handler{
		CLOB: &ClobHandler{
			s: service.NewService(db),
		},
		Tokenized: &TokenizedHandler{
			s: service.NewService(db),
		},
	}
}
