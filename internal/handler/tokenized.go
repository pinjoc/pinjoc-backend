package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/service"
)

type TokenizedHandler struct {
	s service.Service
}

func (h *TokenizedHandler) GetToken(ctx *fiber.Ctx) error {
	payload := new(model.TokenizedPayload)

	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := h.s.Tokenized.GetToken(ctx.Context(), *payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if res == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No available Tokenized",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *TokenizedHandler) GetAllToken(ctx *fiber.Ctx) error {
	res, err := h.s.Tokenized.GetAllToken(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *TokenizedHandler) GetBestPrice(ctx *fiber.Ctx) error {
	payload := new(model.TokenizedPayload)

	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := h.s.Tokenized.GetBestPrice(ctx.Context(), *payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if res == 0 {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"best_rate": "No available rate",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"best_price": fmt.Sprintf("%.2f", res),
	})
}

func (h *TokenizedHandler) UpdateAmount(ctx *fiber.Ctx) error {
	payload := new(model.UpdateAmount)

	if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := h.s.Tokenized.UpdateAmount(ctx.Context(), *payload)
	log.Println(res)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "order not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
