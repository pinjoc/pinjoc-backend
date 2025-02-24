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
		Month             int32  `json:"month" validate:"required"`
		Year              int32  `json:"year" validate:"required"`
	}

	UpdatePayload struct {
		ID             int32 `json:"id" validate:"required"`
		AvailableToken int32 `json:"available_token" validate:"required"`
	}
)

func (p *OrderBookPayload) Validate() error {
	return Validate.Struct(p)
}

func (p *UpdatePayload) Validate() error {
	return Validate.Struct(p)
}
