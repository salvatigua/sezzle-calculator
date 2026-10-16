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
```

## Getting Started

### Option 1: Running with Docker Compose (Recommended)

Make sure you have [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running.

1. Clone the repository:

```bash
   git clone <repository-url>
   cd sezzle-calculator
```

2. Build and start the application:

```bash
docker compose up --build
```

3. Access the application:

```bash
Frontend: http://localhost:3000

Backend API: http://localhost:8081
```

### Option 2: Running Locally

#### Backend (Go)

1. Navigate to the backend directory:

```bash
cd backend
```

2. Run unit tests:

```bash
go test -v -cover ./...
```

3. Start the Go API server:

```bash
go run cmd/api/main.go
```

#### Frontend (React)

1. Open a new terminal and navigate to the frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
```

3. Run unit tests:

```bash
npx vitest run
```

4. Start the development server:

```bash
npm run dev
```

5. Open your browser at http://localhost:5173.

## API Documentation

```bash
POST /api/v1/calculate
```
Performs an arithmetic operation.

### Request Payload Example
```json
{
  "operator": "add",
  "a": 15,
  "b": 25
}
```

### Supported Operations

- add: Requires a and b.
- subtract: Requires a and b.
- multiply: Requires a and b.
- divide: Requires a and b.
- power: Requires a and b.
- sqrt: Requires a.
- percentage: Requires a (or a and b).

### Success Response (200 OK)

```json
{
  "result": 40
}
```

### Error Response (422 Unprocessable Entity)

```json
{
  "error": "La división por cero no está definida."
}
```

### Design Rationale

1. **Go for Backend**: Chosen for its simplicity, fast execution, and explicit error handling. Utilizing standard packages alongside the lightweight Chi router allows for high performance without unnecessary framework overhead.
2. **Multi-Stage Docker Builds**: Minimizes image size and enhances security by separating the build toolchain from the lightweight runtime container (Alpine Linux & Nginx).
3. **Decoupled Architecture**: Clean separation between domain calculation logic (calculator), HTTP routing and transport (handler), and the component-based UI (frontend).
4. **Resilient UX**: Integrated client-side debouncing on network interactions to prevent UI flickering during rapid state calculations.

---

## AI Tooling & Prompts Used

In accordance with the assessment guidelines, AI assistance was utilized throughout the development process for scaffolding, test generation, and architectural refining.

### Key Prompts Executed

1. **Backend Scaffolding & Domain Logic (Go):**
   > *"Create a clean Go backend using chi router for a calculator REST API supporting add, subtract, multiply, divide, power, sqrt, and percentage. Include domain errors for division by zero and negative square roots."*

2. **Backend Unit & Integration Testing:**
   > *"Generate table-driven unit tests for the calculator service and HTTP integration tests for the handler in Go, aiming for high statement coverage."*

3. **Frontend UI & State Management (React + TypeScript):**
   > *"Create a responsive React component in TypeScript with Tailwind CSS for a full-stack calculator. Implement keypad input, error handling banners, and async API integration with a 300ms debounced loading state."*

4. **Frontend Unit Testing (Vitest):**
   > *"Write unit tests using Vitest and React Testing Library for the Calculator component, mocking the API client and verifying key display updates and AC button behavior."*

5. **DevOps & Containerization:**
   > *"Write multi-stage Dockerfiles for a Go backend and a React/Vite frontend (served with Nginx), orchestrated via docker-compose."*