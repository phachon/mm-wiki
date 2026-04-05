.PHONY: help dev dev-backend dev-frontend build build-backend build-frontend \
        docker-up docker-down clean lint test

help:
	@echo "MM-Wiki Development Commands"
	@echo "============================"
	@echo ""
	@echo "  make dev              Start both backend and frontend for development"
	@echo "  make dev-backend      Start backend only (Go, port 8088)"
	@echo "  make dev-frontend     Start frontend only (React, port 3000)"
	@echo ""
	@echo "  make build            Build both backend and frontend"
	@echo "  make build-backend    Build backend binary"
	@echo "  make build-frontend   Build frontend static files"
	@echo ""
	@echo "  make docker-up        Start MySQL with Docker Compose"
	@echo "  make docker-down      Stop MySQL Docker container"
	@echo ""
	@echo "  make lint             Lint backend and frontend code"
	@echo "  make test             Run backend tests"
	@echo "  make clean            Clean all build artifacts"

dev:
	@echo "Starting development servers..."
	@echo "Backend: http://127.0.0.1:8088"
	@echo "Frontend: http://localhost:3000"
	@echo ""
	@echo "Run in separate terminals:"
	@echo "  make dev-backend"
	@echo "  make dev-frontend"

dev-backend:
	cd backend && $(MAKE) dev ENV=dev

dev-frontend:
	cd front && npm start

build: build-backend build-frontend

build-backend:
	cd backend && $(MAKE) build

build-frontend:
	cd front && npm run build

docker-up:
	docker compose up -d

docker-down:
	docker compose down

lint:
	cd backend && $(MAKE) lint
	cd front && npx eslint src/ --ext .ts,.tsx

test:
	cd backend && $(MAKE) test

clean:
	cd backend && $(MAKE) clean
	rm -rf front/build
