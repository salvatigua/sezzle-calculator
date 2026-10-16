# Full-Stack Calculator Application

A modern, full-stack calculator application built with a **React (TypeScript)** frontend and a **Go (REST API)** backend microservice.

---

## Features

- **Basic Operations:** Addition, Subtraction, Multiplication, Division.
- **Advanced Operations:** Exponentiation ($x^y$), Square Root ($\sqrt{x}$), Percentage (%).
- **Error Handling:** Graceful handling of edge cases such as division by zero and square roots of negative numbers.
- **Responsive UI:** Clean and intuitive UI built with Tailwind CSS.
- **Dockerized Deployment:** Containerized with Docker and orchestrated via Docker Compose.

---

## Tech Stack

- **Frontend:** React 18, TypeScript, Vite, Tailwind CSS, Vitest, React Testing Library.
- **Backend:** Go 1.22, Chi Router, Go standard `testing` library.
- **DevOps:** Docker, Docker Compose, Nginx.

---

## Project Structure

```text
sezzle-calculator/
├── backend/
│   ├── cmd/api/main.go         # API Entrypoint
│   ├── internal/
│   │   ├── calculator/         # Mathematical core logic & unit tests
│   │   └── handler/            # HTTP handlers, validation & API tests
│   ├── Dockerfile
│   └── go.mod
├── frontend/
│   ├── src/
│   │   ├── api/                # API client
│   │   ├── components/         # Calculator component & unit tests
│   │   └── types/              # TypeScript contracts
│   ├── Dockerfile
│   └── package.json
├── docker-compose.yml
└── README.md

## Getting Started

### Option 1: Running with Docker Compose (Recommended)

Make sure you have Docker Desktop installed and running.

1. Clone the repository:

git clone <repository-url>
cd sezzle-calculator

2. Build and start the application:

docker compose up --build

3. Access the application:

Frontend: http://localhost:3000

Backend API: http://localhost:8081

### Option 2: Running Locally

#### Backend (Go)

1. Navigate to the backend directory:

cd backend

2. Run unit tests:

go test -v -cover ./...

3. Start the Go API server:

go run cmd/api/main.go

#### Frontend (React)

1. Open a new terminal and navigate to the frontend directory:

cd frontend

2. Install dependencies:

npm install

3. Run unit tests:

npx vitest run

4. Start the development server:

npm run dev

5. Open your browser at http://localhost:5173.

## API Documentation

POST /api/v1/calculate
Performs an arithmetic operation.

### Request Payload Example

{
  "operator": "add",
  "a": 15,
  "b": 25
}

### Supported Operations

add: Requires a and b.

subtract: Requires a and b.

multiply: Requires a and b.

divide: Requires a and b.

power: Requires a and b.

sqrt: Requires a.

percentage: Requires a (or a and b).

### Success Response (200 OK)

{
  "result": 40
}

### Error Response (422 Unprocessable Entity)

{
  "error": "La división por cero no está definida."
}

### Design Rationale

Go for Backend: Chosen for its simplicity, concurrency model, and speed. The standard library and Chi router allow a lightweight architecture without unnecessary abstraction layers.

Multi-Stage Docker Builds: Minimizes image size by isolating the compilation environment from the execution runtime (Alpine Linux / Nginx).

Decoupled Architecture: Separation of domain logic (calculator), HTTP handling (handler), and stateful UI (frontend).

