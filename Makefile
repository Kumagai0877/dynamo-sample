BASE_DOCKER_COMPOSE = ./build/docker-compose.yml
COMPOSE_OPTS        = -f "$(BASE_DOCKER_COMPOSE)"
LIMIT = 1
DEVICE_ID := 123

go-build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.artifacts/api-linux-amd64 ./cmd/

build: go-build
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose $(COMPOSE_OPTS) build --no-cache --parallel
	docker-compose $(COMPOSE_OPTS) up -d

setup: build

up: go-build
	docker-compose $(COMPOSE_OPTS) up -d

stop:
	docker-compose $(COMPOSE_OPTS) stop

down:
	docker-compose $(COMPOSE_OPTS) down

logs:
	docker-compose $(COMPOSE_OPTS) logs -f

test:
	go test -v ./...

lint:
	go fmt ./...
	golangci-lint run
	go vet ./...

exec:
	 go run cmd/main.go --deviceID ${DEVICE_ID}