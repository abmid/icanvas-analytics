BINARY=icanvas

test: 
	go test -v ./...

app:
	cd web/app && npm install && npm run build
	cd db && go run main.go -status=up
	go build -o cmd/dist/${BINARY} cmd/app/main.go
	chmod +x cmd/dist/${BINARY}
	cd cmd/dist && ./${BINARY}

docker:
	docker-compose up -d

docker-stop:
	docker-compose down

test-short:
	go test -short  ./...

.PHONY: all test clean