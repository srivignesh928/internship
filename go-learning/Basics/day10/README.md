# 🚀 Day 10 – Swagger Documentation in Go

## 📌 Overview

On Day 10, I explored Swagger Documentation and learned how to automatically generate API documentation for Go applications using the Gin framework. Swagger provides an interactive UI that allows developers to view, understand, and test APIs directly from the browser.

The objective was to understand how API documentation is maintained in real-world backend projects and how it improves collaboration between developers.

---

## 🎯 Topics Covered

### 🔹 Swagger Fundamentals
- What is Swagger
- Importance of API documentation
- Interactive API testing
- Swagger UI

### 🔹 Swaggo for Go
- Installing Swagger dependencies
- Generating documentation using `swag init`
- Understanding generated files

### 🔹 API Documentation
- Documenting endpoints using annotations
- Grouping APIs using tags
- Defining request and response formats

### 🔹 Swagger UI
- Viewing documented APIs
- Executing API requests directly from browser
- Understanding request and response structures

---

## ⚙️ Features Implemented

### ✅ Swagger Setup
- Configured Swagger in a Gin application
- Generated Swagger documentation files
- Integrated Swagger UI

### ✅ API Documentation
Documented multiple endpoints using Swagger annotations.

### ✅ Interactive Testing
Tested APIs directly from Swagger UI without using Postman.

---

## 📁 Project Structure

```plaintext
day10/
│
├── main.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── README.md
```

---

## 🔗 API Endpoints

### ➤ GET /hello
Returns a welcome message.

### ➤ GET /users
Returns a list of users.

### ➤ POST /user
Creates a new user.

---

## ▶️ How to Run

### 1. Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

### 2. Generate Swagger Documentation

```bash
swag init
```

### 3. Run Application

```bash
go run main.go
```

### 4. Open Swagger UI

```plaintext
http://localhost:8080/swagger/index.html
```

---

## 🧪 Testing

### GET /hello
- Tested through Swagger UI
- Verified successful response

### GET /users
- Retrieved sample user data

### POST /user
- Submitted user details through Swagger UI
- Verified request and response handling

---

## 🧠 Key Learnings

- Understood the purpose of Swagger in backend development
- Learned how API documentation is generated automatically
- Configured Swagger with Gin framework
- Learned to document APIs using annotations
- Tested APIs directly from the browser using Swagger UI

---

## 📄 Files Created

- `main.go` → Gin application with documented APIs
- `docs/docs.go` → Generated Swagger documentation
- `docs/swagger.json` → Swagger specification in JSON format
- `docs/swagger.yaml` → Swagger specification in YAML format

---
