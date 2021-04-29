dependencies:
	go mod download

run:
	go run api/main.go

seed:
	go run seeders/main.go

test_unit:
	go test -v ./...

test_integration:
	go test -v 