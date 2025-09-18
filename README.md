# CodexFlow

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![Node Version](https://img.shields.io/badge/Node-18+-339933?logo=node.js&logoColor=white)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

**CodexFlow** — a lightweight project management system for organizing tasks, iterations, and workflows.  
Built with **Go (Gin)** for backend, **React + Vite** for frontend, and designed to scale into both **web** and **desktop (via Wails)** apps.

## Features

- Buckets, Projects, Tasks, and Iterations management  
- REST API backend in Go (Gin)  
- PostgreSQL database support  
- Modern frontend with React + Vite + TypeScript  
- Axios for API calls  
- Docker-ready for hosting  
- Future-ready for desktop builds with **Wails**


## Project Structure

```
codexflow/
├── cmd/ # Go backend entrypoint(s)
│ └── main.go
├── models/ # Domain models (Task, Project, etc.)
├── frontend/ # React + Vite frontend
│ ├── src/
│ ├── public/
│ └── package.json
├── wails.json
├── scripts/ # Dev/build helper scripts
├── Makefile # Build automation
├── go.mod # Go module definition
├── go.sum # Go module checksums
└── README.md
```

---

## ⚡ Requirements

- [Go](https://go.dev/dl/) (>= 1.21)  
- [Node.js](https://nodejs.org/) (>= 18) + npm or yarn  
- [Air](https://github.com/air-verse/air) (for hot reload in backend)  
- [Wails](https://wails.io/) (optional, for desktop builds)

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

### 5. Build Go backend

```sh
go build -o codexflow ./cmd
```

This will embed frontend/dist into the binary if you configured embed.FS.

## Desktop (Optional)

If you want to run CodexFlow as a native desktop app:

### 1. Install Wails

```sh
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 2. Run in dev

```sh
wails dev
```

### 3. Build desktop app

```sh
wails build
```

You’ll get a single .exe (Windows) or .app (macOS).

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
