package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pinjoc/pinjoc-backend/internal/model"
	"github.com/pinjoc/pinjoc-backend/internal/service"
)

type ClobHandler struct {
	s service.Service
}

func (h *ClobHandler) GetCLOB(ctx *fiber.Ctx) error {
	payload := new(model.OrderBookPayload)

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

	res, err := h.s.CLOB.GetCLOB(ctx.Context(), *payload)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if res == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No available CLOB",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *ClobHandler) GetAvailableToken(ctx *fiber.Ctx) error {
	res, err := h.s.CLOB.GetAvailableToken(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (h *ClobHandler) GetBestRate(ctx *fiber.Ctx) error {
	payload := new(model.OrderBookPayload)

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

	res, err := h.s.CLOB.GetBestRate(ctx.Context(), *payload)
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
		"best_rate": fmt.Sprintf("%.2f", res),
	})
}
