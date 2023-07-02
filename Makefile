.PHONY: run build start-dependencies test

run: build
	REVIEW_DATABASE_USER=postgres 			\
	REVIEW_DATABASE_PASSWORD=postgres		\
	REVIEW_DATABASE_HOST=localhost          \
	REVIEW_DATABASE_DATABASE=postgres       \
	REVIEW_DATABASE_PORT=5432          		\
		./review-server

start-dependencies:
	docker-compose up -d

build: start-dependencies
	go build

test:
	go test ./...