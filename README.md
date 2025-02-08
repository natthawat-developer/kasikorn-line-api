# Kasikorn Line API

## Overview
Kasikorn Line API is a microservice-based API built with Golang and Fiber. It provides account, banner, debit, transaction, and user management functionalities.

## Features
- User authentication and management
- Account and transaction handling
- Banner management
- Debit card details retrieval
- Security features like CORS, rate limiting, and helmet middleware
- Health check endpoint


## Project Structure
```
kasikorn-line-api
├─ .git
├─ .gitignore
├─ Dockerfile
├─ README.md
├─ cmd
│  └─ main.go
├─ config.yaml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ config
│  │  └─ config.go
│  ├─ user
│  │  ├─ handlers
│  │  │  └─ user_handler.go
│  │  ├─ models
│  │  │  └─ user.go
│  │  ├─ repositories
│  │  │  ├─ models
│  │  │  │  └─ user.go
│  │  │  └─ user_repository.go
│  │  ├─ routes
│  │  │  └─ routes.go
│  │  └─ services
│  │     └─ user_service.go
│  ├─ other modules...
└─ pkg
   ├─ database
   │  └─ database.go
   ├─ error
   │  └─ error.go
   ├─ log
   │  └─ logger.go
   ├─ validator
   │  └─ validator.go
```

## Installation
### Prerequisites
- Golang v1.20+
- Docker (optional for containerized deployment)
- PostgreSQL or MySQL database

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/kasikorn-line-api.git
   cd kasikorn-line-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables in `config.yaml`.

## Configuration
Edit the `config.yaml` file to match your database and application settings:
```yaml
db:
  user: "your-db-user"
  password: "your-db-password"
  host: "localhost"
  port: "5432"
  name: "kasikorn_db"
port: "8080"
cors:
  allow_origins: "*"
rate_limiter:
  max_requests: 100
  expiration: 60
```

## Running the Application
### Using Go
```sh
go run cmd/main.go
```
### Using Docker
```sh
docker build -t kasikorn-line-api .
docker run -p 8080:8080 kasikorn-line-api
```

## API Endpoints
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/account/:account_id` | GET | Get account details |
| `/v1/account/user/:user_id` | GET | Get accounts by user |
| `/v1/banner/:user_id` | GET | Get banners for user |
| `/v1/debit/:card_id` | GET | Get debit card details |
| `/v1/debit/user/:user_id` | GET | Get debit cards by user |
| `/v1/transaction/:transaction_id` | GET | Get transaction details |
| `/v1/transaction/user/:user_id` | GET | Get transactions by user |
| `/v1/user/:user_id` | GET | Get user details |

## Testing
Run unit tests:
```sh
go test ./...
```

## Contributing
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Create a Pull Request

## Security Features
- **CORS Protection**
- **Rate Limiting**
- **Helmet Middleware** for setting secure HTTP headers
- 
## License
This project is licensed under the MIT License.

