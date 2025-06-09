# 🧵 Job Queue System in Golang

A lightweight, concurrent job queue system built using Go’s native features like goroutines and channels. Submit jobs via a REST API and process them asynchronously with worker routines. Inspired by real-world systems like Celery and Bull, but simplified and written from scratch in Go.

---

## 🚀 Features

- **Concurrent Workers** using goroutines
- **Job Dispatcher** with buffered channel queue
- **REST API** to submit jobs and check status
- **Simulated Job Processing** with random failures
- **Dockerized** for easy deployment
- Clean and production-friendly Go architecture

---

## 📦 Tech Stack

- **Language:** Golang (1.22+)
- **Router:** Gorilla Mux
- **Concurrency:** Native goroutines and channels
- **Data Store:** In-memory (map with sync.Mutex)
- **Optional Add-ons:** Redis, Web Dashboard, Auth, Retry Logic

---

## 📂 Project Structure

```bash
job-queue-system/
├── handler.go # HTTP handlers
├── job.go # Job struct & status logic
├── main.go # Entry point
├── queue.go # Dispatcher & worker logic
├── go.mod # Go module config
├── Dockerfile # Docker setup
└── README.md # Project docs
```
---

## 🛠️ Setup Instructions

### Clone and Run Locally

```bash
git clone https://github.com/nikhil-tiwary61/job-queue-system.git
cd job-queue-system
go mod tidy
go run .
```

Server will run at `http://localhost:8080`

---

## 🧪 API Endpoints

### Submit Job

```bash
POST /submit
```

Request Body (JSON)
```bash
{
  "payload": "Send welcome email"
}
```

Response
```bash
{
  "job_id": "123e4567-e89b-12d3-a456-426614174000"
}
```


### Get Job Status

```bash
GET /status/{job_id}
```

Example Response
```bash
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "status": "DONE",
  "created_at": "2025-06-09T17:45:00Z",
  "payload": "Send welcome email"
}
```

---

## 🐳 Docker Instructions

Build Image
```bash
docker build -t jobqueue .
```
Run Container
```bash
docker run -p 8080:8080 jobqueue
```
---

## 📌 Possible Enhancements (Roadmap)
- Use Redis for persistent job tracking
- Add retry logic for failed jobs
- Add job prioritization (e.g., high/low)
- Web dashboard for real-time job monitoring
- Rate limiting, metrics, and health checks
- Deploy on Render / Heroku / DigitalOcean

---

## 🤝 Contributing
Pull requests are welcome. Feel free to fork and build your own version!

---

## 📄 License
This project is open-source under the MIT License.