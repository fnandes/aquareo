pi:
	GOOS=linux GOARCH=arm GOARM=7 go build -o aquareo ./cmd/aquareo

clean:
	go clean
	rm aquareo

test:
	go test ./...