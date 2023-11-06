GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

vuln:
	govulncheck ./...

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/server cmd/server/main.go

debug:
	go run -mod $(GOMOD) cmd/server/main.go \
		-nextzen-apikey $(APIKEY) \
		-api \
		-ready-check=false \
		-server-uri http://localhost:8081 

docker:
	docker build -t placeholder-client-www .

up:
	docker-compose up --abort-on-container-exit

lambda:
	if test -f bootstrap; then rm -f bootstrap; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOARCH=arm64 GOOS=linux go build -mod $(GOMOD) -ldflags="-s -w" -tags lambda.norpc -o bootstrap cmd/server/main.go
	zip server.zip bootstrap
	rm -f bootstrap
