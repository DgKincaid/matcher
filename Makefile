dependencies:
	go mod download

run:
	go run api/main.go

seed-up:
	go run seeders/main.go up

seed-down:
	go run seeders/main.go down

test-unit:
	go test -v ./...

test-integration:
	go test -v 