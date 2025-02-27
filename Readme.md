# CLOB Money Legos - Dynamic Lending Protocol
We are trying to build CLOB money legos. Specifically dynamic lending protocol based on supply and demand. This yield rate will be the one that will be traded. Can you help me to change our topics into CLOB money legos

## Overview
A backend service built with Go and PostgreSQL that provides dummy data simulation through APIs to demonstrate a Central Limit Order Book (CLOB) for lending/borrowing operations. This service is designed to be consumed by frontend applications to simulate supply-demand based lending protocols.

## Purpose
- Provide dummy data simulation for CLOB testing
- Demonstrate dynamic yield rate determination based on supply and demand
- Serve as a mock backend for frontend development and testing
- Simulate order book management for lending/borrowing operations

## Tech Stack
- **Backend**: Go 1.24
- **Database**: PostgreSQL 17
- **ORM**: SQLC

## API Documentation

### CLOB Endpoints (`/api/v1/clob`)

#### 1. Get Available Tokens
Lists available token pairs with their lending/borrowing states.

```http
GET /available-token
```

**Response:**
```json
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
  }
]
```

#### 2. Get CLOB State
Returns current order book state for specific token pair.

```http
POST /clob
```

**Request:**
```json
{
  "collateral_address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
  "debt_token_address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
  "month": 8,
  "year": 2025
}
```

**Response:**
```json
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

#### 3. Get Best Rate
Retrieves best available rate for token pair.

```http
POST /best-rate
```

**Request:**
```json
{
  "collateral_address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
  "debt_token_address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
  "month": 8,
  "year": 2025
}
```

**Response:**
```json
{
  "best_rate": "5.00"
}
```

### Update Available Token - PUT
```http
PUT /available-token
```

**Request:**
```json
{
  "available_token": 70000,
  "collateral_address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
  "debt_token_address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
  "month": "MAY",
  "year": 2025,
  "rate": 4.0,
  "order_type": "BORROW"
}
```

### GET Maturity and Best Rate
```http
POST /maturity-best-rate
```

**Response:**
```json
{
  "message": "success"
}
```

**Request:**
```json
{
  "collateral_address": "0xf14442CCE4511D0B5DC34425bceA50Ca67626c3a",
  "debt_token_address": "0x0F848482cC12EA259DA229e7c5C4949EdA7E6475"
}
```

**Response:**
```json
[
    {
        "Maturity": "MAY 2025",
        "BestRate": 6.5
    },
    {
        "Maturity": "AUG 2025",
        "BestRate": 11
    },
    {
        "Maturity": "NOV 2025",
        "BestRate": 19.5
    },
    {
        "Maturity": "FEB 2026",
        "BestRate": 28.1
    }
]
```
### Token Endpoints (`/api/v1/token`)

#### 1. Get Available Tokens
Lists available token pairs with their trading states.

```http
GET /available-token
```

**Response:**
```json
[
  {
    "QuoteTokenName": "Test Token dua",
    "QuoteTokenSymbol": "TT2",
    "QuoteTokenAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB42",
    "BaseTokenName": "Test Token",
    "BaseTokenSymbol": "TT",
    "BaseTokenAddress": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C591",
    "PriceRange": "5000 ~ 7000",
    "MaturityRange": "AUG 2025 ~ NOV 2025",
    "Volume24h": 200000
  }
]
```

#### 2. Get Token State
Returns current order book state for specific token pair.

```http
POST /token
```

**Request:**
```json
{
  "qoute_token": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C591",
  "base_token": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB42",
  "month": "MAY",
  "year": 2025,
  "rate": 2.5
}
```

**Response:**
```json
[
  {
    "Price": 1000,
    "Amount": 1000,
    "OrderType": "SELL"
  }
]
```

#### 3. Get Best Price
Retrieves best available price for token pair.

```http
POST /best-price
```

**Request:**
```json
{
  "qoute_token": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C591",
  "base_token": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB42",
  "month": "MAY",
  "year": 2025,
  "rate": 2.5
}
```

**Response:**
```json
{
  "best_price": "1000.00"
}
```

### Update Available Token - PUT
```http
PUT /available-token
```

**Request:**
```json
{
  "amount": 5000,
  "qoute_token": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C591",
  "base_token": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB42",
  "month": "MAY",
  "year": 2025,
  "rate": 2.5,
  "order_type": "SELL"
} 
```

**Response:**
```json
{
  "message": "success"
}
```

## Quick Start

1. Clone repository
```bash
git clone https://github.com/pinjoc/pinjoc-backend.git
cd pinjoc-backend
```

2. Create `.env` file
```env
DB_ADDR=postgres://username:password@localhost:5432/dbname?sslmode=disable
```

3. Run migrations
```bash
make migrate-up
```

4. Start server
```bash
go run cmd/main.go
```

## Development

### Database Migrations
```bash
# Create new migration
make migrations name=migration_name

# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down
```

### Update Dependencies
```bash
make go-tidy
```