CWD=$(shell pwd)

bake-assets:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	go build -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	# rm -f www/static/*~ www/static/css/*~ www/static/javascript/*~
	# @PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg http static/javascript static/css
	# mv bindata.go http/assetfs.go
	rm -rf templates/html/*~
	bin/go-bindata -pkg templates -o assets/templates/html.go templates/html

docker:
	go mod vendor
	@make bake-assets
	docker build -t placeholder-client-www .

lambda:
	# go mod vendor
	@make bake-assets
	if test -f main; then rm -f main; fi
	if test -f deployment.zip; then rm -f deployment.zip; fi
	GOOS=linux go build -mod vendor -o main cmd/server/main.go
	zip deployment.zip main
	rm -f main
