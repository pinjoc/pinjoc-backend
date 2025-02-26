package model

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type (
	OrderBookPayload struct {
		CollateralAddress string `json:"collateral_address" validate:"required"`
		DebtTokenAddress  string `json:"debt_token_address" validate:"required"`
		Month             string `json:"month" validate:"required"`
		Year              int32  `json:"year" validate:"required"`
	}

	TokenizedPayload struct {
		QouteToken string  `json:"qoute_token" validate:"required"`
		BaseToken  string  `json:"base_token" validate:"required"`
		Month      string  `json:"month" validate:"required"`
		Year       int32   `json:"year" validate:"required"`
		Rate       float64 `json:"rate" validate:"required"`
	}

	UpdatePayload struct {
		ID             int32 `json:"id" validate:"required"`
		AvailableToken int32 `json:"available_token" validate:"required"`
	}
)

func (p *OrderBookPayload) Validate() error {
	return Validate.Struct(p)
}

func (p *TokenizedPayload) Validate() error {
	return Validate.Struct(p)
}

func (p *UpdatePayload) Validate() error {
	return Validate.Struct(p)
}

type (
	UpdateAvailabe struct {
		AvailableToken    int32   `json:"available_token" validate:"required"`
		CollateralAddress string  `json:"collateral_address" validate:"required"`
		DebtTokenAddress  string  `json:"debt_token_address" validate:"required"`
		Month             string  `json:"month" validate:"required"`
		Year              int32   `json:"year" validate:"required"`
		OrderType         string  `json:"order_type" validate:"required"`
		Rate              float64 `json:"rate" validate:"required"`
	}

	UpdateAmount struct {
		Amount     int32   `json:"amount" validate:"required"`
		QouteToken string  `json:"qoute_token" validate:"required"`
		BaseToken  string  `json:"base_token" validate:"required"`
		Month      string  `json:"month" validate:"required"`
		Year       int32   `json:"year" validate:"required"`
		OrderType  string  `json:"order_type" validate:"required"`
		Rate       float64 `json:"rate" validate:"required"`
	}
)

func (p *UpdateAvailabe) Validate() error {
	return Validate.Struct(p)
}

func (p *UpdateAmount) Validate() error {
	return Validate.Struct(p)
}
