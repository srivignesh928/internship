# 🚀 Day 4 – Middleware & Panic Recovery in Go (Gin)

## 📌 Overview

On Day 4, I learned about middleware in Go using the Gin framework and implemented custom middleware for logging and panic recovery. This helped me understand how requests and responses are handled internally in backend systems.

---

## 🎯 Topics Covered

### 1. Middleware Concept

* Understood how middleware works as an intermediate layer
* Observed request flow:

```plaintext
Request → Middleware → Route → Middleware → Response
```

* Learned how middleware executes both **before and after** a request

---

### 2. Logger Middleware

* Implemented a custom logger middleware
* Tracked incoming requests and outgoing responses

Example behavior:

```plaintext
➡️ Request started: /route
⬅️ Response sent: /route
```

---

### 3. Panic & Recovery

* Learned what a panic is (runtime crash)
* Understood how unhandled panic can stop the server
* Implemented recovery middleware using:

  * `defer`
  * `recover()`

---

### 4. Panic Recovery Middleware

* Prevented server crashes
* Returned proper error response to client
* Logged panic details in terminal

---

## ⚙️ Features Implemented

* Custom Logger Middleware
* Custom Recovery Middleware
* Panic testing route
* Safe error handling without crashing server

---

## 📁 Files in this Folder

* `main.go` → Contains middleware implementation and test routes

---

## 🔗 API Endpoints

### ➤ Normal Route

* **Method:** GET
* **URL:** `/`

Response:

```plaintext
Hello Srivignesh
```

---

### ➤ Panic Route

* **Method:** GET
* **URL:** `/panic`

Response:

```json
{
  "error": "Internal Server Error"
}
```

---

## ▶️ How to Run

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Run the server

```bash
go run main.go
```

---

## 🧪 Testing

* Tested endpoints using browser/Postman
* Verified:

  * Middleware execution flow
  * Panic handling without server crash

---

## 🧠 Key Learnings

* Middleware acts as a control layer in backend systems
* Request flow can be customized using middleware
* Panic can crash the server if not handled
* Recovery middleware ensures server stability
* Understood difference between:

  * `gin.Default()` vs `gin.New()`

---

## ⚠️ Notes

* `gin.Default()` already includes built-in Logger and Recovery
* Custom middleware helps understand internal working
* Terminal output is used for logging, not browser

---

## simple

* middleware - before and after request controller

* panic - recovery used to protect from code or system crash.i shows the error msg to the user but they are keep running on the background

* can be tested with the postman ,it returns error for panic