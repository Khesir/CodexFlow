APP_NAME = scrum-app
FRONTEND_DIR = frontend
BACKEND_DIR = cmd/server
BUILD_DIR = build

# Default target
all: build

# Install frontend deps
frontend-deps:
	cd $(FRONTEND_DIR) && npm install

# Build frontend (Vite -> dist)
frontend-build: frontend-deps
	cd $(FRONTEND_DIR) && npm run build

# Build Go backend (embed frontend)
backend-build:
	cd $(BACKEND_DIR) && go build -o ../../$(APP_NAME)

# Full build (frontend + backend, local OS/Arch)
build: frontend-build backend-build

# Cross compile for multiple platforms
release: frontend-build
	mkdir -p $(BUILD_DIR)
	# Windows
	cd $(BACKEND_DIR) && GOOS=windows GOARCH=amd64 go build -o ../../$(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe
	# Linux
	cd $(BACKEND_DIR) && GOOS=linux GOARCH=amd64 go build -o ../../$(BUILD_DIR)/$(APP_NAME)-linux-amd64
	# macOS (Intel)
	cd $(BACKEND_DIR) && GOOS=darwin GOARCH=amd64 go build -o ../../$(BUILD_DIR)/$(APP_NAME)-darwin-amd64
	# macOS (Apple Silicon)
	cd $(BACKEND_DIR) && GOOS=darwin GOARCH=arm64 go build -o ../../$(BUILD_DIR)/$(APP_NAME)-darwin-arm64

# Run dev server (Go only, use Vite proxy for FE in dev)
dev:
	cd $(BACKEND_DIR) && go run main.go

clean:
	rm -f $(APP_NAME)
	rm -rf $(FRONTEND_DIR)/dist
	rm -rf $(BUILD_DIR)