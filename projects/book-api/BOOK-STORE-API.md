# REST API for a Bookstore

Goal: Build a REST API to manage a collection of books.

Features:
Add, update, delete, and fetch book records.
CRUD operations using HTTP methods (GET, POST, PUT, DELETE).
Use a database like SQLite for persistence.

Learning: Building RESTful services, working with databases, swagger, structuring APIs in Go.

1. Setup:
    - Go Environment: Ensure Go is installed.
    - Database: Use SQLite or PostgreSQL for persistence.
    - Go Modules: Initialize your project with Go modules.
```bash
go mod init book-api
```
2. Packages to Install:
    - Use `gorilla/mux` for routing.
    - Use `github.com/mattn/go-sqlite3` for SQLite.
    - Use `gorm` for ORM (optional, if you prefer to work directly with SQL).
```bash
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/mattn/go-sqlite3
```
3. Run the API:
```bash
go run main.go
```
4. Swagger Docs: http://localhost:8000/swagger/