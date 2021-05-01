dependencies:
	go mod download

docker-fresh:
	go run seeders/main.go down && go run seeders/main.go up && go run api/main.go 

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

document:
	godoc -http=:6060