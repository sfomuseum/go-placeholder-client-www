cli:
	go build -mod vendor -o bin/server cmd/server/main.go

debug:
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -api -ready-check=false -server-uri http://localhost:8081 -prefix /places

docker:
	docker build -t placeholder-client-www .

up:
	docker-compose up --abort-on-container-exit

lambda:
	@make bake
	if test -f main; then rm -f main; fi
	if test -f deployment.zip; then rm -f deployment.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/server/main.go
	zip deployment.zip main
	rm -f main
