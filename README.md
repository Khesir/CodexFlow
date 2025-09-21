# CodexFlow

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
![Node Version](https://img.shields.io/badge/Node-22.19.0-339933?logo=node.js&logoColor=white)
![npm Version](https://img.shields.io/badge/npm-10.8.3-CB3837?logo=npm&logoColor=white)

**CodexFlow** — a lightweight project management system for organizing tasks, iterations, and workflows.
Built with **Go (Gin + sqlx)** for the backend, **React + Vite + shadcn/ui** for the frontend, and packaged into a desktop app with **Electron**.

## Features

- Buckets, Projects, Tasks, and Iterations management
- REST API backend with Go (Gin)
- PostgreSQL database + migrations
- Modern frontend with React + Vite + shadcn/ui (Tailwind-based)
- Axios for API calls
- Electron integration → single-binary desktop app
- Docker-ready for hosting backend standalone

## Project Structure

```
codexflow/
├── cmd/                     # Go backend entrypoint(s)
│   └── main.go
├── internal/                # Internal Go packages
│   ├── api/                 # HTTP API (Gin handlers)
│   │   └── user/       
│   ├── core/                # Business logic
│   ├── db/                  # Database init/helpers
│   └── migrations/          # DB migrations
├── frontend/                # React + Vite + shadcn/ui frontend
│   ├── src/
│   ├── public/
│   └── package.json
├── electron/                # Electron app wrapper
│   └── main.js
├── scripts/                 # Dev/build helper scripts
├── docker-compose.dev.yml   # Development external dependencies
├── Makefile                 # Build automation
├── go.mod
├── go.sum
└── README.md

```

---

## ⚡ Requirements

- [Go](https://go.dev/dl/) (>= 1.25)  
- [Node.js](https://nodejs.org/) (>= 22) + npm or yarn  
- [Air](https://github.com/air-verse/air) (1.63.0) for hot reload in backend  

---

## 🔧 Setup

### 1. Clone the repo

```sh
git clone https://github.com/yourname/codexflow.git
cd codexflow

```

### 2. Install dependencies

Using Makefile:

```sh
make deps
```

This will:

- Run go mod download → fetch Go dependencies
- Run npm install in frontend/ → install frontend dependencies

(Manual equivalent: go mod download && cd frontend && npm install)

### 3. Run in development

- Start backend with hot reload:

```sh
air
```

- Start frontend dev server:

```sh
cd frontend
npm run dev
```

Your React app will proxy API calls to Go (configured in vite.config.ts).

### 4. Build frontend

```sh
cd frontend
npm run build
```

The frontend proxies API calls to the Go backend `(see vite.config.ts)`.

### 5. Build Go backend

```sh
go build -o codexflow ./cmd
```

### 6. Electron (Desktop wrapper)

```sh
cd electron
npm install
npm start
```

Electron will open a desktop window pointing at the dev frontend/backend.

## Docker (Optional)

```sh
docker build -t codexflow .
docker run -p 8080:8080 codexflow
```

## Makefile Commands

```sh
# Install backend + frontend deps
make deps

# Run backend with Air
make dev-backend

# Run frontend dev
make dev-frontend

# Build frontend
make build-frontend

# Build backend
make build-backend

```

## License

This project is licensed under the [MIT License](./LICENSE).
