# 🚀 Day 2 – Go Backend Basics (Goroutines, Gin, GORM)

## 📌 Overview

On Day 2 of my Go learning journey, I explored core backend concepts including concurrency, web frameworks, and database integration.

---

## 🎯 Topics Covered

### 1. Goroutines

* Learned how to run functions concurrently using `go` keyword
* Understood basic concurrency behavior
* Used `time.Sleep` to manage execution timing

### 2. Gin (Web Framework)

* Set up a basic HTTP server
* Created API endpoints using `GET` and `POST`
* Handled JSON requests and responses

### 3. GORM (ORM for Go)

* Connected Go application to a database (SQLite)
* Created models using structs
* Performed basic operations:

  * Auto migration (table creation)
  * Insert data into database

---

## 🧪 Files in this Folder

* `routines.go` → Basic goroutine example
* `gin.go` → Simple API server using Gin
* `gorm.go` → Database connection and CRUD basics using GORM

---

## ⚙️ How to Run

### 1. Initialize module (if not already)

```bash
go mod tidy
```

### 2. Run files individually

```bash
go run routines.go
go run gin.go
go run gorm.go
```

---

## 🔗 API Endpoints (Gin)

### GET Request

```
http://localhost:8080/
```

### POST Request (example)

```
http://localhost:8080/user
```

Sample JSON:

```json
{
  "name": "srivignesh"
}
```

---

## 🧠 Key Learnings

* Goroutines allow concurrent execution in Go
* Gin simplifies backend API development
* GORM helps interact with databases without writing SQL

---

## ⚠️ Notes

* Used SQLite for simplicity (local database)
* Database file (`test.db`) is ignored using `.gitignore`

---

## 🚀 Next Plan

* Learn proper goroutine control using `WaitGroup`
* Build full CRUD APIs (Create, Read, Update, Delete)
* Improve project structure for real-world backend applications

---
