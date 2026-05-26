# 🚀 Day 6 – Real-Time Systems in Go (HTTP Streaming & WebSocket Chat)

## 📌 Overview

On Day 6, I explored real-time communication in backend development using Go. I implemented HTTP streaming and built a multi-user WebSocket-based chat system to understand continuous data flow and bidirectional communication.

This marks a shift from traditional REST APIs to real-time backend systems.

---

## 🎯 Topics Covered

### 🔹 HTTP Streaming

* Implemented continuous server response using `net/http`
* Learned about:

  * Chunked transfer encoding
  * Response flushing (`http.Flusher`)
  * Browser buffering limitations
* Built a streaming endpoint that sends delayed messages

---

### 🔹 WebSocket Fundamentals

* Understood difference between HTTP and WebSocket
* Established persistent connection between client and server
* Used `gorilla/websocket` for implementation

---

### 🔹 Real-Time Communication

* Implemented message send/receive system
* Built an echo server for testing
* Handled continuous communication without closing connection

---

### 🔹 Multi-User Chat System

* Managed multiple WebSocket clients
* Implemented broadcasting mechanism
* Enabled real-time message sharing across multiple browser tabs

---

## ⚙️ Features Implemented

### ✅ HTTP Streaming API

* Sends data continuously with delay
* Demonstrates server push behavior

### ✅ WebSocket Server

* Upgrades HTTP to WebSocket connection
* Supports real-time communication

### ✅ Multi-User Chat

* Multiple users can connect simultaneously
* Messages are broadcasted to all connected users
* Simulates real-time chat system

---

## 📁 Project Structure

```plaintext
day6/
│
├── stream.go        → HTTP streaming implementation
├── websocket.go     → Basic WebSocket server
├── multiuser.go     → Multi-user chat (broadcast system)
├── minichat.go      → Improved chat logic
├── chat.html        → Chat UI (multi-user)
├── test.html        → WebSocket testing UI
└── README.md
```

---

## 🔗 Endpoints & Usage

### ➤ HTTP Streaming

* **Method:** GET
* **URL:** `/stream`

Output:

```plaintext
Message 1
Message 2
Message 3
...
```

---

### ➤ WebSocket Connection

* **URL:** `ws://localhost:8080/ws`

Supports:

* Sending messages
* Receiving real-time responses
* Multi-user broadcasting

---

## ▶️ How to Run

### 1. Install dependencies

```bash
go get github.com/gorilla/websocket
go mod tidy
```

---

### 2. Run Streaming Server

```bash
go run stream.go
```

---

### 3. Run WebSocket Server

```bash
go run multiuser.go
```

---

### 4. Test Chat System

* Open `chat.html` in multiple browser tabs
* Send messages from one tab
* Observe real-time updates in other tabs

---

## 🧪 Testing

* Tested streaming using browser and curl
* Tested WebSocket using:

  * Custom HTML UI
  * Multiple browser tabs
* Verified real-time communication and broadcasting

---

## 🧠 Key Learnings

* Difference between REST APIs and real-time systems
* HTTP streaming enables continuous server response
* WebSocket enables persistent bidirectional communication
* Managing multiple clients requires connection tracking
* Broadcasting is core to real-time applications like chat systems

---

## ⚠️ Notes

* Browsers may buffer streaming responses
* WebSocket requires `ws://` protocol
* Multiple tabs simulate multiple users
* Browser security policies may block testing in certain environments

---

## what i done in simple terms

* http : its is normal one way communication where the client send the request and the server respond to them ONCE like the normal CRUD operation.

* http streaming : it is same but here the server sends the response concurrently and continuously with some time delay for each messages.

ex: push and pull,publish and subscriber ,client-server etc .....

* web socket : it is an two way communication where both the client and server will do the request and responses. 

ex: exclusive-pair

---