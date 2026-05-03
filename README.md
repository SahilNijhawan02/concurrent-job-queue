# 🚀 Concurrent Job Processing System in Go

A lightweight backend system built in Go that demonstrates **asynchronous job processing using goroutines, channels, and worker pools**. The system allows users to create jobs via REST APIs and processes them concurrently in the background while enabling real-time status tracking.

---

## 🔗 Live Demo

👉 https://concurrent-job-queue.onrender.com/


## 🧠 Features

- ⚡ Asynchronous job execution (non-blocking API)
- 🧵 Concurrency using Goroutines
- 🔁 Worker Pool implementation
- 📦 Channel-based queue (Producer-Consumer pattern)
- 🔐 Thread-safe operations using Mutex
- 🌐 REST API using Go's `net/http`
- ☁️ Deployed on Render

---

## 🔧 Tech Stack

- **Language:** Go (Golang)
- **Concepts:** Goroutines, Channels, Mutex, WaitGroup
- **API:** net/http (no frameworks)
- **Deployment:** Render

---

## 📡 API Endpoints

### ➕ Create Job
```http
POST /jobs
```

### 🔍 Get Job Status
```http
GET /job?id={jobId}
```

---

## 🧱 Architecture

```
Client → HTTP Handler → Job Queue (Channel) → Worker Pool → Job Store
```

---

## 🧪 How It Works

1. Client sends a request to create a job  
2. Server generates a unique job ID and pushes it into a channel (queue)  
3. Worker goroutines pick jobs from the queue  
4. Jobs are processed asynchronously  
5. Status is updated in a shared in-memory store  
6. Client can fetch job status using the job ID  

---

## 🎯 Use Cases

- Background task processing (emails, notifications)
- Task queues
- Distributed systems
- Asynchronous workflows

---

## 💡 Why Go?

Go provides a simple yet powerful concurrency model using **goroutines and channels**, making it highly efficient for building scalable backend systems compared to traditional threading approaches.

---

## 🧠 Key Learning

> This project demonstrates how to design a concurrent system using Go’s concurrency primitives and implement a scalable worker pool with thread-safe data handling.

---

## ⚠️ Note

- Data is stored in-memory (not persistent)
- Restarting the server will reset all jobs
