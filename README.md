# Project Title: Multi-Database Integration with Go

## Overview
This project demonstrates how to integrate and interact with multiple databases (MySQL, PostgreSQL, SQLite, Redis) using Go. It covers the creation of tables, data insertion, querying, and basic CRUD operations. Docker Compose is used to manage containerized database instances for MySQL, PostgreSQL, and Redis.

## Features
- **MySQL Integration**: Demonstrates table creation, data insertion, and querying using Go's `database/sql` package.
- **PostgreSQL Integration**: Includes examples of table creation, data insertion, and querying using `pgxpool` and `sqlc` for query generation.
- **SQLite Integration**: Implements database operations with and without CGO support.
- **Redis Integration**: Provides examples for setting and retrieving key-value pairs.
- **Dockerized Databases**: Uses Docker Compose for managing database containers.

## Technologies Used
- **Programming Language**: Go
- **Databases**:
  - MySQL
  - PostgreSQL
  - SQLite
  - Redis
- **Tools**:
  - Docker & Docker Compose
  - SQLC for query generation

## Project Structure
```plaintext
├── database
│   ├── mysql
│   │   └── mySQL.go
│   ├── postres
│   │   └── postgres.go
│   ├── redis
│   │   └── redis.go
│   ├── sqlite-integration
│   │   └── sqliteMain.go
│   ├── sqlite-without-cgo
│   │   └── sqliteWithoutCGO.go
│   └── databaseMain.go
├── sql
│   ├── postgresMain.go
│   ├── migrations
│   └── sqlc.yaml
├── main.go
├── docker-compose.yaml
└── README.md
```

## Setup Instructions

### Prerequisites
- [Go](https://go.dev/dl/) installed
- [Docker](https://www.docker.com/) installed
- [Docker Compose](https://docs.docker.com/compose/) installed

### Steps to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

2. Start the database services with Docker Compose:
   ```bash
   docker-compose up -d
   ```

3. Install Go dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

### Environment Configuration
- MySQL connection string: `root:root@tcp(0.0.0.0:3306)/mysql_db`
- PostgreSQL connection string: `postgres://postgres:postgres@localhost:5432/postgres_db`
- Redis connection: `localhost:6379`

## Key Functionalities

### MySQL Integration
- **File**: `database/mysql/mySQL.go`
- **Features**:
  - Creates a table `foo`
  - Inserts a record
  - Queries the table and retrieves results

### PostgreSQL Integration
- **File**: `database/postres/postgres.go`
- **Features**:
  - Creates a table `foo`
  - Inserts a record
  - Queries the table and retrieves results

### SQLite Integration
- **File**: `database/sqlite-integration/sqliteMain.go`
- **Features**:
  - Creates a table `foo`
  - Inserts a record
  - Queries the table and retrieves results

- **File**: `database/sqlite-without-cgo/sqliteWithoutCGO.go`
- **Features**:
  - Demonstrates SQLite integration without CGO dependency

### Redis Integration
- **File**: `database/redis/redis.go`
- **Features**:
  - Sets and retrieves key-value pairs
  - Demonstrates key expiration with TTL

### Postgres Query Generation with SQLC
- **File**: `sql/postgresMain.go`
- **Features**:
  - Creates authors using `CreateAuthor`
  - Lists authors with `ListAuthors`

### Docker Compose
- **File**: `docker-compose.yaml`
- **Services**:
  - MySQL
  - PostgreSQL
  - Redis

## Usage Examples
### MySQL
```go
mysql.DatabaseMySQL()
```

### PostgreSQL
```go
postres.DatabasePostgres()
```

### SQLite
```go
sqliteintegration.SqliteMain()
```

### Redis
```go
redis.RedisDatabase()
```

### Running SQLC Queries
```go
postgres_db.PostgresMain()
```

## Acknowledgments
- [Go-SQL-Driver](https://github.com/go-sql-driver/mysql)
- [PGX](https://github.com/jackc/pgx)
- [Redis Go Client](https://github.com/redis/go-redis)
- [SQLC](https://sqlc.dev/)
