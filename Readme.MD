# Pokemon API - Go & Redis

A lightweight Pokemon API built with Go that retrieves Pokemon data cached in Redis by type (Water, Electric, Grass, Legendary, Fire).

## 📋 Prerequisites

- **Go** (version 1.25.1 or later)
- **Redis** (running locally on port 6379)

## 🚀 Quick Start

### 1. Install Dependencies

```bash
go mod download
```

### 2. Setup Redis

Make sure Redis is running on `localhost:6379`:

```bash
# If Redis is installed via Homebrew
brew services start redis

# Or run Redis directly
redis-server
```

### 3. Load Pokemon Data

Load the Pokemon data into Redis using the provided data file:

```bash
cat pokemon.redis | redis-cli
```

### 4. Run the API

```bash
go run main.go
```

The API will start on `http://localhost:3000`

## API Endpoints

Get Pokemon by type:

- **Water Pokemon**: `http://localhost:3000/water`
- **Electric Pokemon**: `http://localhost:3000/electric`
- **Grass Pokemon**: `http://localhost:3000/grass`
- **Fire Pokemon**: `http://localhost:3000/fire`
- **Legendary Pokemon**: `http://localhost:3000/legendary`

Each endpoint returns a JSON array of Pokemon with their details (name, type, XP, power, level).

## How It Works

1. The app connects to Redis at startup
2. Fetches Pokemon data by filtering keys with the requested type
3. Returns the data as JSON

## Dependencies

- `github.com/go-redis/redis/v8` - Redis client for Go
