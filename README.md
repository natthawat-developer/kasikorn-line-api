# Kasikorn Line API

## Overview
Kasikorn Line API is a microservice-based API built with Golang and Fiber. It provides account, banner, debit, transaction, and user management functionalities.

Note: Let's discuss further details during the interview.

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
│
└─ pkg
   ├─ database
   │  └─ database.go
   ├─ error
   │  └─ error.go
   ├─ health
   │  └─ health.go
   ├─ log
   │  └─ logger.go
   ├─ security
   │  ├─ cors.go
   │  ├─ helmet.go
   │  └─ ratelimit.go
   ├─ utils
   │  ├─ masking.go
   │  └─ masking_test.go
   └─ validator
      └─ validator.go
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

### **Service Optimization**  

To achieve **optimized service latency**, consider the following:

1. **Architectural Design**  
   - Implement a **Backend-for-Frontend (BFF)** pattern, followed by an **Orchestrator**, then a **Core Service**, and finally the **Database**.  
   - Use **Master-Slave database replication** for high read/write performance.  

2. **Load Balancing**  
   - Implement a **latency-aware load balancing algorithm** to dynamically route requests to the fastest available service instance.  

3. **Performance Tuning**  
   - **Step 1: Optimize Database Queries** to ensure efficient indexing and query execution plans.  
   - **Step 2: Refactor Service Layer** to enhance responsiveness. If a service requires multiple external API calls, use **Goroutines** to execute them concurrently and reduce latency.  
   - **Step 3: Use gRPC** instead of REST when high performance is required, as gRPC offers lower latency and better efficiency over HTTP/2.  

### **Database Optimization**  

To improve database performance and integrity, consider the following:

1. **Foreign Key Constraints**  
   - Define **foreign keys** to enforce data integrity. For example:
   ```sql
   ALTER TABLE account_balances ADD CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(account_id);
   ALTER TABLE account_details ADD CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(account_id);
   ```
   - This ensures referential integrity and prevents orphaned records.  

2. **Indexing Frequently Queried Columns**  
   - Add indexes to speed up queries, especially for columns like `user_id` in `transactions`, `accounts`, and `debit_cards`.  
   ```sql
   CREATE INDEX idx_user_id ON transactions(user_id);
   ```

3. **Optimizing Primary Keys**  
   - Instead of using `VARCHAR(50)`, consider **UUID (`CHAR(36)`)** or **BIGINT** for primary keys to improve indexing efficiency.  

4. **Removing or Documenting Dummy Columns**  
   - If columns like `dummy_col_1`, `dummy_col_2`, etc., are placeholders, either **remove them** or provide proper documentation for their intended use.  

5. **Enhancing the Transactions Table**  
   - Add essential fields to make the transaction records more meaningful:  
   ```sql
   ALTER TABLE transactions
   ADD COLUMN amount DECIMAL(15,2) NOT NULL,
   ADD COLUMN transaction_type ENUM('credit', 'debit') NOT NULL,
   ADD COLUMN status ENUM('pending', 'completed', 'failed') DEFAULT 'pending',
   ADD COLUMN timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
   ```
   - This helps in tracking financial transactions more effectively.  

## License
This project is licensed under the MIT License.