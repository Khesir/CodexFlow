# Scrum App

A project management app (like Trello) with Go backend, React frontend, and PostgreSQL database.

## Tech Stack
- **Backend**: Go (Gin, pgx/sqlc, Docker)
- **Frontend**: React (Vite, TanStack/Zustand, Axios)
- **Database**: PostgreSQL
- **Desktop Option**: Wails (native) or Electron (fallback)

## Getting Started

### Backend
```bash
cd scrum-app
go mod init scrum-app
go get github.com/gin-gonic/gin
go run ./cmd/server
