### endpoint
dazzling-enjoyment-production.up.railway.app/api/v1/clob/available-token

```
no-parameter
```

```
[
  {
    "DebtTokenName": "Aave Token",
    "DebtTokenSymbol": "AAVE",
    "DebtTokenAddress": "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9",
    "CollateralTokenName": "USDC",
    "CollateralTokenSymbol": "USDC",
    "CollateralAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
    "MaturityRange": "FEB 2026 ~ NOV 2026",
    "RateRange": "5.00% ~ 8.50%",
    "LendingVault": 0,
    "BorrowVault": 599158
  },
  {
    "DebtTokenName": "Wrapped ETH",
    "DebtTokenSymbol": "WETH",
    "DebtTokenAddress": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
    "CollateralTokenName": "USDC",
    "CollateralTokenSymbol": "USDC",
    "CollateralAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
    "MaturityRange": "AUG 2025 ~ AUG 2026",
    "RateRange": "4.00% ~ 7.50%",
    "LendingVault": 0,
    "BorrowVault": 495000
  },
]
```

dazzling-enjoyment-production.up.railway.app/api/v1/clob/clob

```
{
  "collateral_address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
  "debt_token_address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
  "month":8,
  "year":2025
	
}
```
```
[
  {
    "Rate": 4.50,
    "AvailableToken": 55000,
    "OrderType": "LEND"
  },
  {
    "Rate": 5.00,
    "AvailableToken": 65000,
    "OrderType": "LEND"
  },
  {
    "Rate": 5.50,
    "AvailableToken": 75000,
    "OrderType": "LEND"
  }
]
```

dazzling-enjoyment-production.up.railway.app/api/v1/clob/best-rate

```
{
  "collateral_address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
  "debt_token_address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
  "month":8,
  "year":2025
	
}
```

```
{
  "best_rate": "5.00"
}
```