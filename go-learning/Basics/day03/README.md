# 🚀 Day 3 – PostgreSQL Integration with Go (Gin + GORM)

## 📌 Overview

On Day 3, I focused on integrating a real database (PostgreSQL) with my Go backend. I built a complete REST API using Gin and GORM with full CRUD operations.

---

## 🎯 Topics Covered

### 1. PostgreSQL

* Installed and set up PostgreSQL
* Created database and tables
* Performed basic SQL operations (INSERT, SELECT, UPDATE, DELETE)

### 2. GORM with PostgreSQL

* Connected Go application to PostgreSQL
* Used structs as models for database tables
* Performed automatic migration using `AutoMigrate`

### 3. Gin Framework

* Built REST APIs using Gin
* Handled JSON requests and responses
* Implemented multiple API endpoints

---

## ⚙️ Features Implemented

* Create User (POST)
* Get All Users (GET)
* Update User (PUT)
* Delete User (DELETE)

---

## 📁 Files in this Folder

* `postgres.go` → Main backend file with PostgreSQL, Gin, and GORM integration

---

## 🔗 API Endpoints

### ➤ Create User

* **Method:** POST
* **URL:** `/user`

```json
{
  "name": "srivignesh",
  "age": 21
}
```

---

### ➤ Get Users

* **Method:** GET
* **URL:** `/users`

---

### ➤ Update User

* **Method:** PUT
* **URL:** `/user/:id`

```json
{
  "name": "updated",
  "age": 25
}
```

---

### ➤ Delete User

* **Method:** DELETE
* **URL:** `/user/:id`

---

## ▶️ How to Run

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Run the server

```bash
go run postgres.go
```

---

## 🧪 Testing

* Used **Postman** to test all API endpoints
* Verified CRUD operations with PostgreSQL database

---

## 🧠 Key Learnings

* Understood how backend APIs interact with databases
* Learned how to connect Go with PostgreSQL using GORM
* Gained hands-on experience with REST API development
* Understood difference between HTTP methods (GET, POST, PUT, DELETE)

---

## ⚠️ Notes

* PostgreSQL must be running before starting the server
* Ensure correct database credentials in the connection string
* Database file is not included in repository

---
## things handled for day3

* installed and done the setup on postgressql and done some of the basic query in sql as i know the query part before so revised and skiped some parts and done some basic CRUD operations

* Gin → receives request
GORM → talks to database
PostgreSQL → stores data

* driver installation
go get gorm.io/driver/postgres

* i run the part in URL : http://localhost:8080/user

* i done the CRUD operations of the api testing using the postman as i know this application before.

