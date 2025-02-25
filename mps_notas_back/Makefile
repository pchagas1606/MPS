.PHONY: build clean run docker-build docker-run

APP_NAME = mps_notas_back
GO = go
GOFLAGS = -v
LDFLAGS = -ldflags "-s -w"
DOCKER_IMAGE = $(APP_NAME):latest
PORT = 8080

build:
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o bin/$(APP_NAME) ./cmd/api/main.go

clean:
	rm -rf bin/

run: build
	./bin/$(APP_NAME)

# Docker
docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p 3000:3000 $(APP_NAME):latest

docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

# Hot reload (requer air: go install github.com/air-verse/air@latest)
dev:
	air