vuln:
	govulncheck ./...

cli:
	go build -mod vendor -o bin/server cmd/server/main.go

debug:
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -api -ready-check=false -server-uri http://localhost:8081 -prefix /places

docker:
	docker build -t placeholder-client-www .

up:
	docker-compose up --abort-on-container-exit

lambda:
	if test -f main; then rm -f main; fi
	if test -f server.zip; then rm -f server.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/server/main.go
	zip server.zip main
	rm -f main
